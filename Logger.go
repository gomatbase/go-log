package log

// Logger defines the interface a Logger implementation must provide
type Logger interface {

	// SetLevel sets the current log level of the logger
	SetLevel(severity)

	// Level returns the current log level of the logger
	Level() severity

	// Critical logs the message(s) at the critical level
	Critical(v ...interface{})

	// Criticalf logs the formatted message at the critical level
	Criticalf(format string, v ...interface{})

	// Error logs the message(s) at the error level
	Error(v ...interface{})

	// Errorf logs the formatted message at the error level
	Errorf(format string, v ...interface{})

	// Warning logs the message(s) at the warning level
	Warning(v ...interface{})

	// Warningf logs the formatted message at the warning level
	Warningf(format string, v ...interface{})

	// Info logs the message(s) at the info level
	Info(v ...interface{})

	// Infof logs the formatted message at the info level
	Infof(format string, v ...interface{})

	// Debug logs the message(s) at the debug level
	Debug(v ...interface{})

	// Debugf logs the formatted message at the debug level
	Debugf(format string, v ...interface{})

	// Trace logs the message(s) at the trace level
	Trace(v ...interface{})

	// Tracef logs the formatted message at the trace level
	Tracef(format string, v ...interface{})

	// println logs the message(s) at the provided level
	println(level severity, v ...interface{})

	// printf logs the formatted message at the provided level
	printf(level severity, format string, v ...interface{})
}
