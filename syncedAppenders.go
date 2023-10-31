package log

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type syncedAppenders struct {
	options         *options // the original options used to create the logger
	level           int      // the current log level
	levelHasSource  []bool
	name            string
	criticalFailure bool

	appenders []Appender
	mutex     sync.Mutex
}

func newSyncedAppenders(o *options) *syncedAppenders {
	return &syncedAppenders{}
}

// SetLevel sets the current log level of the logger
func (sa *syncedAppenders) SetLevel(level int) {
	sa.level = level
}

// Level returns the current log level of the logger
func (sa *syncedAppenders) Level() int {
	return sa.level
}

// Critical logs the message(s) at the critical level
func (sa *syncedAppenders) Critical(v ...interface{}) {
	sa.println(CRITICAL, v...)
}

// Criticalf logs the formatted message at the critical level
func (sa *syncedAppenders) Criticalf(format string, v ...interface{}) {
	sa.printf(CRITICAL, format, v...)
}

// Error logs the message(s) at the error level
func (sa *syncedAppenders) Error(v ...interface{}) {
	sa.println(ERROR, v...)
}

// Errorf logs the formatted message at the error level
func (sa *syncedAppenders) Errorf(format string, v ...interface{}) {
	sa.printf(ERROR, format, v...)
}

// Warning logs the message(s) at the warning level
func (sa *syncedAppenders) Warning(v ...interface{}) {
	sa.println(WARNING, v...)
}

// Warningf logs the formatted message at the warning level
func (sa *syncedAppenders) Warningf(format string, v ...interface{}) {
	sa.printf(WARNING, format, v...)
}

// Info logs the message(s) at the info level
func (sa *syncedAppenders) Info(v ...interface{}) {
	sa.println(INFO, v...)
}

// Infof logs the formatted message at the info level
func (sa *syncedAppenders) Infof(format string, v ...interface{}) {
	sa.printf(INFO, format, v...)
}

// Debug logs the message(s) at the debug level
func (sa *syncedAppenders) Debug(v ...interface{}) {
	sa.println(DEBUG, v...)
}

// Debugf logs the formatted message at the debug level
func (sa *syncedAppenders) Debugf(format string, v ...interface{}) {
	sa.printf(DEBUG, format, v...)
}

// Trace logs the message(s) at the trace level
func (sa *syncedAppenders) Trace(v ...interface{}) {
	sa.println(TRACE, v...)
}

// Tracef logs the formatted message at the trace level
func (sa *syncedAppenders) Tracef(format string, v ...interface{}) {
	sa.printf(TRACE, format, v...)
}

// println logs the message(s) at the provided level
func (sa *syncedAppenders) println(level int, v ...interface{}) {
	if level <= sa.level {
		sa.output(level, fmt.Sprintln(v...))
	}
	if level == CRITICAL && sa.criticalFailure {
		panic("critical failure")
	}
}

// printf logs the formatted message at the provided level
func (sa *syncedAppenders) printf(level int, format string, v ...interface{}) {
	if level <= sa.level {
		sa.output(level, fmt.Sprintf(format, v...))
	}
	if level == CRITICAL && sa.criticalFailure {
		panic("critical failure")
	}
}

func (sa *syncedAppenders) output(level int, message string) {
	sa.mutex.Lock()
	defer sa.mutex.Unlock()

	entry := &LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
	}
	if sa.levelHasSource[level] {
		sa.mutex.Unlock()
		_, file, line, ok := runtime.Caller(3)
		if !ok {
			file = "???"
			line = 0
		}
		entry.Source = &file
		entry.Line = line
		sa.mutex.Lock()
	}

	for _, appender := range sa.appenders {
		appender.Print(entry)
	}
}
