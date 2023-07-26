package logger

import (
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLog() *zap.SugaredLogger {
	json := os.Getenv("ZAP_JSON")
	jsonLogs := false
	jsonLogs, _ = strconv.ParseBool(json)
	if jsonLogs {
		return initZapProdLog().Sugar()
	}
	return InitZapDeveloperLog().Sugar()

}
func initZapProdLog() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}

func InitZapDeveloperLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}
