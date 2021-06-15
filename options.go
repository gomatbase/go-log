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

type Options struct {
	dateFlags        int
	writer           io.Writer
	failingCriticals bool
	startingLevel    uint
}

func WithoutOptions() *Options {
	return WithOptions()
}

func WithOptions() *Options {
	return &Options{
		dateFlags:        0,
		writer:           os.Stdout,
		failingCriticals: false,
		startingLevel:    WARNING,
	}
}

func (o *Options) DateFlags(flags int) *Options {
	o.dateFlags = flags
	return o
}
func (o *Options) WithWriter(writer io.Writer) *Options {
	o.writer = writer
	return o
}

func (o *Options) WithFailingCriticals() *Options {
	o.failingCriticals = true
	return o
}

func (o *Options) WithoutFailingCriticals() *Options {
	o.failingCriticals = false
	return o
}

func (o *Options) WithStartingLevel(startingLevel uint) *Options {
	o.startingLevel = startingLevel
	return o
}

func (o *Options) equals(options *Options) bool {
	return o.failingCriticals == options.failingCriticals && o.dateFlags == options.dateFlags && o.startingLevel == options.startingLevel && o.writer == options.writer
}
