# Logger Package

A simple and customizable logging package for Go.

## Features
- Log messages with different levels: `INFO`, `WARNING`, `ERROR`, and `FATAL`.
- Colored output for better visibility in terminal.
- Graceful handling of fatal errors with custom exit codes.

## Installation

To install the package, run the following command:

```bash
go get github.com/StarGames2025/Logger
```

## Usage

Below is an example of how to use the Logger package in your Go project:

```go
package main

import (
    "github.com/StarGames2025/Logger"
    "time"
)

func main() {
    // Create a new logger instance with INFO level
    logger := Logger.NewLogger(Logger.INFO)

    // Log messages with different levels
    logger.Info("This is an info message.")
    logger.Warning("This is a warning message.")
    logger.Error("This is an error message.")
    
    // Log a fatal error with a custom exit code
    logger.Fatal("This is a fatal error message.", "DB_INIT_ERROR")
}
```

### Log Levels
- **INFO**: Standard informational messages.
- **WARNING**: Warnings about potential issues.
- **ERROR**: Errors that need attention but do not cause the program to stop.
- **FATAL**: Critical errors that cause the program to stop.

### Exit Codes
The logger handles different exit codes for fatal errors. You can specify a custom exit code when calling `Fatal()`.

Available exit codes are:
- `SUCCESS` (0)
- `DB_INIT_ERROR` (10)
- `DEVICE_STORE_ERROR` (11)
- and more...

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to fork the repository and open a pull request for improvements or bug fixes. Contributions are always welcome!

## Author

DevStarByte