package main

import (
	"log"
	chat "netcat/functions"
)

func main() {
	server := chat.NewServer(":3000")
	log.Fatal(server.Start())
}
