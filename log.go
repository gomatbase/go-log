// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"errors"
	l "log"
	"sync"
)

const (
	CRITICAL = iota // Critical log level (Always logged)
	ERROR           // Error log level
	WARNING         // Warning log level (Default level)
	INFO            // Info log level
	DEBUG           // Debug log level
	TRACE           // Trace log level
)

// DEFAULT is the name of the default logger
const DEFAULT = ""

// dictionary to translate the log level to it's name (for the default severity scale)
var levelNames = []string{"ERROR", "WARNING", "INFO", "DEBUG", "TRACE"}

var (
	loggers          = make(map[string]*logger)                  // map of all existing loggers. Indexed by their names.
	lock             = sync.Mutex{}                              // mutex to manipulate the loggers map
	defaultLogger, _ = getWithOptions(DEFAULT, WithoutOptions()) // the default logger provided by the package for out-of-the-box usage with default options.
)

// Logger defines the interface a Logger implementation must provide
type Logger interface {

	// SetLevel sets the current log level of the logger
	SetLevel(uint)

	// Level returns the current log level of the logger
	Level() uint

	// Critical logs the message(s) at the critical level
	Critical(v ...interface{})

	// Criticalf logs the formated message at the critical level
	Criticalf(format string, v ...interface{})

	// Error logs the message(s) at the error level
	Error(v ...interface{})

	// Errorf logs the formated message at the error level
	Errorf(format string, v ...interface{})

	// Warning logs the message(s) at the warning level
	Warning(v ...interface{})

	// Warningf logs the formated message at the warning level
	Warningf(format string, v ...interface{})

	// Info logs the message(s) at the info level
	Info(v ...interface{})

	// Infof logs the formated message at the info level
	Infof(format string, v ...interface{})

	// Debug logs the message(s) at the debug level
	Debug(v ...interface{})

	// Debugf logs the formated message at the debug level
	Debugf(format string, v ...interface{})

	// Trace logs the message(s) at the trace level
	Trace(v ...interface{})

	// Tracef logs the formated message at the trace level
	Tracef(format string, v ...interface{})

	// Println logs the message(s) at the provided level
	Println(level uint, v ...interface{})

	// Printf logs the formated message at the provided level
	Printf(level uint, format string, v ...interface{})
}

// Get will create or get an existing logger with the given name. If the logger doesn't exist it will be created with
// the default options (warning level, logs to stdout and non-failing criticals). The name must be a non-empty string
// (may be spaces).
func Get(name string) (Logger, error) {
	logger, e := GetWithOptions(name, WithoutOptions())
	var returnError error
	if name != DEFAULT && e != nil && e != errReinitializingExistingLogger {
		returnError = e
	}
	return logger, returnError
}

// GetWithOptions will create a log with the provided options if it doesn't exist yet or returns an existing log if
// the provided options are the same as the options the existing logger was created with. Trying to get an existing
// logger with different options. The name logger may not be an empty string (can be filled spaces).
func GetWithOptions(name string, options *Options) (Logger, error) {
	if len(name) == 0 {
		return nil, errEmptyLoggerName
	}
	return getWithOptions(name, options)
}

// private getWithOptions function that actually fetches or creates the logger. The internal version allows an empty
// string as a name allowing the creation of the DEFAULT logger.
func getWithOptions(name string, options *Options) (*logger, error) {
	lock.Lock()
	defer lock.Unlock()

	logger, found := loggers[name]
	var e error
	if !found {
		loggers[name] = newLogger(name, options)
		logger = loggers[name]
	} else if !logger.options.equals(options) {
		e = errReinitializingExistingLogger
	}

	return logger, e
}

// OverrideLogWithOptions allows resetting the options of an existing logger. The function is provided for convenience
// and allows changing the options of the default logger. It is not thread-safe and aside the use to set the options
// for the default logger it shouldn't be used
func OverrideLogWithOptions(name string, options *Options) (Logger, error) {
	lock.Lock()
	defer lock.Unlock()

	logger, found := loggers[name]
	if !found {
		return nil, errors.New("log doesn't exist")
	} else if !logger.options.equals(options) {
		logger.log = l.New(options.writer, logger.name+" - ", options.dateFlags)
		logger.level = options.startingLevel
		if options.failingCriticals {
			logger.critical = logger.log.Fatalln
			logger.criticalf = logger.log.Fatalf
		} else {
			logger.critical = logger.log.Println
			logger.criticalf = logger.log.Printf
		}
		logger.options = options
	}

	return logger, nil
}

// newLogger creates a new logger with the given name from the provided options
func newLogger(name string, options *Options) *logger {
	o := options
	if o == nil {
		o = WithOptions()
	}

	prefix := name
	if len(prefix) > 0 {
		prefix = prefix + " - "
	}
	logger := logger{
		options: o,
		level:   o.startingLevel,
		name:    name,
		log:     l.New(o.writer, prefix, o.dateFlags),
	}

	if o.failingCriticals {
		logger.critical = logger.log.Fatalln
		logger.criticalf = logger.log.Fatalf
	} else {
		logger.critical = logger.log.Println
		logger.criticalf = logger.log.Printf
	}

	return &logger
}

// SetLevel sets the log level of the default logger
func SetLevel(level uint) {
	defaultLogger.SetLevel(level)
}

// SetLoggerLevel sets the log level of a logger by name. DEFAULT may be used to set the default logger level.
func SetLoggerLevel(name string, level uint) error {
	lock.Lock()
	logger, found := loggers[name]
	lock.Unlock()

	if !found {
		return errLoggerDoesNotExist
	}

	logger.SetLevel(level)
	return nil
}

// Level returns the current log level of the default logger
func Level() uint {
	return defaultLogger.Level()
}

// LevelName is a convenience method to translate the log level into a name. It only works for loggers implementing
// the default severity scale.
func LevelName(level uint) string {
	if int(level) >= len(levelNames) {
		return "UNKNOWN"
	}
	return levelNames[level]
}

// Critical logs a critical log entry through the default logger
func Critical(v ...interface{}) {
	defaultLogger.Critical(v...)
}

// Criticalf logs a formatted critical log entry through the default logger
func Criticalf(format string, v ...interface{}) {
	defaultLogger.Criticalf(format, v...)
}

// Error logs a error log entry through the default logger
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

// Errorf logs a formatted error log entry through the default logger
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

// Warning logs a warning log entry through the default logger
func Warning(v ...interface{}) {
	defaultLogger.Println(WARNING, v...)
}

// Warningf logs a formatted warning log entry through the default logger
func Warningf(format string, v ...interface{}) {
	defaultLogger.Printf(WARNING, format, v...)
}

// Info logs a info log entry through the default logger
func Info(v ...interface{}) {
	defaultLogger.Println(INFO, v...)
}

// Infof logs a formatted info log entry through the default logger
func Infof(format string, v ...interface{}) {
	defaultLogger.Printf(INFO, format, v...)
}

// Debug logs a debug log entry through the default logger
func Debug(v ...interface{}) {
	defaultLogger.Println(DEBUG, v...)
}

// Debugf logs a formatted debug log entry through the default logger
func Debugf(format string, v ...interface{}) {
	defaultLogger.Printf(DEBUG, format, v...)
}

// Trace logs a trace log entry through the default logger
func Trace(v ...interface{}) {
	defaultLogger.Println(TRACE, v...)
}

// Tracef logs a formatted trace log entry through the default logger
func Tracef(format string, v ...interface{}) {
	defaultLogger.Printf(TRACE, format, v...)
}

// Println logs a log entry at the given log level  through the default logger
func Println(level uint, v ...interface{}) {
	defaultLogger.Println(level, v...)
}

// Printf logs a formatted log entry at the given log level through the default logger
func Printf(level uint, format string, v ...interface{}) {
	defaultLogger.Printf(level, format, v...)
}
