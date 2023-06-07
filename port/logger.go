package port

import "go.uber.org/zap"

type Logger interface {
	Info(string, ...zap.Field)
	Debug(string, ...zap.Field)
}
