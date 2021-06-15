package log

import (
	"errors"
	l "log"
	"sync"
)

const (
	CRITICAL = iota
	ERROR
	WARNING
	INFO
	DEBUG
	TRACE
)

var levelNames = []string{"ERROR", "WARNING", "INFO", "DEBUG", "TRACE"}

var (
	loggers          = make(map[string]*logger)
	lock             = sync.Mutex{}
	defaultLogger, _ = getWithOptions("", WithoutOptions())
)

type Logger interface {
	SetLevel(uint)
	Level() uint
	Critical(v ...interface{})
	Criticalf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Warning(v ...interface{})
	Warningf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Trace(v ...interface{})
	Tracef(format string, v ...interface{})
	Println(level uint, v ...interface{})
	Printf(level uint, format string, v ...interface{})
}

type logger struct {
	options   *Options
	level     uint
	name      string
	log       *l.Logger
	critical  func(...interface{})
	criticalf func(string, ...interface{})
}

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

func Get(name string) (Logger, error) {
	return GetWithOptions(name, WithoutOptions())
}

func GetWithOptions(name string, options *Options) (Logger, error) {
	if len(name) == 0 {
		return nil, errors.New("name of logger cannot be empty")
	}
	return getWithOptions(name, options)
}

func getWithOptions(name string, options *Options) (Logger, error) {
	lock.Lock()
	defer lock.Unlock()

	logger, found := loggers[name]
	var e error
	if !found {
		loggers[name] = newLogger(name, options)
		logger = loggers[name]
	} else if !logger.options.equals(options) {
		e = errors.New("trying to initialize an already initialized logger with different options")
	}

	return logger, e
}

func OverrideLogWithOptions(name string, options *Options) (Logger, error) {
	if len(name) == 0 {
		return nil, errors.New("name of logger cannot be empty")
	}

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

func SetLevel(level uint) {
	defaultLogger.SetLevel(level)
}

func Level() uint {
	return defaultLogger.Level()
}

func LevelName(level uint) string {
	if int(level) >= len(levelNames) {
		return "UNKNOWN"
	}
	return levelNames[level]
}

func Critical(v ...interface{}) {
	defaultLogger.Critical(v...)
}

func Criticalf(format string, v ...interface{}) {
	defaultLogger.Criticalf(format, v...)
}

func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Warning(v ...interface{}) {
	defaultLogger.Println(WARNING, v...)
}

func Warningf(format string, v ...interface{}) {
	defaultLogger.Printf(WARNING, format, v...)
}

func Info(v ...interface{}) {
	defaultLogger.Println(INFO, v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Printf(INFO, format, v...)
}

func Debug(v ...interface{}) {
	defaultLogger.Println(DEBUG, v...)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Printf(DEBUG, format, v...)
}

func Trace(v ...interface{}) {
	defaultLogger.Println(TRACE, v...)
}

func Tracef(format string, v ...interface{}) {
	defaultLogger.Printf(TRACE, format, v...)
}

func Println(level uint, v ...interface{}) {
	defaultLogger.Println(level, v...)
}

func Printf(level uint, format string, v ...interface{}) {
	defaultLogger.Printf(level, format, v...)
}
