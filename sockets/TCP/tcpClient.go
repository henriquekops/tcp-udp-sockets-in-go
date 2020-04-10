package TCP

import (
	c "../common"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

const clientType = "TCP:Client"

func CreateTCPClient(network string, serverAddress string) {

	// Startup
	fmt.Println("Starting up " + clientType + " ...")

	connection, err := net.Dial(network, serverAddress)
	c.CheckError(clientType, err)
	defer connection.Close()

	c.Log(clientType, "Connected to '" + connection.RemoteAddr().String() + "'")

	// Open file
	file, err := os.Open("./test/send/test1.txt")
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
	c.Log(clientType, "SENT " + strconv.FormatInt(fileInfo.Size(), 10) + " BYTES")
}