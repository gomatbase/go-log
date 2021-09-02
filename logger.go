// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"fmt"
	l "log"
	"os"
)

// logger Simple implementation of a Logger, using the standard go log as the actual logging framework
type logger struct {
	options         *options  // the original options used to create the logger
	level           uint      // the current log level
	name            string    // the name of the logger (also the prefix for logging)
	log             *l.Logger // standard fo log.Logger to actually log the entries
	criticalFailure bool
}

// SetLevel sets the level of the logger
func (logger *logger) SetLevel(level uint) {
	logger.level = level
}

func (logger *logger) Level() uint {
	return logger.level
}

func (logger *logger) Critical(v ...interface{}) {
	logger.log.Output(2, fmt.Sprintln(v...))
	if logger.criticalFailure {
		os.Exit(1)
	}
}

func (logger *logger) Criticalf(format string, v ...interface{}) {
	logger.log.Output(2, fmt.Sprintf(format, v...))
	if logger.criticalFailure {
		os.Exit(1)
	}
}

func (logger *logger) Error(v ...interface{}) {
	logger.println(ERROR, v...)
}

func (logger *logger) Errorf(format string, v ...interface{}) {
	logger.printf(ERROR, format, v...)
}

func (logger *logger) Warning(v ...interface{}) {
	logger.println(WARNING, v...)
}

func (logger *logger) Warningf(format string, v ...interface{}) {
	logger.printf(WARNING, format, v...)
}

func (logger *logger) Info(v ...interface{}) {
	logger.println(INFO, v...)
}

func (logger *logger) Infof(format string, v ...interface{}) {
	logger.printf(INFO, format, v...)
}

func (logger *logger) Debug(v ...interface{}) {
	logger.println(DEBUG, v...)
}

func (logger *logger) Debugf(format string, v ...interface{}) {
	logger.printf(DEBUG, format, v...)
}

func (logger *logger) Trace(v ...interface{}) {
	logger.println(TRACE, v...)
}

func (logger *logger) Tracef(format string, v ...interface{}) {
	logger.printf(TRACE, format, v...)
}

func (logger *logger) Println(level uint, v ...interface{}) {
	logger.println(level, v...)
}

func (logger *logger) Printf(level uint, format string, v ...interface{}) {
	logger.printf(level, format, v...)
}

func (logger *logger) println(level uint, v ...interface{}) {
	if level <= logger.level {
		logger.log.Output(3, fmt.Sprintln(v...))
	}
}

func (logger *logger) printf(level uint, format string, v ...interface{}) {
	if level <= logger.level {
		logger.log.Output(3, fmt.Sprintf(format, v...))
	}
}
