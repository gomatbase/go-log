// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"io"
	l "log"
	"os"
)

const (
	// log pattern variables copied from standard log package for import easyness

	Ldate         = l.Ldate         // the date in the local time zone: 2009/01/23
	Ltime         = l.Ltime         // the time in the local time zone: 01:23:23
	Lmicroseconds = l.Lmicroseconds // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile     = l.Llongfile     // full file name and line number: /a/b/c/d.go:23
	Lshortfile    = l.Lshortfile    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC          = l.LUTC          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix    = l.Lmsgprefix    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = l.LstdFlags     // initial values for the standard logger
)

// Options holds the configuration for a new logger and provides methods to setup the configurable options
type Options struct {
	dateFlags        int       // format flags for the logger as per the go standard log package
	writer           io.Writer // writer where the logs should be logged
	failingCriticals bool      // flag setting if a critical log should result in a fatal entry (process exits)
	startingLevel    uint      // the log level the logger should start in
}

// WithoutOptions is syntax-candy returning an Options object with default settings. When used, it implies that the
// options object is not meant to be configured, but there is no restriction in doing so.
func WithoutOptions() *Options {
	return WithOptions()
}

// WithOptions is syntax-candy to create an Options object with default settings
func WithOptions() *Options {
	return &Options{
		dateFlags:        0,
		writer:           os.Stdout,
		failingCriticals: false,
		startingLevel:    WARNING,
	}
}

// DateFlags sets the format flags for the logger
func (o *Options) DateFlags(flags int) *Options {
	o.dateFlags = flags
	return o
}

// WithWriter defines the writer the logger should use
func (o *Options) WithWriter(writer io.Writer) *Options {
	o.writer = writer
	return o
}

// WithFailingCriticals sets the logger to fail (exit process) when logging a critical
func (o *Options) WithFailingCriticals() *Options {
	o.failingCriticals = true
	return o
}

// WithoutFailingCriticals sets the logger to log criticals as plain log entries (process doesn't break)
func (o *Options) WithoutFailingCriticals() *Options {
	o.failingCriticals = false
	return o
}

// WithStartingLevel sets the initial log level the logger has
func (o *Options) WithStartingLevel(startingLevel uint) *Options {
	o.startingLevel = startingLevel
	return o
}

// equals compares if the options object is an exact match to another options object
func (o *Options) equals(options *Options) bool {
	return o.failingCriticals == options.failingCriticals && o.dateFlags == options.dateFlags && o.startingLevel == options.startingLevel && o.writer == options.writer
}
