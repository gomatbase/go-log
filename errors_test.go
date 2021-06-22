package log

import (
	"testing"
)

func TestLogError(t *testing.T) {
	if ErrEmptyLoggerName.Error() != "Logger name may not be empty" {
		t.Error("Log Error not returning the proper error message.")
	}
}
