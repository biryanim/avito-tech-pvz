package config

import (
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	loggerLevelEnvKey = "LOGGER_LEVEL"
)

type LoggerConfig interface {
	GetCore() zapcore.Core
}

type loggerConfig struct {
	level zapcore.Level
}

func NewLoggerConfig() (LoggerConfig, error) {
	loglevel := os.Getenv(loggerLevelEnvKey)
	if len(loglevel) == 0 {
		loglevel = "info"
	}

	var level zapcore.Level
	if err := level.Set(loglevel); err != nil {
		return nil, err
	}

	return &loggerConfig{
		level: level,
	}, nil
}

func (l *loggerConfig) GetCore() zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Цветные уровни
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // Формат времени
		EncodeDuration: zapcore.StringDurationEncoder,    // Длительность как строка
		EncodeCaller:   zapcore.ShortCallerEncoder,       // Короткий формат caller'а
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	return zapcore.NewCore(
		consoleEncoder,
		stdout,
		l.level,
	)
}
