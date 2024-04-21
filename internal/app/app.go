package app

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/Shemistan/platform_common/pkg/closer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/Shemistan/grpc_user_api/internal/config"
	"github.com/Shemistan/grpc_user_api/internal/interceptor"
	descAccess "github.com/Shemistan/grpc_user_api/pkg/access_api_v1"
	descAuth "github.com/Shemistan/grpc_user_api/pkg/auth_api_v1"
	descUser "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"

	_ "github.com/Shemistan/grpc_user_api/statik" // необходим что бы подтянуть статику при инициализации
)

// App - структура приложения
type App struct {
	serviceProvider   *serviceProvider
	grpcUserServer    *grpc.Server
	grpcAuthServer    *grpc.Server
	grpcAccessServer  *grpc.Server
	httpUserServer    *http.Server
	swaggerUserServer *http.Server
}

var configPath string

func init() {
	//flag.StringVar(&configPath, "config-path", ".env", "path to config file")
	configPath = ".env"
}

// NewApp - создать новый экземпляр структуры приложения
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run - запустить сервис
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := a.runGRPCUserServer()
		if err != nil {
			log.Fatalf("failed to run GRPC server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runGRPCAuthServer()
		if err != nil {
			log.Fatalf("failed to run GRPC auth server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runGRPCAccessServer()
		if err != nil {
			log.Fatalf("failed to run GRPC access server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := a.runSwaggerServer()
		if err != nil {
			log.Fatalf("failed to run Swagger server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCUserServer,
		a.initHTTPServer,
		a.initSwaggerServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCUserServer(ctx context.Context) error {
	a.grpcUserServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

	reflection.Register(a.grpcUserServer)

	descUser.RegisterUserV1Server(a.grpcUserServer, a.serviceProvider.UserAPI(ctx))

	return nil
}

func (a *App) initGRPCAuthServer(ctx context.Context) error {
	a.grpcAuthServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

	reflection.Register(a.grpcAuthServer)

	descAuth.RegisterAuthV1Server(a.grpcAuthServer, a.serviceProvider.AuthAPI(ctx))

	return nil
}

func (a *App) initGRPCAccessServer(ctx context.Context) error {
	a.grpcAccessServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

	reflection.Register(a.grpcUserServer)

	descAccess.RegisterAccessV1Server(a.grpcAccessServer, a.serviceProvider.AccessAPI(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := descUser.RegisterUserV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCUserConfig().Address(), opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpUserServer = &http.Server{
		Addr:              a.serviceProvider.HTTPConfig().Address(),
		Handler:           corsMiddleware.Handler(mux),
		ReadHeaderTimeout: 10 * time.Second,
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/api.swagger.json", serveSwaggerFile("/api.swagger.json"))

	a.swaggerUserServer = &http.Server{
		Addr:              a.serviceProvider.SwaggerConfig().Address(),
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	return nil
}

func (a *App) runGRPCUserServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCUserConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCUserConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcUserServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.httpConfig.Address())

	err := a.httpUserServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	log.Printf("Swagger server is running on %s", a.serviceProvider.SwaggerConfig().Address())

	err := a.swaggerUserServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving swagger file: %s", path)

		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Open swagger file: %s", path)

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close() //nolint:errcheck

		log.Printf("Read swagger file: %s", path)

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Write swagger file: %s", path)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Served swagger file: %s", path)
	}
}

func (a *App) runGRPCAuthServer() error {
	log.Printf("GRPC auth server is running on %s", a.serviceProvider.GRPCAuthConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCAuthConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcAuthServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPCAccessServer() error {
	log.Printf("GRPC access server is running on %s", a.serviceProvider.GRPCAccessConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCAccessConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcAccessServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
