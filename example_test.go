// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"bytes"
)

var buf = &bytes.Buffer{}

func resetLoggers() {
	buf.Reset()
	loggers = make(map[string]Logger)
	defaultLogger, _ = newLogger(DEFAULT, Standard().WithWriter(buf).(*options))
	loggers[DEFAULT] = defaultLogger
}

func ExampleTrace() {
	SetDefaultLogger(Standard())
	SetLevel(TRACE)

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
	// TRC
}

func ExampleDebug() {
	SetDefaultLogger(Standard().WithStartingLevel(DEBUG))

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
}

func ExampleInfo() {
	SetDefaultLogger(Standard().WithStartingLevel(INFO))

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
}

func ExampleWarning() {
	SetDefaultLogger(Standard())

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	// Output:
	// CRT
	// ERR
	// WRN
}

func ExampleError() {
	SetDefaultLogger(Standard().WithStartingLevel(ERROR))

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	// Output:
	// CRT
	// ERR
}

func ExampleCritical() {
	SetDefaultLogger(Standard().WithStartingLevel(CRITICAL))

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	// Output:
	// CRT
}

func ExampleTracef() {
	SetDefaultLogger(Standard().WithStartingLevel(TRACE))

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
	// TRC
}

func ExampleDebugf() {
	SetDefaultLogger(Standard().WithStartingLevel(DEBUG))

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
}

func ExampleInfof() {
	SetDefaultLogger(Standard().WithStartingLevel(INFO))

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
}

func ExampleWarningf() {
	SetDefaultLogger(Standard())

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
	// WRN
}

func ExampleErrorf() {
	SetDefaultLogger(Standard().WithStartingLevel(ERROR))

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
}

func ExampleCriticalf() {
	SetDefaultLogger(Standard().WithStartingLevel(CRITICAL))

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	// Output:
	// CRT
}

func ExampleCustomLoggerTrace() {
	logger, _ := GetWithOptions("TRC", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(TRACE)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	// Output:
	// TRC - CRT
	// TRC - ERR
	// TRC - WRN
	// TRC - INF
	// TRC - DBG
	// TRC - TRC
}

func ExampleCustomLoggerDebug() {
	logger, _ := GetWithOptions("DBG", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(DEBUG)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	// Output:
	// DBG - CRT
	// DBG - ERR
	// DBG - WRN
	// DBG - INF
	// DBG - DBG
}

func ExampleCustomLoggerInfo() {
	logger, _ := GetWithOptions("INF", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(INFO)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	// Output:
	// INF - CRT
	// INF - ERR
	// INF - WRN
	// INF - INF
}

func ExampleCustomLoggerWarning() {
	logger, _ := GetWithOptions("WRN", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(WARNING)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	// Output:
	// WRN - CRT
	// WRN - ERR
	// WRN - WRN
}

func ExampleCustomLoggerError() {
	logger, _ := GetWithOptions("ERR", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(ERROR)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	// Output:
	// ERR - CRT
	// ERR - ERR
}

func ExampleCustomLoggerCritical() {
	logger, _ := GetWithOptions("CRT", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(CRITICAL)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	// Output:
	// CRT - CRT
}

func ExampleCustomLoggerTracef() {
	logger, _ := GetWithOptions("TRCF", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(TRACE)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	// Output:
	// TRCF - CRT
	// TRCF - ERR
	// TRCF - WRN
	// TRCF - INF
	// TRCF - DBG
	// TRCF - TRC
}

func ExampleCustomLoggerDebugf() {
	logger, _ := GetWithOptions("DBGF", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(DEBUG)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	// Output:
	// DBGF - CRT
	// DBGF - ERR
	// DBGF - WRN
	// DBGF - INF
	// DBGF - DBG
}

func ExampleCustomLoggerInfof() {
	logger, _ := GetWithOptions("INFF", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(INFO)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	// Output:
	// INFF - CRT
	// INFF - ERR
	// INFF - WRN
	// INFF - INF
}

func ExampleCustomLoggerWarningf() {
	logger, _ := GetWithOptions("WRNF", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(WARNING)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	// Output:
	// WRNF - CRT
	// WRNF - ERR
	// WRNF - WRN
}

func ExampleCustomLoggerErrorf() {
	logger, _ := GetWithOptions("ERRF", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(ERROR)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	// Output:
	// ERRF - CRT
	// ERRF - ERR
}

func ExampleCustomLoggerCriticalf() {
	logger, _ := GetWithOptions("CRTF", Standard().WithLogPrefix(Name, Separator))
	logger.SetLevel(CRITICAL)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	// Output:
	// CRTF - CRT
}

func ExampleCustomLoggerWithPattern() {
	logger, _ := GetWithOptions("PATTERN", Standard().WithLogPrefix(Name, Source, Separator))
	logger.SetLevel(ERROR)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	// Output:
	// PATTERN example_test.go:433 - CRT
	// PATTERN example_test.go:434 - ERR
}
