// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	levelTokens = []string{"[CRT]", "[ERR]", "[WRN]", "[INF]", "[DBG]", "[TRC]"}
)

type headerFormat struct {
	hasSource bool
	format    []uint
}

// logger Simple implementation of a Logger, using the standard go log as the actual logging framework
type standardLogger struct {
	options         *options // the original options used to create the logger
	level           severity // the current log level
	name            string
	writer          io.Writer
	buffer          []byte
	levelFormats    []headerFormat
	mutex           sync.Mutex
	criticalFailure bool
}

func newStandardLogger(name string, options *options) Logger {
	levelFormats := make([]headerFormat, len(options.levelFormats))
	for i, levelFormat := range options.levelFormats {
		levelFormats[i] = headerFormat{format: levelFormat}
		for _, format := range levelFormat {
			if format == Source || format == LongSource {
				levelFormats[i].hasSource = true
				break
			}
		}
	}
	return &standardLogger{
		options:         options,
		level:           options.startingLevel,
		name:            name,
		writer:          os.Stdout,
		levelFormats:    levelFormats,
		criticalFailure: options.failingCriticals,
	}
}

// SetLevel sets the level of the logger
func (logger *standardLogger) SetLevel(level severity) {
	logger.level = level
}

func (logger *standardLogger) Level() severity {
	return logger.level
}

func (logger *standardLogger) Critical(v ...interface{}) {
	logger.output(0, 2, fmt.Sprintln(v...))
	if logger.criticalFailure {
		os.Exit(1)
	}
}

func (logger *standardLogger) Criticalf(format string, v ...interface{}) {
	logger.output(0, 2, fmt.Sprintf(format, v...))
	if logger.criticalFailure {
		os.Exit(1)
	}
}

func (logger *standardLogger) Error(v ...interface{}) {
	logger.println(ERROR, v...)
}

func (logger *standardLogger) Errorf(format string, v ...interface{}) {
	logger.printf(ERROR, format, v...)
}

func (logger *standardLogger) Warning(v ...interface{}) {
	logger.println(WARNING, v...)
}

func (logger *standardLogger) Warningf(format string, v ...interface{}) {
	logger.printf(WARNING, format, v...)
}

func (logger *standardLogger) Info(v ...interface{}) {
	logger.println(INFO, v...)
}

func (logger *standardLogger) Infof(format string, v ...interface{}) {
	logger.printf(INFO, format, v...)
}

func (logger *standardLogger) Debug(v ...interface{}) {
	logger.println(DEBUG, v...)
}

func (logger *standardLogger) Debugf(format string, v ...interface{}) {
	logger.printf(DEBUG, format, v...)
}

func (logger *standardLogger) Trace(v ...interface{}) {
	logger.println(TRACE, v...)
}

func (logger *standardLogger) Tracef(format string, v ...interface{}) {
	logger.printf(TRACE, format, v...)
}

func (logger *standardLogger) println(level severity, v ...interface{}) {
	if level <= logger.level {
		logger.output(level, 3, fmt.Sprintln(v...))
	}
}

func (logger *standardLogger) printf(level severity, format string, v ...interface{}) {
	if level <= logger.level {
		logger.output(level, 3, fmt.Sprintf(format, v...))
	}
}

func (logger *standardLogger) output(level severity, callDepth int, s string) error {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	if level > logger.level {
		return nil
	}

	now := time.Now()

	var file string
	var line int
	levelHeaderFormat := logger.levelFormats[level]
	if levelHeaderFormat.hasSource {
		// Release lock while getting caller info - it's expensive.
		logger.mutex.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		}
		logger.mutex.Lock()
	}
	logger.buffer = logger.buffer[:0]
	logger.formatHeader(&logger.buffer, level, now, file, line, levelHeaderFormat.format)
	logger.buffer = append(logger.buffer, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		logger.buffer = append(logger.buffer, '\n')
	}
	_, err := logger.writer.Write(logger.buffer)
	return err
}

var digits = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func itoa(buf *[]byte, i int, padding int) {
	var b [20]byte
	n := 19
	for i > 10 || padding > 1 {
		padding--
		m := i % 10
		b[n] = digits[m]
		n--
		i = i / 10
	}
	b[n] = digits[i]
	*buf = append(*buf, b[n:]...)
}

func timetoa(buf *[]byte, flags int, t time.Time) {
	if flags&LUTC != 0 {
		t = t.UTC()
	}
	if flags&Ldate != 0 {
		year, month, day := t.Date()
		itoa(buf, year, 4)
		*buf = append(*buf, '/')
		itoa(buf, int(month), 2)
		*buf = append(*buf, '/')
		itoa(buf, day, 2)
		*buf = append(*buf, ' ')
	}
	if flags&(Ltime|Lmicroseconds) != 0 {
		hour, min, sec := t.Clock()
		itoa(buf, hour, 2)
		*buf = append(*buf, ':')
		itoa(buf, min, 2)
		*buf = append(*buf, ':')
		itoa(buf, sec, 2)
		if flags&Lmicroseconds != 0 {
			*buf = append(*buf, '.')
			itoa(buf, t.Nanosecond()/1e3, 6)
		}
		*buf = append(*buf, ' ')
	}
}

func sourcetobuf(buf *[]byte, source string, line int) {
	*buf = append(*buf, source...)
	*buf = append(*buf, ':')
	itoa(buf, line, 0)
}

func (logger *standardLogger) formatHeader(buf *[]byte, level severity, t time.Time, file string, line int, format []uint) {
	for _, f := range format {
		switch f {
		case Separator:
			*buf = append(*buf, '-')
		case Name:
			*buf = append(*buf, logger.name...)
		case LogLevel:
			*buf = append(*buf, levelTokens[level]...)
		case LongSource:
			sourcetobuf(buf, file, line)
		case Source:
			i := strings.LastIndexByte(file, '/')
			sourcetobuf(buf, file[i+1:], line)
		case Time:
			timetoa(buf, logger.options.dateFlags, t)
		}
		*buf = append(*buf, ' ')
	}
}
