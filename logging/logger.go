package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use ISO8601 timestamp format

	var err error
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	defer Logger.Sync()
}
