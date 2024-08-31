package main

import (
	"fmt"
	"log"
	"os"

	chat "netcat/functions"
)

func isValidPort(input string) bool {
	if len(input) == 0 {
		return false
	}

	for i := 0; i < len(input); i++ {
		if input[i] < '0' || input[i] > '9' {
			return false
		}
	}

	port := 0
	for i := 0; i < len(input); i++ {
		port = port*10 + int(input[i]-'0')
	}

	return port >= 0 && port <= 65535
}

func main() {
	if len(os.Args) < 1 || len(os.Args) > 2 {
		fmt.Println("Error, the program runs as the following: go run . [PORT]")
		return
	}

	var port string = "8989"
	if len(os.Args) == 2 {
		if isValidPort(os.Args[1]) {
			port = os.Args[1]
		} else {
			fmt.Println("Error: invalid port number")
			return
		}
	}

	server := chat.NewServer(string(":" + port))
	log.Fatal(server.Start())
}
