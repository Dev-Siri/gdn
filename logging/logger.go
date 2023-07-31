package logging

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

func InitLogger() error {
	logger, err := zap.NewProduction()

	if err != nil {
		return err
	}

	zapLogger = logger

	return nil
}

func Log(level zapcore.Level, message string, status int64, cache string, requestCtx *fasthttp.RequestCtx) {
	zapLogger.Log(
		level, message,
		zapcore.Field{
			Key:    "method",
			Type:   zapcore.StringType,
			String: string(requestCtx.Method()),
		},
		zapcore.Field{
			Key:     "status",
			Type:    zapcore.Uint16Type,
			Integer: status,
		},
		zapcore.Field{
			Key:    "request_ip",
			Type:   zapcore.StringType,
			String: requestCtx.RemoteIP().String(),
		},
		zapcore.Field{
			Key:    "cache",
			Type:   zapcore.StringType,
			String: cache,
		},
	)
}

func FlushLogger() {
	zapLogger.Sync()
}
