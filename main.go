package Logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

// LogLevel represents the severity level of a log message
type LogLevel int

// Defining different log levels
const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger struct holds the log level, file writer, console flag, and exit codes
type Logger struct {
	level        LogLevel
	logFile      *os.File
	logToConsole bool
	ExitCodes    map[string]int
}

// NewLogger creates a new Logger instance with the provided log level and file path.
// Automatically sets a finalizer to close the file when the logger is garbage collected.
// Parameters:
// - level: The minimum log level the logger should display.
// - logFilePath: The path to the log file.
// - logToConsole: Whether to also print logs to the terminal.
// Returns:
// - A pointer to a Logger instance and an error if file creation fails.
func NewLogger(level LogLevel, logFilePath string, logToConsole bool) (*Logger, error) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := &Logger{
		level:        level,
		logFile:      file,
		logToConsole: logToConsole,
		ExitCodes: map[string]int{
			"ERROR":    -1,
			"SHUTDOWN": 0,
			"SUCCESS":  0,
		},
	}

	// Automatically close the log file when the logger is garbage collected
	runtime.SetFinalizer(logger, func(l *Logger) {
		fmt.Println("Finalizer: Closing log file.")
		l.Close()
	})

	return logger, nil
}

// Close closes the log file.
// Should be called when logging is no longer needed.
func (l *Logger) Close() {
	if l.logFile != nil {
		l.logFile.Close()
	}
}

// log is the core logging function. It prints log messages with a timestamp,
// log level, and color (to console) according to the specified log level.
// Parameters:
// - level: The log level for the message (DEBUG, INFO, WARNING, ERROR, FATAL).
// - msg: The log message to be displayed.
func (l *Logger) log(level LogLevel, msg string) {
	if level < l.level {
		return
	}

	var levelString string
	var levelColor *color.Color

	// Assign the appropriate color for the log level
	switch level {
	case DEBUG:
		levelString = "DEBUG"
		levelColor = color.New(color.FgCyan) // Cyan for DEBUG
	case INFO:
		levelString = "INFO"
		levelColor = color.New(color.FgGreen) // Green for INFO
	case WARNING:
		levelString = "WARNING"
		levelColor = color.New(color.FgYellow) // Yellow for WARNING
	case ERROR:
		levelString = "ERROR"
		levelColor = color.New(color.FgRed) // Red for ERROR
	case FATAL:
		levelString = "FATAL"
		levelColor = color.New(color.FgMagenta) // Magenta for FATAL
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] %s: %s\n", timestamp, levelString, msg)

	// Write to file (without color)
	if l.logFile != nil {
		l.logFile.WriteString(logLine)
	}

	// Print to console (with color)
	if l.logToConsole {
		fmt.Printf("[%s] %s: %s\n",
			color.New(color.FgWhite).Sprint(timestamp),
			levelColor.Sprint(levelString),
			msg,
		)
	}

	// Exit if level is FATAL
	if level == FATAL {
		l.handleFatal(l.ExitCodes["ERROR"])
	}
}

// Info logs a message with INFO level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Info(msg ...string) {
	l.log(INFO, join(msg))
}

// Warning logs a message with WARNING level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Warning(msg ...string) {
	l.log(WARNING, join(msg))
}

// Debug logs a message with DEBUG level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Debug(msg ...string) {
	l.log(DEBUG, join(msg))
}

// Error logs a message with ERROR level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Error(msg ...string) {
	l.log(ERROR, join(msg))
}

// Fatal logs a message with FATAL level and exits the program with the corresponding exit code.
// Parameters:
// - exitCodeName: The name of the exit code to be used from the ExitCodes map.
// - msg: The log message to be displayed.
func (l *Logger) Fatal(exitCodeName string, msg ...string) {
	message := join(msg)
	l.log(FATAL, message)

	// Fetch the exit code from the map by its name
	exitCode, exists := l.ExitCodes[exitCodeName]
	if !exists {
		// If the exit code name is not valid, use "SUCCESS" (0) as a fallback
		log.Printf("Invalid exit code name. Defaulting to 'SUCCESS' (0).\n")
		exitCode = l.ExitCodes["SUCCESS"]
	}

	// Handle fatal error by exiting the program with the specified exit code
	l.handleFatal(exitCode)
}

// handleFatal is responsible for handling fatal errors. It performs any necessary cleanup
// and then exits the program using the specified exit code.
// Parameters:
// - exitCode: The exit code to be used when exiting the program.
func (l *Logger) handleFatal(exitCode int) {
	log.Println("A fatal error occurred. Exiting...")
	l.Close()
	os.Exit(exitCode)
}

// join joins multiple strings with spaces.
// Parameters:
// - parts: Variadic string parts to concatenate.
// Returns:
// - A single string with all parts separated by spaces.
func join(parts []string) string {
	result := ""
	for _, part := range parts {
		result += part + " "
	}
	return result
}
