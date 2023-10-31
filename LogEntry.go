package log

import (
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     int
	Message   string
	Source    *string
	Line      int
}
