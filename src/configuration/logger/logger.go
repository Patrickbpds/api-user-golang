package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LEVEL_LOGS = "LEVEL_LOGS"
	OUTPUT_LOGS = "OUTPUT_LOGS"
)

func init() {
	logConfig := zap.Config{
		OutputPaths:      []string{getOutputLogs()},
		Level:            zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:         "json",
		EncoderConfig:    zapcore.EncoderConfig{
			LevelKey: 	 "level",
			TimeKey:      "time",
			MessageKey: "message",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfig.Build()  
}


func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(OUTPUT_LOGS)))
	if output == "" {
		return "stdout"
	}
	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LEVEL_LOGS))){
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.InfoLevel
	}
}