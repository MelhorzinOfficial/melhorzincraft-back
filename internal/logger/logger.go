package logger

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/logger"
	"go.uber.org/zap"
)

type Config struct {
	Debug bool `mapstructure:"debug"`
}

func New(cfg *Config) logger.Logger {
	var l *zap.Logger

	if cfg.Debug {
		l, _ = zap.NewDevelopment()
	} else {
		l, _ = zap.NewProduction()
	}

	sugar := l.Sugar()
	return &log{
		l: sugar,
	}
}

type log struct {
	l *zap.SugaredLogger
}

func (l *log) Info(msg string, keysAndValues ...interface{}) {
	l.l.Infow(msg, keysAndValues...)
}

func (l *log) Error(msg string, keysAndValues ...interface{}) {
	l.l.Errorw(msg, keysAndValues...)
}

func (l *log) Warn(msg string, keysAndValues ...interface{}) {
	l.l.Warnw(msg, keysAndValues...)
}

func (l *log) Debug(msg string, keysAndValues ...interface{}) {
	l.l.Debugw(msg, keysAndValues...)
}
