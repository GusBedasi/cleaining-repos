package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

var logger *zap.Logger

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warning(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func init() {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	zapLogger := zap.Must(cfg.Build())
	defer zapLogger.Sync()

	logger = zapLogger
}
