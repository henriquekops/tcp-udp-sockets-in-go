package TCP

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

const clientType = "TCP:Client"

func CreateTCPClient(network string, serverAddress string) {

	// Startup
	fmt.Println("Starting up " + clientType + " ...")

	connection, err := net.Dial(network, serverAddress)
	c.CheckError(clientType, err)
	defer connection.Close()

	c.Log(clientType, "Connected to '" + connection.RemoteAddr().String() + "'")
	c.Log(clientType, "Input your data to send to server\n['exit' to quit]")

	for {
		// Args
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		parsedInput := strings.TrimSpace(input)

		// Exit?
		if parsedInput == "exit" {
			_, err = fmt.Fprintf(connection, parsedInput + "\n")
			c.CheckError(clientType, err)
			c.Log(clientType, "Safe quit... Bye :)")
			return
		}

		// Send file to TCP server
		fileName, fileSize, err := sendFileToServer(connection, parsedInput)
		if err == nil {
			c.Log(clientType, "SENT " + fileName + ", " + strconv.FormatInt(fileSize, 10) + " BYTES")

			// Wait for response
			serverMsg, _ := bufio.NewReader(connection).ReadString('\n')
			parsedMsg := strings.TrimSuffix(strings.TrimSpace(serverMsg),  "\n")
			c.Log(clientType, "RECEIVED -> '" + parsedMsg + "'")
		}
	}
}

func sendFileToServer(connection net.Conn, filename string) (string, int64, error) {
	// Open file
	buffer := make([]byte, c.BUFFER_SIZE)

	file, err := os.Open(filename)
	if err != nil {
		c.Log(clientType, "File does not exist!")
		return "", 0, err
	}
	defer file.Close()

	// Send file
	fileInfo, _ := file.Stat()

	_, err = connection.Write(buffer)
	c.CheckError(clientType, err)

	_, err =  io.Copy(connection, file)
	c.CheckError(clientType, err)

	return fileInfo.Name(), fileInfo.Size(), nil
}
