// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log_test

import (
	"github.com/gomatbase/go-log"
)

func ExampleTrace() {
	_ = log.SetDefaultLogger(log.Standard())
	log.SetLevel(log.TRACE)

	log.Critical("CRT")
	log.Error("ERR")
	log.Warning("WRN")
	log.Info("INF")
	log.Debug("DBG")
	log.Trace("TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
	// TRC
}

func ExampleDebug() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.DEBUG))

	log.Critical("CRT")
	log.Error("ERR")
	log.Warning("WRN")
	log.Info("INF")
	log.Debug("DBG")
	log.Trace("TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
}

func ExampleInfo() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.INFO))

	log.Critical("CRT")
	log.Error("ERR")
	log.Warning("WRN")
	log.Info("INF")
	log.Debug("DBG")
	log.Trace("TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
}

func ExampleWarning() {
	log.SetDefaultLogger(log.Standard())

	log.Critical("CRT")
	log.Error("ERR")
	log.Warning("WRN")
	log.Info("INF")
	log.Debug("DBG")
	log.Trace("TRC")

	// Output:
	// CRT
	// ERR
	// WRN
}

func ExampleError() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.ERROR))

	log.Critical("CRT")
	log.Error("ERR")
	log.Warning("WRN")
	log.Info("INF")
	log.Debug("DBG")
	log.Trace("TRC")

	// Output:
	// CRT
	// ERR
}

func ExampleCritical() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.CRITICAL))

	log.Critical("CRT")
	log.Error("ERR")
	log.Warning("WRN")
	log.Info("INF")
	log.Debug("DBG")
	log.Trace("TRC")

	// Output:
	// CRT
}

func ExampleTracef() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.TRACE))

	log.Criticalf("%v", "CRT")
	log.Errorf("%v", "ERR")
	log.Warningf("%v", "WRN")
	log.Infof("%v", "INF")
	log.Debugf("%v", "DBG")
	log.Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
	// TRC
}

func ExampleDebugf() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.DEBUG))

	log.Criticalf("%v", "CRT")
	log.Errorf("%v", "ERR")
	log.Warningf("%v", "WRN")
	log.Infof("%v", "INF")
	log.Debugf("%v", "DBG")
	log.Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
	// DBG
}

func ExampleInfof() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.INFO))

	log.Criticalf("%v", "CRT")
	log.Errorf("%v", "ERR")
	log.Warningf("%v", "WRN")
	log.Infof("%v", "INF")
	log.Debugf("%v", "DBG")
	log.Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
	// WRN
	// INF
}

func ExampleWarningf() {
	log.SetDefaultLogger(log.Standard())

	log.Criticalf("%v", "CRT")
	log.Errorf("%v", "ERR")
	log.Warningf("%v", "WRN")
	log.Infof("%v", "INF")
	log.Debugf("%v", "DBG")
	log.Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
	// WRN
}

func ExampleErrorf() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.ERROR))

	log.Criticalf("%v", "CRT")
	log.Errorf("%v", "ERR")
	log.Warningf("%v", "WRN")
	log.Infof("%v", "INF")
	log.Debugf("%v", "DBG")
	log.Tracef("%v", "TRC")

	// Output:
	// CRT
	// ERR
}

func ExampleCriticalf() {
	_ = log.SetDefaultLogger(log.Standard().WithStartingLevel(log.CRITICAL))

	log.Criticalf("%v", "CRT")
	log.Errorf("%v", "ERR")
	log.Warningf("%v", "WRN")
	log.Infof("%v", "INF")
	log.Debugf("%v", "DBG")
	log.Tracef("%v", "TRC")

	// Output:
	// CRT
}

func ExampleSetDefaultLogger() {
	_ = log.SetDefaultLogger(log.Standard().WithLogPrefix(log.Name, log.Source, log.Separator).WithStartingLevel(log.ERROR))

	log.Criticalf("%v", "CRT")
	log.Errorf("%v", "ERR")
	log.Warningf("%v", "WRN")
	log.Infof("%v", "INF")
	log.Debugf("%v", "DBG")
	log.Tracef("%v", "TRC")

	// Output:
	// DEFAULT example_test.go:213 - CRT
	// DEFAULT example_test.go:214 - ERR
}

func ExampleStandard() {
	logger, _ := log.GetWithOptions("TRC", log.Standard().WithLogPrefix(log.Name, log.Separator).WithStartingLevel(log.TRACE))

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	logger, _ = log.GetWithOptions("DBG", log.Standard().WithLogPrefix(log.Name, log.Separator))
	logger.SetLevel(log.DEBUG)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	logger, _ = log.GetWithOptions("INF", log.Standard().WithLogPrefix(log.Name, log.Separator))
	logger.SetLevel(log.INFO)

	logger.Critical("CRT")
	logger.Error("ERR")
	logger.Warning("WRN")
	logger.Info("INF")
	logger.Debug("DBG")
	logger.Trace("TRC")

	logger, _ = log.GetWithOptions("WRN", log.Standard().WithLogPrefix(log.Name, log.Separator))
	logger.SetLevel(log.WARNING)

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
	// DBG - CRT
	// DBG - ERR
	// DBG - WRN
	// DBG - INF
	// DBG - DBG
	// INF - CRT
	// INF - ERR
	// INF - WRN
	// INF - INF
	// WRN - CRT
	// WRN - ERR
	// WRN - WRN
}
