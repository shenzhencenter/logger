package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var i *zap.Logger

func init() {
	var (
		env    string     = os.Getenv("ENV")
		host   string     = os.Getenv("HOSTNAME")
		config zap.Config = zap.NewDevelopmentConfig()
	)
	if env == "production" || env == "prod" || env == "prd" {
		config = zap.NewProductionConfig()
	}
	config.Encoding = "json"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	config.InitialFields = map[string]interface{}{"_host": host}
	i, _ = config.Build()
}

func WithHook(hook ...func(zapcore.Entry) error) {
	i = i.WithOptions(zap.Hooks(hook...))
}

func WithOptions(opts ...zap.Option) {
	i = i.WithOptions(opts...)
}

func Info(msg string, fields ...zap.Field) {
	i.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	i.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	i.Fatal(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	i.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	i.Warn(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	i.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	i.Panic(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return i.With(fields...)
}

func Sync() {
	i.Sync()
}
