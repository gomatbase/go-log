// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"testing"
)

func TestLogError(t *testing.T) {
	if ErrEmptyLoggerName.Error() != "Logger name may not be empty" {
		t.Error("Log Error not returning the proper error message.")
	}
}
