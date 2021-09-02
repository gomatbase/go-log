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

func ExamplePrintln() {
	resetLoggers()
	SetDefaultLogger(Standard().WithLevels(TRACE + 1).WithStartingLevel(TRACE + 1))

	Println(CRITICAL, "CRT")
	Println(ERROR, "ERR")
	Println(WARNING, "WRN")
	Println(INFO, "INF")
	Println(DEBUG, "DBG")
	Println(TRACE, "TRC")
	Println(TRACE+1, "CUSTOM")
	Println(TRACE+2, "CUSTOM2")

	// os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
	// TRC
	// CUSTOM
}

func ExamplePrintf() {
	resetLoggers()
	SetDefaultLogger(Standard().WithLevels(TRACE + 1).WithStartingLevel(TRACE + 1))

	Printf(CRITICAL, "%v", "CRT")
	Printf(ERROR, "%v", "ERR")
	Printf(WARNING, "%v", "WRN")
	Printf(INFO, "%v", "INF")
	Printf(DEBUG, "%v", "DBG")
	Printf(TRACE, "%v", "TRC")
	Printf(TRACE+1, "%v", "CUSTOM")
	Printf(TRACE+2, "%v", "CUSTOM2")

	// os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
	// TRC
	// CUSTOM
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

func ExampleCustomLoggerPrintln() {
	logger, _ := GetWithOptions("PLN", Standard().WithLevels(TRACE+1).WithLogPrefix(Name, Separator))
	logger.SetLevel(TRACE + 1)

	logger.Println(CRITICAL, "CRT")
	logger.Println(ERROR, "ERR")
	logger.Println(WARNING, "WRN")
	logger.Println(INFO, "INF")
	logger.Println(DEBUG, "DBG")
	logger.Println(TRACE, "TRC")
	logger.Println(TRACE+1, "CUSTOM")
	logger.Println(TRACE+2, "CUSTOM2")

	// Output:
	// PLN - CRT
	// PLN - ERR
	// PLN - WRN
	// PLN - INF
	// PLN - DBG
	// PLN - TRC
	// PLN - CUSTOM
}

func ExampleCustomLoggerPrintf() {
	logger, _ := GetWithOptions("PRTF", Standard().WithLevels(TRACE+1).WithLogPrefix(Name, Separator))
	logger.SetLevel(TRACE + 1)

	logger.Printf(CRITICAL, "%v", "CRT")
	logger.Printf(ERROR, "%v", "ERR")
	logger.Printf(WARNING, "%v", "WRN")
	logger.Printf(INFO, "%v", "INF")
	logger.Printf(DEBUG, "%v", "DBG")
	logger.Printf(TRACE, "%v", "TRC")
	logger.Printf(TRACE+1, "%v", "CUSTOM")
	logger.Printf(TRACE+2, "%v", "CUSTOM2")

	// Output:
	// PRTF - CRT
	// PRTF - ERR
	// PRTF - WRN
	// PRTF - INF
	// PRTF - DBG
	// PRTF - TRC
	// PRTF - CUSTOM
}

func ExampleCustomLoggerWithPattern() {
	logger, _ := GetWithOptions("PATTERN", Standard().WithLogPrefix(Name, Source, Separator))
	logger.SetLevel(ERROR)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Printf(ERROR, "%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	// Output:
	// PATTERN example_test.go:529 - CRT
	// PATTERN example_test.go:530 - ERR
	// PATTERN example_test.go:531 - ERR
}
