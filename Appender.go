// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

// Appender An Appender takes log entries and outputs it to whatever medium it implements
type Appender interface {
	Print(logEntry *LogEntry)
}
