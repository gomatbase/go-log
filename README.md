# go-log
Simple logger wrapper allowing for log levels and runtime level setup.

### Usage

The log usage is pretty much the same as the default log package from go.

The main difference is the addition of a level variable that allows a decision to be made if the log will be printed based on the current log level.

The package initializes a default log which can be called directly from package functions. The log is initialized in WARNING level.

```go
package main

import "github.com/gomatbase/go-log"

func main() {
    log.Println(log.ERROR,"This log entry will be printed")
    log.Println(log.TRACE,"This log entry will be ignored")

    // output:
    // This log entry will be printed
}
```

The `Println` function has the counterpart `Printf` function allowing for formatted output.

```go
package main

import "github.com/gomatbase/go-log"

func main() {
    log.Printf(log.ERROR,"This %v entry will be printed","log")
    log.Printf(log.TRACE,"This %v entry will be ignored","log")
    
    // output:
    // This log entry will be printed
}
```

### Custom Loggers

The package will create a default logger which can be used for out-of-the-box logging.

The default logger will be initialized in level WARNING, it will output entries to stdout without timestamps
and a critical log will not result in a panic.

These are options that can be customized for custom loggers.

To create a new custom logger the function `Get(name string)` will create a new logger with the default options (same as 
the default logger) or return an existing logger. 

Any custom logger will output its name as a prefix of the log entry (more precisely it will output the name followed 
by " - " and then the log entry). The name of the logger will typically identify the source it is logging from,
like the package name, a filename or some other identifier the developer may find useful.

```go
package main

import (
    "os"

    "github.com/gomatbase/go-log"
)

func main() {
    // First time getting a logger named "TESTER", logger will be created
    logger, e := log.Get("TESTER")
    if e != nil {
        log.Critical("Unable to create custom logger : ", e)
        os.Exit(1)
    }
    logger.Warning("Logging a warning with custom logger")
    logger.Info("Logging information with custom logger")
    logger.SetLevel(log.INFO)
    
    // Getting the logger will return the same logger as it already exists
    customLogger, e := log.Get("TESTER")
    if e != nil {
        log.Critical("Unable to get custom logger TESTER : ", e)
        os.Exit(1)
    }
    customLogger.Warning("Logging a warning with custom logger")
    customLogger.Info("Logging information with custom logger")

    // output:
    // TESTER - Logging a warning with custom logger
    // TESTER - Logging a warning with custom logger
    // TESTER - Logging information with custom logger
}
```

#### Configuring a custom logger

To set specific options  for a logger (like changing the output stream,
adding log flags or ending fatally when logging criticals), `Get(name string, options log.Options)` can be used
instead.

To create the logger options, you may either create a new log.Options object or use the syntax-candy functions
`WithOptions()` or `WithoutOptions()`. Both these functions will return a log.Options object initialized with
the default values (the difference is only to make it more readable the intention of the
options object).

`log.Options` has several methods which allow setting the options values effectively acting as a simplistic builder.

After creating a logger, using `Get(name string, options log.Options)` with different options than the ones used
to create the logger will result in error. Using the same options will be the same as using `Get(name string)` for an
existing logger.

If for some reason there is a need to change the initial configurations of an existing logger, 
`OverrideLogWithOptions(name string, options log.Options)` must be used instead. This function will result in error if
overriding options for a non-existing logger. The function is also not thread safe, so if a logger is already in use
in another running thread, the operation may result in unexpected results. The method is provided for completeness, but
it's not recommended. Typically, a logger should be created and used unchanged throughout a process life-cycle eventually
changing its log level.

```go
package main

import (
    "os"

    "github.com/gomatbase/go-log"
)

func main() {
    // First time getting a logger named "TESTER", logger will be created
    logger, e := log.GetWithOptions("TESTER",
        log.WithOptions().
            DateFlags(log.Ldate).        // Set the log flags following the same options as the standard log package
            WithWriter(os.Stderr).       // Write logs to another writer (stderr in this case)
            WithoutFailingCriticals().   // Write a critical log as a failure, panicking
            WithFailingCriticals().      // Write a critical log as a plain log entry, not causing the process to exit
            WithStartingLevel(log.INFO)) // Set the starting log level to INFO
    if e != nil {
        log.Critical("Unable to create custom logger TESTER : ", e)
        os.Exit(1)
    }
    
    logger.Critical("A critical, non-failing error entry")
    logger.Info("Logging information")

    logger, _ = log.Get("TESTER")
    logger.Error("Logging an error")
    logger.Info("Logging information")
    
    _, _ = log.OverrideLogWithOptions("TESTER",
        log.WithOptions().
            DateFlags(0).
            WithFailingCriticals().
            WithStartingLevel(log.ERROR))

    logger.Info("Logging some more info")
    logger.Critical("Fatal critical, exiting...")

    // output (assuming date 15/06/2021) :
    // TESTER - 2021/06/15 A critical, non-failing error entry
    // TESTER - 2021/06/15 Logging information
    // TESTER - 2021/06/15 Logging an error
    // TESTER - 2021/06/15 Logging information
    // TESTER - Fatal critical, exiting...
}
```

### Log Levels

The log level is a simple `uint` which the logger uses to compare with the level of a log entry and will print any log entry whose level is
less or equal than the logger level.

Constants are provided with names for the severity of the intended log. Some namesake functions are also provided to allow logging for the level without 
explicitly stating the log level in every call. The provided constants by decreasing order of severity are:

0. CRITICAL (it will always log the entry)
    * `log.Critical(log ...interface{})`
    * `log.Criticalf(format string, variables ...interface{})`
0. ERROR
    * `log.Error(log ...interface{})`
    * `log.Errorf(format string, variables ...interface{})`
0. WARNING (default value for any newly created logger unless specified otherwise)
    * `log.Warning(log ...interface{})`
    * `log.Warningf(log string, variables ...interface{})`
0. INFO
    * `log.Info(log ...interface{})`
    * `log.Infof(format string, variables ...interface{})`
0. DEBUG
    * `log.Debug(log ...interface{})`
    * `log.Debugf(format string, variables ...interface{})`
0. TRACE
    * `log.Trace(log ...interface{})`
    * `log.Tracef(format string, variables ...interface{})`
    
These constants and methods are provided for convenience but the level can be specified
as any unsigned integer, and using `log.Println(level uint, log ...interface{})` and 
`log.Printf(level uint, format string, variables ...interface{})` any kind of level range may be used.
