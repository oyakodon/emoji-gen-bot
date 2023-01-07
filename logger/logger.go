package logger

const (
	LOG_PATH = "./log/latest.log"

	LOG_MAX_SIZE_MB = 512 // megabytes
	LOG_MAX_BACKUPS = 3
	LOG_MAX_AGE     = 28 //days
)

type Logger interface {
	Panicf(string, ...interface{})
	Panic(...interface{})

	Fatalf(string, ...interface{})
	Fatal(...interface{})

	Errorf(string, ...interface{})
	Error(...interface{})

	Warnf(string, ...interface{})
	Warn(...interface{})

	Infof(string, ...interface{})
	Info(...interface{})

	Debugf(string, ...interface{})
	Debug(...interface{})
}
