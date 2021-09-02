package log

import (
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     severity
	Message   string
	Source    *string
	Line      int
}
