package server

import (
	"fmt"
	"multispore/internal/client"
	"multispore/internal/commands"
	"multispore/internal/database"
	"multispore/internal/managers"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
)

type ServerConfig struct {
	Host string
	Port string
	Name string
}

type Server struct {
	config   *ServerConfig
	dbConfig *database.DBConfig
	maxConn  int
	db       *database.Database
	gm       *managers.GameManager
}

func New(config *ServerConfig, dbConfig *database.DBConfig) (*Server, error) {
	gm, err := managers.NewGameManager()
	if err != nil {
		return nil, err
	}
	return &Server{
		config:   config,
		dbConfig: dbConfig,
		gm:       gm,
	}, nil
}

func (s *Server) Run() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	db, err := database.ConnectToDB(s.dbConfig)
	if err != nil {
		log.Errorf("Failed to connect to database.")
		return
	}
	defer db.Close()
	log.Infof("Server connected to database.")

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
			c.SetClientID(s.gm.NextClientID())
			s.gm.AddClient(c)
			c.Start()

			go commands.HandleClient(c, s.gm)
		}
	}()
	// Wait for interrupt signal
	<-sigChan
	log.Info("Received interrupt signal. Saving all data...")
}
