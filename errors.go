// Copyright 2020 GOM. All rights reserved.
// Since 25/06/2021 By GOM
// Licensed under MIT License

package log

import (
	"github.com/gomatbase/go-error"
)

const (

	// ErrEmptyLoggerName Error raised when trying to refer to a logger with an empty name
	ErrEmptyLoggerName = err.Error("Logger name may not be empty")

	// ErrReinitializingExistingLogger Error raised when trying to initialize an existing logger with different options
	ErrReinitializingExistingLogger = err.Error("trying to initialize an already initialized logger with different options")

	// ErrLoggerDoesNotExist Error raised when referring to a non-existing logger
	ErrLoggerDoesNotExist = err.Error("logger with given name doesn't exist")
)
