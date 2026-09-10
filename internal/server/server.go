package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
)

type ServerConfig struct {
	Host string
	Port string
}

type Server struct {
	config *ServerConfig
}

func New(config *ServerConfig) (*Server, error) {
	return &Server{
		config: config,
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

	log.Infof("Server started, and listening on %s", address)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			defer conn.Close()

			//c := client.New(conn, db, s.gm)
			//c.SetClientID(s.gm.NextClientID())
			//s.gm.AddClient(c)
			//c.Start()
			//go commands.HandleClient(c, s.gm)
		}
	}()
	// Wait for interrupt signal
	<-sigChan
	log.Info("Received interrupt signal. Saving all data...")
}
