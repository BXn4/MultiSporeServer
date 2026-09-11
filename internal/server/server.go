package server

import (
	"fmt"
	"multispore/internal/client"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/charmbracelet/log"
)

type ServerConfig struct {
	Host string
	Port string
	Name string
}

type Server struct {
	config  *ServerConfig
	clients map[string]net.Conn
	mu      sync.Mutex
}

func New(config *ServerConfig) (*Server, error) {
	return &Server{
		config:  config,
		clients: make(map[string]net.Conn),
	}, nil
}

func (s *Server) Run() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Infof("Server: [%s] started, and listening on %s", s.config.Name, address)

	// Handle connections
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			defer conn.Close()

			c := client.New(conn)
			log.Infof("Client connected %s ", c.GetIP())
			/*c.SetClientID(s.gm.NextClientID())
			s.gm.AddClient(c)
			c.Start()
			go commands.HandleClient(c, s.gm)*/
		}
	}()
	// Wait for interrupt signal
	<-sigChan
	log.Info("Received interrupt signal. Saving all data...")
}
