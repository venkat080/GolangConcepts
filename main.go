package main

import (
	"golangconcepts/concepts"
)

func main() {
	// Embedding //
	// concepts.Embedding()

	////Goroutines//
	// concepts.GoRoutines()

	//// channels: Example 1//
	// concepts.Channels()

	//// Race conditions cannot be seen in output as small program
	concepts.ChannelsWithWaitGroup()

	//// We are forcing the raceconditions here and it can be seen in this
	concepts.ChannelsWithMutexExampleClearUnderstanding()

	//// RWMutex
	concepts.ChannelsWithRWMutex()

	// Select statement
	concepts.ChannelsWithSelectStatement()
}
