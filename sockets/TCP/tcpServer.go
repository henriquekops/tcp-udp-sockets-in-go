package TCP

import (
	c "../common"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

const serverType  = "TCP:Server"

func CreateTCPServer(network string, serverPort string) {

	// Startup
	fmt.Println("Starting up " + serverType + " ...")

	listener, err := net.Listen(network, serverPort)
	c.CheckError(serverType, err)
	defer listener.Close()

	c.Log(serverType, "Up and running!\n[Crtl+C to quit]")

	// Listening to TCP
	for {
		connection, err := listener.Accept()
		c.CheckError(serverType, err)

		// Incoming request!
		go handleClient(connection)
	}
}

func handleClient(connection net.Conn) {
	// New connection
	remoteAddr := connection.RemoteAddr().String()
	c.Log(serverType, "NEW CONNECTION (SOURCE=''" + remoteAddr + "')")
	defer connection.Close()

	// Create new file
	file, err := os.Create("./test/receive/TCP_RECEIVED.txt")
	c.CheckError(serverType, err)
	defer file.Close()

	// Copy file data
	_, err = io.Copy(file, connection)
	c.CheckError(serverType, err)

	// Log
	fileInfo, _ := file.Stat()
	c.Log(serverType, "RECEIVED " + strconv.FormatInt(fileInfo.Size(), 10) + " BYTES")
	c.Log(serverType, "CLOSING CONNECTION WITH " + remoteAddr)
}