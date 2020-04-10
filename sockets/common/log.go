package common

import (
	"fmt"
	"os"
	"time"
)

// Socket struct keeps socket type string for logging
type Socket struct {
	SocketType string
}

// BUFFERSIZE defines communication buffer size
const BUFFERSIZE = 1024
// TIMESTAMP defines timestamp log format
const TIMESTAMP = "02 Jan 06 15:04"

// Log creates common log format
func (socket *Socket) Log(logMessage string) {
	t := time.Now()
	fmt.Printf("[%s] %s - %s\n",
		socket.SocketType,
		t.Format(TIMESTAMP),
		logMessage,
	)
}

// CheckError creates common error validation
func (socket *Socket) CheckError(err error) {
	if err != nil {
		t := time.Now()
		fmt.Printf("[%s] %s - Exception '%s', exiting...",
			socket.SocketType,
			t.Format(TIMESTAMP),
			err.Error(),
		)
		os.Exit(1)
	}
}
