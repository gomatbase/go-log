// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"bytes"
	"os"
	"testing"
)

var buf = &bytes.Buffer{}

func TestMain(m *testing.M) {
	defaultLogger = newLogger(DEFAULT, WithOptions().WithWriter(buf))
	loggers[DEFAULT] = defaultLogger
	os.Exit(m.Run())
}

func ExampleTrace() {
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
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
	buf.Reset()
	SetLevel(TRACE + 1)

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
	buf.Reset()
	SetLevel(TRACE + 1)

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
