package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log, _ = zap.NewProduction(zap.AddCallerSkip(1))

func SetOptions(opts ...zap.Option) {
	log = log.WithOptions(opts...)
}

func GetLogger() *zap.Logger {
	return log
}

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return log.Check(lvl, msg)
}

func Core() zapcore.Core {
	return log.Core()
}

func Sync() error {
	return log.Sync()
}

func prepareLog(ctx context.Context) *zap.Logger {
	return loggerFromContext(ctx)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func DebugCtx(ctx context.Context, msg string, fields ...zap.Field) {
	prepareLog(ctx).Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func CtxInfo(ctx context.Context, msg string, fields ...zap.Field) {
	prepareLog(ctx).Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

func CtxWarn(ctx context.Context, msg string, fields ...zap.Field) {
	prepareLog(ctx).Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func CtxError(ctx context.Context, msg string, fields ...zap.Field) {
	prepareLog(ctx).Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	log.DPanic(msg, fields...)
}

func CtxDPanic(ctx context.Context, msg string, fields ...zap.Field) {
	prepareLog(ctx).DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	log.Panic(msg, fields...)
}

func CtxPanic(ctx context.Context, msg string, fields ...zap.Field) {
	prepareLog(ctx).Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

func CtxFatal(ctx context.Context, msg string, fields ...zap.Field) {
	prepareLog(ctx).Fatal(msg, fields...)
}

func Named(s string) *zap.Logger {
	return log.Named(s)
}

func With(ctx context.Context, fields ...zap.Field) *zap.Logger {
	return log.With(fields...)
}

func CtxNamed(ctx context.Context, name string) context.Context {
	return newContextWithLogger(ctx, loggerFromContext(ctx).Named(name))
}

func CtxWith(ctx context.Context, fields ...zap.Field) context.Context {
	return newContextWithLogger(ctx, loggerFromContext(ctx).With(fields...))
}

type keyContextType uint8

const (
	loggerContextKey keyContextType = 1
)

func newContextWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

func loggerFromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(loggerContextKey).(*zap.Logger); ok {
		return l
	}
	return log
}
