package log

// Simple error type for log errors
type logError string

func (le logError) Error() string {
	return string(le)
}

const (

	// ErrEmptyLoggerName Error raised when trying to refer to a logger with an empty name
	ErrEmptyLoggerName = logError("Logger name may not be empty")

	// ErrReinitializingExistingLogger Error raised when trying to initialize an existing logger with different options
	ErrReinitializingExistingLogger = logError("trying to initialize an already initialized logger with different options")

	// ErrLoggerDoesNotExist Error raised when referring to a non-existing logger
	ErrLoggerDoesNotExist = logError("logger with given name doesn't exist")
)
