// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"io"
)

// constants for date format. Borrowing the same names from standard log package
const (
	Ldate         = iota // the date in the local time zone: 2009/01/23
	Ltime                // the time in the local time zone: 01:23:23
	Lmicroseconds        // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	LUTC                 // if Ldate or Ltime is set, use UTC rather than the local time zone
)

// Log message format constants to use in message header format pattern
const (
	Time = iota
	Name
	Source
	LongSource
	Separator
	LogLevel
)

// Types of loggers
const (
	standard = iota
	syncedAppender
)

// Options represents a logger options object and methods available to configure it
type Options interface {
	// DateFlags sets the format flags for the logger
	DateFlags(flags int) Options

	// WithFailingCriticals sets the logger to fail (exit process) when logging a critical
	WithFailingCriticals() Options

	// WithoutFailingCriticals sets the logger to log criticals as plain log entries (process doesn't break)
	WithoutFailingCriticals() Options

	// WithStartingLevel sets the initial log level the logger has
	WithStartingLevel(startingLevel severity) Options

	// WithLevelLogPrefix sets the log prefix format for a specific level
	WithLevelLogPrefix(logLevel severity, flags ...uint) Options

	// WithLogPrefix sets the log prefix format for all levels
	WithLogPrefix(flags ...uint) Options
}

type StandardWriter interface {
	Options
	WithWriter(writer io.Writer) Options
}

type AppendersLogger interface {
	Options
}

// options holds the configuration for a new logger and provides methods to setup the configurable options
type options struct {
	loggerType       uint      // type of logger the options are for
	dateFlags        int       // format flags for the logger as per the go standard log package
	failingCriticals bool      // flag setting if a critical log should result in a fatal entry (process exits)
	startingLevel    severity  // the log level the logger should start in
	levelFormats     [][]uint  // formats used for each of the log levels
	writer           io.Writer // writer that should be used for a standard writer logger
}

// Standard creates an Options object for standard logging
func Standard() StandardWriter {
	return &options{
		loggerType:       standard,
		dateFlags:        0,
		failingCriticals: false,
		startingLevel:    WARNING,
		levelFormats:     make([][]uint, TRACE+1),
	}
}

func SyncedAppenders() AppendersLogger {
	return &options{
		loggerType:       syncedAppender,
		failingCriticals: false,
		startingLevel:    WARNING,
	}
}

// WithWriter sets the writer for a StandardWriter logger
func (o *options) WithWriter(writer io.Writer) Options {
	o.writer = writer
	return o
}

// DateFlags sets the format flags for the logger
func (o *options) DateFlags(flags int) Options {
	o.dateFlags = flags
	return o
}

// WithFailingCriticals sets the logger to fail (exit process) when logging a critical
func (o *options) WithFailingCriticals() Options {
	o.failingCriticals = true
	return o
}

// WithoutFailingCriticals sets the logger to log criticals as plain log entries (process doesn't break)
func (o *options) WithoutFailingCriticals() Options {
	o.failingCriticals = false
	return o
}

// WithStartingLevel sets the initial log level the logger has
func (o *options) WithStartingLevel(startingLevel severity) Options {
	o.startingLevel = startingLevel
	return o
}

func validatePrefixFlags(flags []uint) {
	foundFlags := []bool{false, false, false, false, false}
	for _, flag := range flags {
		if flag > LogLevel {
			panic("Unknown lof prefix flag")
		} else if foundFlags[flag] {
			panic("duplicating  prefix flags")
		}
		foundFlags[flag] = true
	}
}

// WithLevelLogPrefix sets the log prefix format for a specific level
func (o *options) WithLevelLogPrefix(logLevel severity, flags ...uint) Options {
	validatePrefixFlags(flags)
	if int(logLevel) >= len(o.levelFormats) {
		for i := len(o.levelFormats); i <= int(logLevel); i++ {
			o.levelFormats = append(o.levelFormats, nil)
		}
	}
	o.levelFormats[logLevel] = flags
	return o
}

// WithLogPrefix sets the log prefix format for all levels
func (o *options) WithLogPrefix(flags ...uint) Options {
	validatePrefixFlags(flags)
	for i, _ := range o.levelFormats {
		o.levelFormats[i] = flags
	}
	return o
}

// equals compares if the options object is an exact match to another options object
func (o *options) equals(options *options) bool {
	return o.failingCriticals == options.failingCriticals && o.dateFlags == options.dateFlags && o.startingLevel == options.startingLevel
}
