package netcat

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
	length     int
	maxClients int
	mu         sync.Mutex
	clients    map[net.Conn]string
	clientMu   sync.Mutex
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
		length:     0,
		maxClients: 10,
		clients:    make(map[net.Conn]string),
	}
}

func WelcomeMessage() []byte {
	file, err := os.Open("/assets/welcome.txt")
	if err != nil {
		log.Fatal("Error opening txt file")
	}
	defer file.Close()
	result, _ := io.ReadAll(file)
	return result
}