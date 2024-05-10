package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var i *zap.Logger

func init() {
	var (
		host   string     = os.Getenv("HOSTNAME")
		level  string     = os.Getenv("LOG_LEVEL")
		config zap.Config = zap.NewProductionConfig()
	)
	if len(level) > 0 {
		l, err := zap.ParseAtomicLevel(level)
		if err != nil {
			panic("invalid log level")
		}
		config.Level = l
	}
	config.Encoding = "json"
	config.Development = false
	config.Sampling = nil
	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	config.InitialFields = map[string]interface{}{"_host": host}
	i, _ = config.Build(zap.WithCaller(true), zap.AddCallerSkip(1))
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

func I() *zap.Logger {
	return i
}

func Sync() {
	i.Sync()
}
