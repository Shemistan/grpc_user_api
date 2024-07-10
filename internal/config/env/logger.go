package env

import (
	"os"
	"strconv"

	"github.com/pkg/errors"

	"github.com/Shemistan/grpc_user_api/internal/config"
)

const (
	logFileNameEnvName    = "LOG_FILE_NAME"
	logFileMaxSizeEnvName = "LOG_MAX_SIZE"
	logMaxBackupsEnvName  = "LOG_MAX_BACKUPS"
	logMaxAgeEnvName      = "LOG_MAX_AGE"
	logLevelEnvName       = "LOG_LEVEL"
	defaultLogLevel       = "info"
)

type zapLoggerConfig struct {
	fileName    string
	fileMaxSize int
	maxBackups  int
	maxAge      int
	logLevel    string
}

// NewZapLoggerConfig - новый конфиг для логгера
func NewZapLoggerConfig() (config.ZapLogger, error) {
	fileName := os.Getenv(logFileNameEnvName)
	if len(fileName) == 0 {
		return nil, errors.New("log file name not found")
	}

	fileMaxSizeString := os.Getenv(logFileMaxSizeEnvName)

	fileMaxSize, err := strconv.Atoi(fileMaxSizeString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log file max size")
	}

	maxBackupsString := os.Getenv(logMaxBackupsEnvName)
	maxBackups, err := strconv.Atoi(maxBackupsString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log max backups")
	}

	maxAgeString := os.Getenv(logMaxAgeEnvName)
	maxAge, err := strconv.Atoi(maxAgeString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log max age")
	}

	logLevel := os.Getenv(logLevelEnvName)
	if len(logLevel) == 0 {
		logLevel = defaultLogLevel
	}

	return &zapLoggerConfig{
		fileName:    fileName,
		fileMaxSize: fileMaxSize,
		maxBackups:  maxBackups,
		maxAge:      maxAge,
		logLevel:    logLevel,
	}, nil
}

func (cfg *zapLoggerConfig) FileName() string {
	return cfg.fileName
}

func (cfg *zapLoggerConfig) FileMaxSize() int {
	return cfg.fileMaxSize
}

func (cfg *zapLoggerConfig) MaxBackups() int {
	return cfg.maxBackups
}

func (cfg *zapLoggerConfig) MaxAge() int {
	return cfg.maxAge
}

func (cfg *zapLoggerConfig) LogLevel() string {
	return cfg.logLevel
}
