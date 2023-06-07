package adapter

import (
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() *Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		// Generally don't do this, but as this is basically
		// essential for everything else, just fail here.
		log.Fatalln(err.Error())
	}

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}
