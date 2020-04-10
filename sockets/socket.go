package sockets

import (
	"./tcp"
	"./udp"
	"fmt"
	env "github.com/joho/godotenv"
	"os"
)

func getEnvParams() (host string, port string) {
	if err := env.Load(".env"); err != nil {
		fmt.Println("Where is .env?")
		os.Exit(1)
	}
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
	return
}

func GetSocket(socketType string, networkType string)  {
	host, port := getEnvParams()
	if socketType == "client" && networkType == "tcp" {
		tcp.CreateTCPClient(networkType, host+":"+port)
	} else if socketType == "server" && networkType == "tcp" {
		tcp.CreateTCPServer(networkType, ":"+port)
	} else if socketType == "client" && networkType == "udp" {
		udp.CreateUDPClient(networkType, host+":"+port)
	} else if socketType == "server" && networkType == "udp" {
		udp.CreateUDPServer(networkType, ":"+port)
	} else {
		fmt.Println("Your entry is invalid")
		os.Exit(1)
	}
}