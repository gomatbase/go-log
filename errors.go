package log

type logError string

func (le logError) Error() string {
	return string(le)
}

const (
	errEmptyLoggerName              = logError("Logger name may not be empty")
	errReinitializingExistingLogger = logError("trying to initialize an already initialized logger with different options")
)
