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
	INFO LogLevel = iota
	WARNING
	ERROR
	FATAL
)

// Mapping of custom exit codes for different error types
var exitCodes = map[string]int{
	"SUCCESS":              0,
	"DB_INIT_ERROR":        10,
	"DEVICE_STORE_ERROR":   11,
	"CLIENT_CONNECT_ERROR": 12,
	"QR_GENERATE_ERROR":    13,
	"QR_OPEN_ERROR":        14,
	"QR_DECODE_ERROR":      15,
	"QR_RESIZE_ERROR":      16,
	"QR_FILE_CREATE_ERROR": 17,
	"QR_FILE_ENCODE_ERROR": 18,
	"QR_RENDER_ERROR":      19,
	"GROUP_FETCH_ERROR":    20,
	"CONTACT_FETCH_ERROR":  21,
	"SHUTDOWN":             99,
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
			// Pass the corresponding exit code for SUCCESS, which is 0
			l.handleFatal(exitCodes["SUCCESS"])
		}
	}
}

// Info logs a message with INFO level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

// Warning logs a message with WARNING level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Warning(msg string) {
	l.log(WARNING, msg)
}

// Error logs a message with ERROR level.
// Parameters:
// - msg: The log message to be displayed.
func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}

// Fatal logs a message with FATAL level and exits the program with the corresponding exit code.
// Parameters:
// - msg: The log message to be displayed.
// - exitCodeName: The name of the exit code to be used from the exitCodes map.
func (l *Logger) Fatal(msg string, exitCodeName string) {
	l.log(FATAL, msg)

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
