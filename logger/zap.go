package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() *ZapLogger {
	config := zap.NewDevelopmentConfig()
	enc := zapcore.NewJSONEncoder(config.EncoderConfig)

	wsfile := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   LOG_PATH,
			MaxSize:    LOG_MAX_SIZE_MB,
			MaxBackups: LOG_MAX_BACKUPS,
			MaxAge:     LOG_MAX_AGE,
		},
	)
	wsstdout := zapcore.AddSync(
		zapcore.Lock(os.Stdout),
	)

	logger := zap.New(
		zapcore.NewTee(
			zapcore.NewCore(enc, wsfile, config.Level),
			zapcore.NewCore(enc, wsstdout, config.Level),
		),
	)
	defer logger.Sync()

	return &ZapLogger{
		logger: logger,
	}
}

func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	l.logger.Log(zap.PanicLevel, fmt.Sprintf(format, args...))
}

func (l *ZapLogger) Panic(v ...interface{}) {
	l.logger.Log(zap.PanicLevel, fmt.Sprint(v...))
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Log(zap.FatalLevel, fmt.Sprintf(format, args...))
}

func (l *ZapLogger) Fatal(v ...interface{}) {
	l.logger.Log(zap.FatalLevel, fmt.Sprint(v...))
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Log(zap.ErrorLevel, fmt.Sprintf(format, args...))
}

func (l *ZapLogger) Error(v ...interface{}) {
	l.logger.Log(zap.ErrorLevel, fmt.Sprint(v...))
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.logger.Log(zap.WarnLevel, fmt.Sprintf(format, args...))
}

func (l *ZapLogger) Warn(v ...interface{}) {
	l.logger.Log(zap.WarnLevel, fmt.Sprint(v...))
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.logger.Log(zap.InfoLevel, fmt.Sprintf(format, args...))
}

func (l *ZapLogger) Info(v ...interface{}) {
	l.logger.Log(zap.InfoLevel, fmt.Sprint(v...))
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.logger.Log(zap.DebugLevel, fmt.Sprintf(format, args...))
}

func (l *ZapLogger) Debug(v ...interface{}) {
	l.logger.Log(zap.DebugLevel, fmt.Sprint(v...))
}
