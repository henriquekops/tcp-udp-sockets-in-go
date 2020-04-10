package UDP

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
const clientType = "UDP:Client"

// CreateUDPClient creates new UDP connection
func CreateUDPClient(network string, serverAddress string) {

	// Startup
	fmt.Println("Starting up "+clientType+" ...")

	udpAddr, err := net.ResolveUDPAddr(network, serverAddress)
	c.CheckError(clientType, err)

	connection, err := net.DialUDP(network, nil, udpAddr)
	c.CheckError(clientType, err)
	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)
	c.Log(clientType, "Connected to '"+connection.RemoteAddr().String()+"'")

	// Read from cmdline
	c.Log(clientType, "Input file path to send [HINT: use './test/send/test.txt']")
	fmt.Print("[PATH]: ")
	filePath, _ := reader.ReadString('\n')

	// Open file
	file, err := os.Open(strings.TrimSpace(strings.TrimSuffix(filePath, "\n")))
	if err != nil {
		c.Log(clientType, "File does not exist!")
		return
	}
	defer file.Close()

	// Send file to TCP server
	fileInfo, _ := file.Stat()
	_, err =  io.Copy(connection, file)
	c.CheckError(clientType, err)

	// Log
	formatBytes := strconv.FormatInt(fileInfo.Size(), 10)
	c.Log(clientType, "SENT "+formatBytes+" BYTES (DEST='"+serverAddress+"')")
}