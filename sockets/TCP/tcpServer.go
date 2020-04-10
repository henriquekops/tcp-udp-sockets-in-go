package TCP

import (
	c "../common"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
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
	buffer := make([]byte, c.BUFFER_SIZE)
	remoteAddr := connection.RemoteAddr().String()
	c.Log(serverType, "NEW CONNECTION (SOURCE=''" + remoteAddr + "')")
	defer connection.Close()

	for {
		// Capture message
		length, err := connection.Read(buffer)
		c.CheckError(serverType, err)
		clientMsg := strings.TrimSuffix(strings.TrimSpace(string(buffer[:length])), "\n")

		// Aborted ?
		if clientMsg == "exit" {
			c.Log(serverType, "Connection " + remoteAddr + " aborted")
			break
		}

		// File received
		fileName, fileSize := receiveFile(clientMsg, connection)
		c.Log(serverType, "RECEIVED (SOURCE='"+ remoteAddr +"') -> '" + fileName + ", " + strconv.FormatInt(fileSize, 10) + " BYTES")

		// Return to client
		c.Log(serverType, "SENDING (SOURCE='" + remoteAddr + "') OK")
		_, err = connection.Write([]byte("OK\n"))
		c.CheckError(serverType, err)
	}
}

func receiveFile(filename string, conn net.Conn) (string, int64){
	// Create new file
	file, err := os.Create(filename)
	c.CheckError(serverType, err)
	defer file.Close()

	// Copy file data
	_, err = io.Copy(conn, file)
	c.CheckError(serverType, err)
	fileInfo, _ := file.Stat()

	return fileInfo.Name(), fileInfo.Size()
}
