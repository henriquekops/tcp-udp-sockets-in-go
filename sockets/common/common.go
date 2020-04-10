package common

import (
	"fmt"
	"os"
	"time"
)

// BUFFER_SIZE defines communication buffer size
const BUFFER_SIZE = 1024
// TIMESTAMP defines timestamp log format
const TIMESTAMP = "02 Jan 06 15:04"

// Log creates common log format
func Log(socketType string, logMessage string) {
	t := time.Now()
	fmt.Printf("[%s] %s - %s\n",
		socketType,
		t.Format(TIMESTAMP),
		logMessage,
	)
}

// CheckError creates common error validation
func CheckError(socketType string, err error) {
	if err != nil {
		t := time.Now()
		fmt.Printf("[%s] %s - Exception '%s', exiting...",
			socketType,
			t.Format(TIMESTAMP),
			err.Error(),
		)
		os.Exit(1)
	}
}
