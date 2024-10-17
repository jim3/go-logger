package pocketlog_test

// To access pocketlog functions, we need to import it:
import "learngo-pockets/logger/pocketlog"

// test the standard output
func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}

// ----------------------------------------------
