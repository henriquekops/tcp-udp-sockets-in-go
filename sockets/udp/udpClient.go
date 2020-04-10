package udp

import (
	c "../common"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

// clientType defines client identification when logging
const UDPClientType = "UDP:Client"

// CreateUDPClient creates new UDP connection
func CreateUDPClient(network string, serverAddress string) {

	// Startup
	fmt.Println("Starting up "+UDPClientType+" ...")

	udpSocket := c.Socket{UDPClientType}

	udpAddr, err := net.ResolveUDPAddr(network, serverAddress)
	udpSocket.CheckError(err)

	connection, err := net.DialUDP(network, nil, udpAddr)
	udpSocket.CheckError(err)
	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)
	udpSocket.Log("Connected to '"+connection.RemoteAddr().String()+"'")

	// Read from cmdline
	udpSocket.Log("Input file path to send [HINT: use './test/send/test.txt']")
	fmt.Print("[PATH]: ")
	filePath, _ := reader.ReadString('\n')

	// Open file
	file, err := os.Open(strings.TrimSpace(strings.TrimSuffix(filePath, "\n")))
	if err != nil {
		udpSocket.Log("File does not exist!")
		return
	}
	defer file.Close()

	// Send file to TCP server
	fileInfo, _ := file.Stat()
	_, err =  io.Copy(connection, file)
	udpSocket.CheckError(err)

	// Log
	formatBytes := strconv.FormatInt(fileInfo.Size(), 10)
	udpSocket.Log("SENT "+formatBytes+" BYTES (DEST='"+serverAddress+"')")
}