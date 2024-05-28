package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

// Init - инициализация логгера
func Init(core zapcore.Core, options ...zap.Option) {
	globalLogger = zap.New(core, options...)
}

// Debug - логирование отладочной информации
func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

// Info - логирование информации
func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

// Warn - логирование предупреждения
func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}

// Error - логирование ошибки
func Error(msg string, fields ...zap.Field) {
	globalLogger.Error(msg, fields...)
}

// Fatal - логирование ошибки и завершение работы приложения
func Fatal(msg string, fields ...zap.Field) {
	globalLogger.Fatal(msg, fields...)
}

// WithOptions - создать новый логгер с опциями
func WithOptions(opts ...zap.Option) *zap.Logger {
	return globalLogger.WithOptions(opts...)
}
