package main

import (
	"fmt"
	"io"
	"multispore/internal/utils"
	"net"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	envFile, err := godotenv.Read("../.env")

	hasConfig := err == nil

	if !hasConfig {
		log.Warnf("Cannot find .env file, using default!")
	}

	address := utils.If(hasConfig, envFile["SERVER_HOST"], "127.0.0.1")
	port := utils.If(hasConfig, envFile["SERVER_PORT"], "5523")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Info("Connected to the server!")

	/*_, err = fmt.Fprintf(conn, "")
	if err != nil {
		log.Fatal(err)
		}*/

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Print(string(buf[:n]))
}
