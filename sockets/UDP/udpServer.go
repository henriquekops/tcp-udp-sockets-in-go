package UDP

import (
	c "../common"
	"fmt"
	"net"
	"strings"
)

const serverType  = "UDP:Server"

func CreateUDPServer(network string, serverPort string) {

	// Startup
	fmt.Println("Starting up " + serverType + " ...")

	udpAddr, err := net.ResolveUDPAddr(network, serverPort)
	c.CheckError(serverType, err)

	connection, err := net.ListenUDP(network, udpAddr)
	c.CheckError(serverType, err)
	defer connection.Close()

	buffer := make([]byte, c.BUFFER_SIZE)
	c.Log(serverType, "Up and running!\n[Crtl+C to quit]")

	// Wait for UDP connection
	for {
		length, remoteAddr, err := connection.ReadFromUDP(buffer)
		c.CheckError(serverType, err)

		clientMsg := strings.TrimSuffix(strings.TrimSpace(string(buffer[:length])), "\n")
		parsedAddr := remoteAddr.String()

		// Aborted ?
		if clientMsg == "exit" {
			c.Log(serverType, "CONNECTION ABORTED -> '"+parsedAddr+"'")
			return
		}

		c.Log(serverType, "RECEIVED (SOURCE='"+parsedAddr+"') -> '"+clientMsg+"'")
	}
}