package UDP

import (
	c "../common"
	"fmt"
	"net"
	"os"
	"strconv"
)

// serverType defines server identification when logging
const serverType = "UDP:Server"

// CreateUDPServer create server and listen to UDP
func CreateUDPServer(network string, serverPort string) {

	// Startup
	fmt.Println("Starting up "+serverType+" ...")

	listener, err := net.ListenPacket(network, serverPort)
	c.CheckError(serverType, err)
	defer listener.Close()

	c.Log(serverType, "Up and running!\n[Crtl+C to quit]")

	// Listening to UDP
	for {
		buffer := make([]byte, c.BUFFER_SIZE)
		length, remoteAddr, err := listener.ReadFrom(buffer)
		c.CheckError(serverType, err)

		// Incoming request!
		handleClient(remoteAddr.String(), buffer[:length])
	}
}

// handleClient handles incoming UDP client connections
func handleClient(remoteAddr string, buffer []byte) {
	// Create new file
	file, err := os.Create("./test/receive/UDP_RECEIVED.txt")
	c.CheckError(serverType, err)
	defer file.Close()

	// Copy file data
	_, err = file.Write(buffer)
	c.CheckError(serverType, err)

	// Log
	fileInfo, _ := file.Stat()
	formatBytes := strconv.FormatInt(fileInfo.Size(), 10)
	c.Log(serverType, "RECEIVED "+formatBytes+" BYTES (SOURCE='"+remoteAddr+"')")
}