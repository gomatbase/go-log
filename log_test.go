// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"reflect"
	"testing"
)

func TestGettingLog(t *testing.T) {
	log, e := Get("Test Override")
	if e != nil {
		t.Error("Getting log should not raise an error.")
	}
	t.Run("Test getting existing log with same options", func(t *testing.T) {
		overrideLog, e := GetWithOptions("Test Override", WithoutOptions())
		if e != nil {
			t.Error("Getting log with same options should not fail.")
		}
		if overrideLog != log {
			t.Error("Getting log with options not returning the same log with options.")
		}
	})

	t.Run("Test getting existing log with new options", func(t *testing.T) {
		overrideLog, e := GetWithOptions("Test Override", WithOptions().WithFailingCriticals())
		if e == nil {
			t.Error("Getting log with new options should fail.")
		}
		if overrideLog != log {
			t.Error("Getting log with new options failing but returning a new log.")
		}
	})

	t.Run("Test overriding non-existing log", func(t *testing.T) {
		overrideLog, e := OverrideLogWithOptions("Another Test Override", WithoutOptions())
		if e == nil {
			t.Error("Overriding nonexistent log should fail.")
		}
		if overrideLog != nil {
			t.Error("Overriding nonexistent log should not return one.")
		}
	})

	t.Run("Test overriding existing log with same options", func(t *testing.T) {
		options := WithoutOptions()
		logWriter := log.(*logger).log
		overrideLog, e := OverrideLogWithOptions("Test Override", options)
		if e != nil {
			t.Error("Overriding existing log with same options should not fail.")
		}
		if overrideLog != log {
			t.Error("Overriding log options with same options should return the same log.")
		}
		if options == log.(*logger).options {
			t.Error("Overriding log options with same options should not change the options.")
		}
		if logWriter != log.(*logger).log {
			t.Error("Overriding log options with same options should not affect the log writer.")
		}
		if reflect.ValueOf(log.(*logger).critical).Pointer() != reflect.ValueOf(logWriter.Println).Pointer() ||
			reflect.ValueOf(log.(*logger).criticalf).Pointer() != reflect.ValueOf(logWriter.Printf).Pointer() {
			t.Error("Overriding log options with same options should not affect the critical functions")
		}
	})

	t.Run("Test overriding existing log with new options", func(t *testing.T) {
		options := WithOptions().WithStartingLevel(TRACE)
		logWriter := log.(*logger).log
		overrideLog, e := OverrideLogWithOptions("Test Override", options)
		if e != nil {
			t.Error("Overriding existing log with new options should not fail.")
		}
		if overrideLog != log {
			t.Error("Overriding log options with new options should return the same (updated) log.")
		}
		if options != log.(*logger).options {
			t.Error("Overriding log options with new options should update the options.")
		}
		if logWriter == log.(*logger).log {
			t.Error("Overriding log options with new options should update the log writer.")
		}
		if reflect.ValueOf(log.(*logger).critical).Pointer() != reflect.ValueOf(log.(*logger).log.Println).Pointer() ||
			reflect.ValueOf(log.(*logger).criticalf).Pointer() != reflect.ValueOf(log.(*logger).log.Printf).Pointer() {
			t.Error("Overriding log options with new options should affect the critical functions to use the new log writter.")
		}
		if log.Level() != TRACE {
			t.Error("Overriding log level should update current log level")
		}
		log, e = Get("Test Override")
		if e != nil {
			t.Error("Getting log with non-default options should not raise an error:", e)
		}
	})
}

func TestLevelNames(t *testing.T) {
	for i, name := range levelNames {
		if returnedName := LevelName(uint(i)); returnedName != name {
			t.Errorf("Unexpected level name for level %d: (expected %v, got %v)", i, name, returnedName)
		}
	}

	if returnedName := LevelName(1000); returnedName != "UNKNOWN" {
		t.Errorf("Unexpected levels should be unknown got %v with level 1000", returnedName)
	}
}
