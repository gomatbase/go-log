// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"bytes"
	"testing"
)

var buf = &bytes.Buffer{}

func resetLoggers() {
	buf.Reset()
	loggers = make(map[string]Logger)
	defaultLogger, _ = newLogger(DEFAULT, Standard().WithWriter(buf).(*options))
	loggers[DEFAULT] = defaultLogger
}

func TestGettingLog(t *testing.T) {
	log, e := Get("Test Override")
	if e != nil {
		t.Error("Getting log should not raise an error.")
	}

	t.Run("Test getting empty log name", func(t *testing.T) {
		_, e := Get("")
		if e == nil {
			t.Error("Empty logger names should result in error.")
		}
	})

	t.Run("Test getting existing log with same options", func(t *testing.T) {
		overrideLog, e := GetWithOptions("Test Override", Standard())
		if e != nil {
			t.Error("Getting log with same options should not fail.")
		}
		if overrideLog != log {
			t.Error("Getting log with options not returning the same log with options.")
		}
	})

	t.Run("Test getting existing log with new options", func(t *testing.T) {
		overrideLog, e := GetWithOptions("Test Override", Standard().WithFailingCriticals())
		if e != nil {
			t.Error("Getting log with new options should not fail.")
		}
		if overrideLog != log {
			t.Error("Getting log with new options failing but returning a new log.")
		}
	})

	t.Run("Test getting logger with null options", func(t *testing.T) {
		overrideLog, e := GetWithOptions("Logger With Empty Options", nil)
		if e != nil {
			t.Error("Nil options should not result in error.")
		}
		if overrideLog == nil {
			t.Error("Nil options for non-exiting log should still result in a new logger.")
		} else if overrideLog.Level() != WARNING {
			t.Error("Default logger level WARNING is not set for nil options.")
		}
	})
}

func TestLevelNames(t *testing.T) {
	for i, name := range levelNames {
		if returnedName := LevelName(severity(i)); returnedName != name {
			t.Errorf("Unexpected level name for level %d: (expected %v, got %v)", i, name, returnedName)
		}
	}

	if returnedName := LevelName(1000); returnedName != "UNKNOWN" {
		t.Errorf("Unexpected levels should be unknown got %v with level 1000", returnedName)
	}
}

func TestSettingLogLevels(t *testing.T) {
	resetLoggers()

	t.Run("Test setting default logger log level", func(t *testing.T) {
		SetLevel(ERROR)
		if Level() != ERROR {
			t.Error("Failed to set default logger level to ERROR")
		}

		if e := SetLoggerLevel(DEFAULT, CRITICAL); e != nil {
			t.Error("Failed to set default logger level to CRITICAL")
		}
		if Level() != CRITICAL {
			t.Error("Default logger level did not change")
		}
	})

	t.Run("Test setting logger log level", func(t *testing.T) {
		loggerName := "LOG-LEVEL-LOG"
		logger, e := Get(loggerName)
		if e != nil {
			t.Error("Failed to get LOG-LEVEL-LOG")
		}
		if logger.Level() != WARNING {
			t.Error("Logger with default init not in warning level")
		}

		if e := SetLoggerLevel(loggerName, CRITICAL); e != nil {
			t.Error("Failed to set", loggerName, "logger level to CRITICAL")
		}
		if logger.Level() != CRITICAL {
			t.Error("Failed to set", loggerName, "logger level to CRITICAL")
		}
	})

	t.Run("Test setting non-existing logger log level", func(t *testing.T) {
		if e := SetLoggerLevel("NON-EXISTING", CRITICAL); e == nil {
			t.Error("Setting log level of a non-existing logger should end in error")
		}
	})

	t.Run("Test setting logger log level", func(t *testing.T) {
		loggerName := "LOG-LEVEL-LOG"
		logLevels := make(map[string]severity)
		logLevels[DEFAULT] = DEBUG
		logLevels[loggerName] = CRITICAL
		logLevels["NON-EXISTING"] = ERROR

		logLevels = SetLoggerLevels(logLevels)

		if _, found := logLevels["NON-EXISTING"]; len(logLevels) != 2 || found {
			t.Error("Setting log levels including non-existing logger reports as having log level set.")
		}

		if logLevels[DEFAULT] != DEBUG || logLevels[DEFAULT] != Level() {
			t.Error("Failed to set log level of default logger.")
		}

		if logLevels[loggerName] != CRITICAL || logLevels[DEFAULT] != Level() {
			t.Error("Failed to set log level of default logger.")
		}
	})

}

func TestGettingLogLevels(t *testing.T) {
	resetLoggers()
	_, _ = GetWithOptions("LOG1", Standard().WithStartingLevel(DEBUG).WithoutFailingCriticals())
	_, _ = GetWithOptions("LOG2", Standard().WithStartingLevel(TRACE+1))

	t.Run("Test getting logger log levels", func(t *testing.T) {
		loggerLevels := LoggerLevels()
		if len(loggerLevels) != 3 {
			t.Error("Unexpected number of log level entries :", loggerLevels)
		}

		if level, found := loggerLevels[DEFAULT]; !found {
			t.Error("DEFAULT logger level not reported.")
		} else if level != WARNING {
			t.Error("DEFAULT logger level incorrect.")
		}

		if level, found := loggerLevels["LOG1"]; !found {
			t.Error("LOG1 logger level not reported.")
		} else if level != DEBUG {
			t.Error("LOG1 logger level incorrect.")
		}

		if level, found := loggerLevels["LOG2"]; !found {
			t.Error("LOG2 logger level not reported.")
		} else if level != TRACE+1 {
			t.Error("LOG2 logger level incorrect.")
		}
	})

	t.Run("Test getting logger log level", func(t *testing.T) {
		if level, e := LoggerLevel(DEFAULT); e != nil {
			t.Error("DEFAULT logger level errored :", e)
		} else if level != WARNING {
			t.Error("DEFAULT logger level incorrect.")
		}

		if level, e := LoggerLevel("LOG1"); e != nil {
			t.Error("LOG1 logger level errored :", e)
		} else if level != DEBUG {
			t.Error("LOG1 logger level incorrect.")
		}

		if level, e := LoggerLevel("LOG2"); e != nil {
			t.Error("LOG2 logger level errored :", e)
		} else if level != TRACE+1 {
			t.Error("LOG2 logger level incorrect.")
		}

		if level, e := LoggerLevel("UNKNOWN"); e == nil {
			t.Error("Getting level for unknown logger should error :", level)
		}
	})

	t.Run("Test getting logger log level names", func(t *testing.T) {
		loggerLevels := LoggerLevelNames()
		if len(loggerLevels) != 3 {
			t.Error("Unexpected number of log level entries")
		}

		if level, found := loggerLevels[DEFAULT]; !found {
			t.Error("DEFAULT logger level not reported.")
		} else if level != "WARNING" {
			t.Error("DEFAULT logger level incorrect.")
		}

		if level, found := loggerLevels["LOG1"]; !found {
			t.Error("LOG1 logger level not reported.")
		} else if level != "DEBUG" {
			t.Error("LOG1 logger level incorrect.")
		}

		if level, found := loggerLevels["LOG2"]; !found {
			t.Error("LOG2 logger level not reported.")
		} else if level != "UNKNOWN" {
			t.Error("LOG2 logger level incorrect.")
		}
	})

	t.Run("Test getting logger log level name", func(t *testing.T) {
		if level, e := LoggerLevelName(DEFAULT); e != nil {
			t.Error("DEFAULT logger level errored :", e)
		} else if level != "WARNING" {
			t.Error("DEFAULT logger level incorrect :", level)
		}

		if level, e := LoggerLevelName("LOG1"); e != nil {
			t.Error("LOG1 logger level errored :", e)
		} else if level != "DEBUG" {
			t.Error("LOG1 logger level incorrect :", level)
		}

		if level, e := LoggerLevelName("LOG2"); e != nil {
			t.Error("LOG2 logger level errored :", e)
		} else if level != "UNKNOWN" {
			t.Error("LOG2 logger level incorrect :", level)
		}

		if level, e := LoggerLevelName("UNKNOWN"); e == nil {
			t.Error("Getting level for unknown logger should error :", level)
		}
	})

}
