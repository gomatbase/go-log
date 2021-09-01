// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"bytes"
	"os"
)

var buf = &bytes.Buffer{}

func resetLoggers() {
	buf.Reset()
	loggers = make(map[string]Logger)
	defaultLogger, _ = newLogger(DEFAULT, Standard())
	defaultLogger.(*standardLogger).writer = buf
	loggers[DEFAULT] = defaultLogger
}

func ExampleTrace() {
	resetLoggers()
	SetLevel(TRACE)

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
	// TRC
}

func ExampleDebug() {
	resetLoggers()
	SetLevel(DEBUG)

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
}

func ExampleInfo() {
	resetLoggers()
	SetLevel(INFO)

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
}

func ExampleWarning() {
	resetLoggers()
	SetLevel(WARNING)

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
}

func ExampleError() {
	resetLoggers()
	SetLevel(ERROR)

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
}

func ExampleCritical() {
	resetLoggers()
	SetLevel(CRITICAL)

	Critical("CRT")
	Error("ERR")
	Warning("WRN")
	Info("INF")
	Debug("DBG")
	Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
}

func ExampleTracef() {
	resetLoggers()
	SetLevel(TRACE)

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
	// TRC
}

func ExampleDebugf() {
	resetLoggers()
	SetLevel(DEBUG)

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
}

func ExampleInfof() {
	resetLoggers()
	SetLevel(INFO)

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
}

func ExampleWarningf() {
	resetLoggers()
	SetLevel(WARNING)

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
	// WRN
}

func ExampleErrorf() {
	resetLoggers()
	SetLevel(ERROR)

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
	// ERR
}

func ExampleCriticalf() {
	resetLoggers()
	SetLevel(CRITICAL)

	Criticalf("%v", "CRT")
	Errorf("%v", "ERR")
	Warningf("%v", "WRN")
	Infof("%v", "INF")
	Debugf("%v", "DBG")
	Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT
}

func ExamplePrintln() {
	resetLoggers()
	defaultLogger, _ = newLogger(DEFAULT, Standard().WithLevels(TRACE+1).WithStartingLevel(TRACE+1))
	defaultLogger.(*standardLogger).writer = buf

	Println(CRITICAL, "CRT")
	Println(ERROR, "ERR")
	Println(WARNING, "WRN")
	Println(INFO, "INF")
	Println(DEBUG, "DBG")
	Println(TRACE, "TRC")
	Println(TRACE+1, "CUSTOM")
	Println(TRACE+2, "CUSTOM2")

	os.Stdout.WriteString(buf.String())

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
	defaultLogger, _ = newLogger(DEFAULT, Standard().WithLevels(TRACE+1).WithStartingLevel(TRACE+1))
	defaultLogger.(*standardLogger).writer = buf

	Printf(CRITICAL, "%v", "CRT")
	Printf(ERROR, "%v", "ERR")
	Printf(WARNING, "%v", "WRN")
	Printf(INFO, "%v", "INF")
	Printf(DEBUG, "%v", "DBG")
	Printf(TRACE, "%v", "TRC")
	Printf(TRACE+1, "%v", "CUSTOM")
	Printf(TRACE+2, "%v", "CUSTOM2")

	os.Stdout.WriteString(buf.String())

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
	buf.Reset()
	logger, _ := GetWithOptions("TRC", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(TRACE)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// TRC - CRT
	// TRC - ERR
	// TRC - WRN
	// TRC - INF
	// TRC - DBG
	// TRC - TRC
}

func ExampleCustomLoggerDebug() {
	buf.Reset()
	logger, _ := GetWithOptions("DBG", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(DEBUG)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// DBG - CRT
	// DBG - ERR
	// DBG - WRN
	// DBG - INF
	// DBG - DBG
}

func ExampleCustomLoggerInfo() {
	buf.Reset()
	logger, _ := GetWithOptions("INF", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(INFO)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// INF - CRT
	// INF - ERR
	// INF - WRN
	// INF - INF
}

func ExampleCustomLoggerWarning() {
	buf.Reset()
	logger, _ := GetWithOptions("WRN", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(WARNING)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// WRN - CRT
	// WRN - ERR
	// WRN - WRN
}

func ExampleCustomLoggerError() {
	buf.Reset()
	logger, _ := GetWithOptions("ERR", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(ERROR)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// ERR - CRT
	// ERR - ERR
}

func ExampleCustomLoggerCritical() {
	buf.Reset()
	logger, _ := GetWithOptions("CRT", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(CRITICAL)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRT - CRT
}

func ExampleCustomLoggerTracef() {
	buf.Reset()
	logger, _ := GetWithOptions("TRCF", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(TRACE)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// TRCF - CRT
	// TRCF - ERR
	// TRCF - WRN
	// TRCF - INF
	// TRCF - DBG
	// TRCF - TRC
}

func ExampleCustomLoggerDebugf() {
	buf.Reset()
	logger, _ := GetWithOptions("DBGF", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(DEBUG)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// DBGF - CRT
	// DBGF - ERR
	// DBGF - WRN
	// DBGF - INF
	// DBGF - DBG
}

func ExampleCustomLoggerInfof() {
	buf.Reset()
	logger, _ := GetWithOptions("INFF", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(INFO)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// INFF - CRT
	// INFF - ERR
	// INFF - WRN
	// INFF - INF
}

func ExampleCustomLoggerWarningf() {
	buf.Reset()
	logger, _ := GetWithOptions("WRNF", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(WARNING)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// WRNF - CRT
	// WRNF - ERR
	// WRNF - WRN
}

func ExampleCustomLoggerErrorf() {
	buf.Reset()
	logger, _ := GetWithOptions("ERRF", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(ERROR)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// ERRF - CRT
	// ERRF - ERR
}

func ExampleCustomLoggerCriticalf() {
	buf.Reset()
	logger, _ := GetWithOptions("CRTF", Standard().WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(CRITICAL)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// CRTF - CRT
}

func ExampleCustomLoggerPrintln() {
	buf.Reset()
	logger, _ := GetWithOptions("PLN", Standard().WithLevels(TRACE+1).WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(TRACE + 1)

	logger.Println(CRITICAL, "CRT")
	logger.Println(ERROR, "ERR")
	logger.Println(WARNING, "WRN")
	logger.Println(INFO, "INF")
	logger.Println(DEBUG, "DBG")
	logger.Println(TRACE, "TRC")
	logger.Println(TRACE+1, "CUSTOM")
	logger.Println(TRACE+2, "CUSTOM2")

	os.Stdout.WriteString(buf.String())

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
	buf.Reset()
	logger, _ := GetWithOptions("PRTF", Standard().WithLevels(TRACE+1).WithLogPrefix(Name, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(TRACE + 1)

	logger.Printf(CRITICAL, "%v", "CRT")
	logger.Printf(ERROR, "%v", "ERR")
	logger.Printf(WARNING, "%v", "WRN")
	logger.Printf(INFO, "%v", "INF")
	logger.Printf(DEBUG, "%v", "DBG")
	logger.Printf(TRACE, "%v", "TRC")
	logger.Printf(TRACE+1, "%v", "CUSTOM")
	logger.Printf(TRACE+2, "%v", "CUSTOM2")

	os.Stdout.WriteString(buf.String())

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
	buf.Reset()
	logger, _ := GetWithOptions("PATTERN", Standard().WithLogPrefix(Name, Source, Separator))
	logger.(*standardLogger).writer = buf
	logger.SetLevel(ERROR)

	logger.Criticalf("%v", "CRT")
	logger.Errorf("%v", "ERR")
	logger.Printf(ERROR, "%v", "ERR")
	logger.Warningf("%v", "WRN")
	logger.Infof("%v", "INF")
	logger.Debugf("%v", "DBG")
	logger.Tracef("%v", "TRC")

	os.Stdout.WriteString(buf.String())

	// Output:
	// PATTERN example_test.go:626 - CRT
	// PATTERN example_test.go:627 - ERR
	// PATTERN example_test.go:628 - ERR
}
