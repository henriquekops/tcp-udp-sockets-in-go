package tcp

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
const TCPClientType = "tcp:Client"

// CreateTCPClient creates new tcp connection
func CreateTCPClient(network string, serverAddress string) {

	// Startup
	fmt.Println("Starting up "+TCPClientType+" ...")

	tcpSocket := c.Socket{TCPClientType}

	connection, err := net.Dial(network, serverAddress)
	tcpSocket.CheckError(err)
	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)
	tcpSocket.Log("Connected to '"+connection.RemoteAddr().String()+"'")

	// Read from cmdline
	tcpSocket.Log("Input file path to send [HINT: use './test/send/test.txt']")
	fmt.Print("[PATH]: ")
	filePath, _ := reader.ReadString('\n')

	// Open file
	file, err := os.Open(strings.TrimSpace(strings.TrimSuffix(filePath, "\n")))
	if err != nil {
		tcpSocket.Log("File does not exist!")
		return
	}
	defer file.Close()

	// Send file to tcp server
	fileInfo, _ := file.Stat()
	_, err =  io.Copy(connection, file)
	tcpSocket.CheckError(err)

	// Log
	formatBytes := strconv.FormatInt(fileInfo.Size(), 10)
	tcpSocket.Log("SENT "+formatBytes+" BYTES (DEST='"+serverAddress+"')")
}