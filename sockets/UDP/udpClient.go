package UDP

import (
	c "../common"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const clientType = "UDP:Client"

func CreateUDPClient(network string, serverAddress string) {

	// Startup
	fmt.Println("Starting up " + clientType + " ...")

	udpAddr, err := net.ResolveUDPAddr(network, serverAddress)
	c.CheckError(clientType, err)

	connection, err := net.DialUDP(network, nil, udpAddr)
	c.CheckError(clientType, err)
	defer connection.Close()

	c.Log(clientType, "Connected to '" + connection.RemoteAddr().String() + "'")
	c.Log(clientType, "Input your data to send to server\n['exit' to quit]")

	for {
		// Args
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')

		// Write to UDP server
		data := []byte(input)
		_, err := connection.Write(data)
		c.CheckError(clientType, err)

		// Exit?
		if strings.TrimSpace(input) == "exit" {
			c.Log(clientType, "Safe quit... Bye :)")
			return
		}
	}
}