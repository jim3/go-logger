package pocketlog

// Level is used to define the level of logging.
type Level byte

// The following constants are used to define the level of logging.
const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable.
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace. errors
	LevelError
)
