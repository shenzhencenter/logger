package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	i    *zap.Logger
	host zap.Field
)

func init() {
	host = zap.String("_host", os.Getenv("HOSTNAME"))
	var config = zap.NewDevelopmentConfig()
	if os.Getenv("ENV") == "production" {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	i, _ = config.Build()
}

func WithHook(hook ...func(zapcore.Entry) error) {
	i = i.WithOptions(zap.Hooks(hook...))
}

func WithOptions(opts ...zap.Option) {
	i = i.WithOptions(opts...)
}

func Info(msg string, fields ...zap.Field) {
	i.With(host).Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	i.With(host).Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	i.With(host).Fatal(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	i.With(host).Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	i.With(host).Warn(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	i.With(host).DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	i.With(host).Panic(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return i.With(host).With(fields...)
}

func Sync() {
	i.Sync()
}
