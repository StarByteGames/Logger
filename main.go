package Logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

type LogLevel int

// Defining different log levels
const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Mapping of custom exit codes for different error types
var exitCodes = map[string]int{
	"ERROR":    -1,
	"SHUTDOWN": 0,
	"SUCCESS":  0,
}

// Logger struct holds the log level and methods to log messages
type Logger struct {
	level LogLevel
}

// NewLogger creates a new Logger instance with the provided log level.
// Parameters:
// - level: The minimum log level the logger should display.
// Returns:
// - A pointer to a new Logger instance.
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level: level,
	}
}

// log is the core logging function. It prints log messages with a timestamp,
// log level, and color according to the specified log level.
// Parameters:
// - level: The log level for the message (INFO, WARNING, ERROR, FATAL).
// - msg: The log message to be displayed.
func (l *Logger) log(level LogLevel, msg string) {
	if level >= l.level {
		levelString := ""
		var levelColor *color.Color

		// Assign the appropriate color for the log level
		switch level {
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

		// Set the timestamp to white
		timestamp := color.New(color.FgWhite)

		// Print timestamp and log message with colors
		fmt.Printf("[%s] %s: %s\n", timestamp.Sprint(time.Now().Format("2006-01-02 15:04:05")), levelColor.Sprint(levelString), msg)

		// If Fatal, exit the program
		if level == FATAL {
			// Pass the corresponding exit code for ERROR, which is -1
			l.handleFatal(exitCodes["ERROR"])
		}
	}
}

// Info logs a message with INFO level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Info(msg ...string) {
	message := ""
	for _, part := range msg {
		message += part + " "
	}
	l.log(INFO, message)
}

// Warning logs a message with WARNING level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Warning(msg ...string) {
	message := ""
	for _, part := range msg {
		message += part + " "
	}
	l.log(WARNING, message)
}

// Debug logs a message with DEBUG level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Debug(msg ...string) {
	message := ""
	for _, part := range msg {
		message += part + " "
	}
	l.log(DEBUG, message)
}

// Error logs a message with ERROR level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Error(msg ...string) {
	message := ""
	for _, part := range msg {
		message += part + " "
	}
	l.log(ERROR, message)
}

// Fatal logs a message with FATAL level and exits the program with the corresponding exit code.
// Parameters:
// - exitCodeName: The name of the exit code to be used from the exitCodes map.
// - msg: The log message to be displayed.
func (l *Logger) Fatal(exitCodeName string, msg ...string) {
	message := ""
	for _, part := range msg {
		message += part + " "
	}
	l.log(FATAL, message)

	// Fetch the exit code from the map by its name
	exitCode, exists := exitCodes[exitCodeName]
	if !exists {
		// If the exit code name is not valid, use "SUCCESS" (0) as a fallback
		log.Printf("Invalid exit code name. Defaulting to 'SUCCESS' (0).\n")
		exitCode = exitCodes["SUCCESS"]
	}

	// Handle fatal error by exiting the program with the specified exit code
	l.handleFatal(exitCode)
}

// handleFatal is responsible for handling fatal errors. It performs any necessary cleanup
// and then exits the program using the specified exit code.
// Parameters:
// - exitCode: The exit code to be used when exiting the program.
func (l *Logger) handleFatal(exitCode int) {
	// Simulate the handling of a fatal error (e.g., triggering cleanup, etc.)
	log.Println("A fatal error occurred. Exiting...")

	// You could also do cleanup here if needed before exiting

	// Exit the program with the provided exit code
	os.Exit(exitCode)
}
