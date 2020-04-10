package tcp

import (
	c "../common"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

// TCPServerType defines server identification when logging
const TCPServerType = "TCP:Server"
// savePath defines path for saving received file data
const savePath = "./test/receive/TCP_RECEIVED.txt"

// CreateTCPServer create server and listen to TCP
func CreateTCPServer(network string, serverPort string) {

	// Startup
	fmt.Println("Starting up "+TCPServerType+" ...")

	tcpSocket := c.Socket{TCPServerType}

	listener, err := net.Listen(network, serverPort)
	tcpSocket.CheckError(err)
	defer listener.Close()

	tcpSocket.Log("Up and running!\n[Crtl+C to quit]")

	// Listening to TCP
	for {
		connection, err := listener.Accept()
		tcpSocket.CheckError(err)

		// Incoming request!
		go handleClient(tcpSocket, connection)
	}
}

// handleClient handles incoming TCP client connections
func handleClient(tcpSocket c.Socket, connection net.Conn) {
	// New connection
	remoteAddr := connection.RemoteAddr().String()
	defer connection.Close()

	// Create new file
	file, err := os.Create(savePath)
	tcpSocket.CheckError(err)
	defer file.Close()

	// Copy file data
	_, err = io.Copy(file, connection)
	tcpSocket.CheckError(err)

	// Log
	fileInfo, _ := file.Stat()
	formatBytes := strconv.FormatInt(fileInfo.Size(), 10)
	tcpSocket.Log("RECEIVED "+formatBytes+" BYTES (SOURCE='"+remoteAddr+"')")
}