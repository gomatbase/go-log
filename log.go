// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"sync"
)

// Log Severity levels
const (
	CRITICAL = iota // Critical log level (Always logged)
	ERROR           // Error log level
	WARNING         // Warning log level (Default level)
	INFO            // Info log level
	DEBUG           // Debug log level
	TRACE           // Trace log level

	UNKNOWN = -1 // used as error log type for unknown levels
)

// DEFAULT is the name of the default logger
const DEFAULT = "DEFAULT"

var (
	// dictionary to translate the log level to it's name (for the default severity scale)
	levelNames = []string{"CRITICAL", "ERROR", "WARNING", "INFO", "DEBUG", "TRACE"}

	// dictionary to convert a level name to severity
	levelSeverities = map[string]int{
		levelNames[CRITICAL]: CRITICAL,
		levelNames[ERROR]:    ERROR,
		levelNames[WARNING]:  WARNING,
		levelNames[INFO]:     INFO,
		levelNames[DEBUG]:    DEBUG,
		levelNames[TRACE]:    TRACE,
	}
)

var (
	loggers          = make(map[string]Logger)             // map of all existing loggers. Indexed by their names.
	lock             = sync.Mutex{}                        // mutex to manipulate the loggers map
	defaultLogger, _ = getWithOptions(DEFAULT, Standard()) // the default logger provided by the package for out-of-the-box usage with default options.
)

// Get will create or get an existing logger with the given name. If the logger doesn't exist it will be created with
// the default options (warning level, logs to stdout and non-failing criticals). The name must be a non-empty string
// (may be spaces).
func Get(name string) (Logger, error) {
	logger, e := GetWithOptions(name, Standard())
	var returnError error
	if e != nil && e != ErrReinitializingExistingLogger {
		returnError = e
	}
	return logger, returnError
}

// GetWithOptions will create a log with the provided options if it doesn't exist yet or returns an existing log if
// the provided options are the same as the options the existing logger was created with. Trying to get an existing
// logger with different options. The name logger may not be an empty string (can be filled spaces).
func GetWithOptions(name string, options Options) (Logger, error) {
	if len(name) == 0 {
		return nil, ErrEmptyLoggerName
	}
	return getWithOptions(name, options)
}

// private getWithOptions function that actually fetches or creates the logger. The internal version allows an empty
// string as a name allowing the creation of the DEFAULT logger.
func getWithOptions(name string, o Options) (Logger, error) {
	lock.Lock()
	defer lock.Unlock()

	if logger, found := loggers[name]; !found {
		if o == nil {
			o = Standard()
		}
		if logger, e := newLogger(name, o.(*options)); e != nil {
			return nil, e
		} else {
			loggers[name] = logger
			return logger, nil
		}
	} else {
		return logger, nil
	}
}

// SetDefaultLogger allows overriding the default logger with different options
func SetDefaultLogger(o Options) error {
	lock.Lock()
	defer lock.Unlock()

	if logger, e := newLogger(DEFAULT, o.(*options)); e != nil {
		return e
	} else {
		defaultLogger = logger
		loggers[DEFAULT] = logger
	}

	return nil
}

// newLogger creates a new logger with the given name from the provided options
func newLogger(name string, o *options) (Logger, error) {
	switch o.loggerType {
	case standard:
		return newStandardLogger(name, o), nil
	case syncedAppender:
		return newStandardLogger(name, o), nil
	default:
		return nil, ErrUnknownLoggerType
	}
}

// SetLevel sets the log level of the default logger
func SetLevel(level int) {
	defaultLogger.SetLevel(level)
}

// SetLoggerLevel sets the log level of a logger by name. DEFAULT may be used to set the default logger level.
func SetLoggerLevel(name string, level int) error {
	lock.Lock()
	logger, found := loggers[name]
	lock.Unlock()

	if !found {
		return ErrLoggerDoesNotExist
	}

	logger.SetLevel(level)
	return nil
}

// SetLoggerLevels sets the log levels of several loggers at once. If any logger is not found it will be omitted from the response
func SetLoggerLevels(loggerLevels map[string]int) map[string]int {
	result := make(map[string]int)

	for k, l := range loggerLevels {
		if e := SetLoggerLevel(k, l); e == nil {
			result[k] = l
		}
	}

	return result
}

// Level returns the current log level of the default logger
func Level() int {
	return defaultLogger.Level()
}

// LoggerLevels gets the current log levels of all known loggers
func LoggerLevels() map[string]int {
	loggerLevels := make(map[string]int)
	lock.Lock()
	for k, l := range loggers {
		loggerLevels[k] = l.Level()
	}
	lock.Unlock()
	return loggerLevels
}

// LoggerLevel gets the current log level of the logger with the given name. ErrLoggerDoesNotExist is returned as an
// error if a logger with the given name doesn't is unknown.
func LoggerLevel(name string) (int, error) {
	lock.Lock()
	logger, found := loggers[name]
	lock.Unlock()

	if !found {
		return 0, ErrLoggerDoesNotExist
	}

	return logger.Level(), nil
}

// LoggerLevelName gets the current log level name of the logger with the given name. ErrLoggerDoesNotExist is returned as an
// error if a logger with the given name doesn't is unknown. Utility method for loggers using the standard severity scale.
func LoggerLevelName(name string) (string, error) {
	if level, e := LoggerLevel(name); e == nil {
		return LevelName(level), nil
	} else {
		return "", e
	}
}

// LoggerLevelNames gets the current log level names of all known loggers. Only relevant if logger is using the standard severity scale.
func LoggerLevelNames() map[string]string {
	loggerLevels := LoggerLevels()
	loggerLevelNames := make(map[string]string)
	for k, l := range loggerLevels {
		loggerLevelNames[k] = LevelName(l)
	}
	return loggerLevelNames
}

// LevelName is a convenience method to translate the log level into a name. It only works for loggers implementing
// the default severity scale.
func LevelName(level int) string {
	if int(level) >= len(levelNames) {
		return "UNKNOWN"
	}
	return levelNames[level]
}

func LevelSeverity(name string) int {
	if s, found := levelSeverities[name]; found {
		return s
	}
	return UNKNOWN
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
	defaultLogger.Warning(v...)
}

// Warningf logs a formatted warning log entry through the default logger
func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

// Info logs a info log entry through the default logger
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Infof logs a formatted info log entry through the default logger
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Debug logs a debug log entry through the default logger
func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

// Debugf logs a formatted debug log entry through the default logger
func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

// Trace logs a trace log entry through the default logger
func Trace(v ...interface{}) {
	defaultLogger.Trace(v...)
}

// Tracef logs a formatted trace log entry through the default logger
func Tracef(format string, v ...interface{}) {
	defaultLogger.Tracef(format, v...)
}
