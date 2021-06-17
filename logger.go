// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	l "log"
)

// logger Simple implementation of a Logger, using the standard go log as the actual logging framework
type logger struct {
	options   *Options                     // the original options used to create the logger
	level     uint                         // the current log level
	name      string                       // the name of the logger (also the prefix for logging)
	log       *l.Logger                    // standard fo log.Logger to actually log the entries
	critical  func(...interface{})         // function used to log criticals
	criticalf func(string, ...interface{}) // function used to log criticals with a format string
}

// SetLevel sets the level of the logger
func (logger *logger) SetLevel(level uint) {
	logger.level = level
}

func (logger *logger) Level() uint {
	return logger.level
}

func (logger *logger) Critical(v ...interface{}) {
	logger.critical(v...)
}

func (logger *logger) Criticalf(format string, v ...interface{}) {
	logger.criticalf(format, v...)
}

func (logger *logger) Error(v ...interface{}) {
	logger.Println(ERROR, v...)
}

func (logger *logger) Errorf(format string, v ...interface{}) {
	logger.Printf(ERROR, format, v...)
}

func (logger *logger) Warning(v ...interface{}) {
	logger.Println(WARNING, v...)
}

func (logger *logger) Warningf(format string, v ...interface{}) {
	logger.Printf(WARNING, format, v...)
}

func (logger *logger) Info(v ...interface{}) {
	logger.Println(INFO, v...)
}

func (logger *logger) Infof(format string, v ...interface{}) {
	logger.Printf(INFO, format, v...)
}

func (logger *logger) Debug(v ...interface{}) {
	logger.Println(DEBUG, v...)
}

func (logger *logger) Debugf(format string, v ...interface{}) {
	logger.Printf(DEBUG, format, v...)
}

func (logger *logger) Trace(v ...interface{}) {
	logger.Println(TRACE, v...)
}

func (logger *logger) Tracef(format string, v ...interface{}) {
	logger.Printf(TRACE, format, v...)
}

func (logger *logger) Println(level uint, v ...interface{}) {
	if level <= logger.level {
		logger.log.Println(v...)
	}
}

func (logger *logger) Printf(level uint, format string, v ...interface{}) {
	if level <= logger.level {
		logger.log.Printf(format, v...)
	}
}
