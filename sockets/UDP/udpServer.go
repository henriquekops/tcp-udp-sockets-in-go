package udp

import (
	c "../common"
	"fmt"
	"net"
	"os"
	"strconv"
)

// UDPServerType defines server identification when logging
const UDPServerType = "udp:Server"
// savePath defines path for saving received file data
const savePath = "./test/receive/UDP_RECEIVED.txt"

// CreateUDPServer create server and listen to udp
func CreateUDPServer(network string, serverPort string) {

	// Startup
	fmt.Println("Starting up "+UDPServerType+" ...")

	udpSocket := c.Socket{UDPServerType}

	listener, err := net.ListenPacket(network, serverPort)
	udpSocket.CheckError(err)
	defer listener.Close()

	udpSocket.Log("Up and running!\n[Crtl+C to quit]")

	// Listening to udp
	for {
		buffer := make([]byte, c.BUFFERSIZE)
		length, remoteAddr, err := listener.ReadFrom(buffer)
		udpSocket.CheckError(err)

		// Incoming request!
		handleClient(udpSocket, remoteAddr.String(), buffer[:length])
	}
}

// handleClient handles incoming udp client connections
func handleClient(udpSocket c.Socket, remoteAddr string, buffer []byte) {
	// Create new file
	file, err := os.Create(savePath)
	udpSocket.CheckError(err)
	defer file.Close()

	// Copy file data
	_, err = file.Write(buffer)
	udpSocket.CheckError(err)

	// Log
	fileInfo, _ := file.Stat()
	formatBytes := strconv.FormatInt(fileInfo.Size(), 10)
	udpSocket.Log("RECEIVED "+formatBytes+" BYTES (SOURCE='"+remoteAddr+"')")
}