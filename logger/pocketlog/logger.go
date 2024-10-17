package pocketlog

// Logger is used to log information.
type Logger struct {
	threshold Level
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// A receiver method is a function that operates on the `struct` (structure) specified in parentheses before the function’s name
// Receiver methods can accept a copy or a reference - a pointer - to the structure

// Debugf formats and logs a message at the debug level or higher
func (l *Logger) Debugf(format string, args ...any) {
	// implementation
}

// Infof formats and logs a message at the info level or higher
func (l *Logger) Infof(format string, args ...any) {
	// implementation
}

// Errorf formats and logs a message at the error level or higher
func (l *Logger) Errorf(format string, args ...any) {
	// implementation
}
