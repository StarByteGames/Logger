# Logger

A simple, customizable logging package for Go with colored console output and file logging.

## Features

- Multiple log levels: `DEBUG`, `INFO`, `WARNING`, `ERROR`, `FATAL`
- Colored output for terminal logs
- Log messages to both file and console
- Customizable exit codes for fatal errors
- Automatic file cleanup via Go's garbage collector

## Installation

Add the package to your project:

```sh
go get github.com/StarGames2025/Logger
```

## Usage

```go
package main

import (
    "github.com/StarGames2025/Logger"
)

func main() {
    // Create a new logger (log level, log file path, log to console)
    logger, err := Logger.NewLogger(Logger.INFO, "app.log", true)
    if err != nil {
        panic(err)
    }
    defer logger.Close()

    logger.Debug("This is a debug message")
    logger.Info("Application started")
    logger.Warning("This is a warning")
    logger.Error("An error occurred")
    logger.Fatal("SHUTDOWN", "Fatal error, shutting down")
}
```

## Log Levels

- `DEBUG`: Detailed debug information
- `INFO`: General information
- `WARNING`: Warnings about potential issues
- `ERROR`: Errors that do not stop the program
- `FATAL`: Critical errors that terminate the program

## Exit Codes

Customize exit codes via the `ExitCodes` map in the logger. By default:
- `"ERROR"`: -1
- `"SHUTDOWN"`: 0
- `"SUCCESS"`: 0

You can add more codes as needed.

## License

MIT License. See [LICENSE](LICENSE) for details.