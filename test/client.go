package main

import (
	"fmt"
	"io"
	"multispore/internal/utils"
	"net"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

const END_CHAR = "\x00"

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

	send(conn, "%xt%spore%lgn%1%2")

	TEST_CELL := "0+-1+-1+11+5+0.4000+-0.0000+-0.2851+0.0000+0.0000+0.0000+1.0000+0+1080123392+3244096132#1+0+-1+9+5+0.6000+-0.0000+-0.1734+0.0000+0.2500+0.0000+1.0000+0+1080123392+3244096132#2+1+-1+9+5+1.0750+-0.0000+-0.0618+0.0000+0.8438+0.0000+1.0000+0+1080123392+3244096132#3+2+-1+9+5+0.9750+-0.0000+0.0499+0.0000+0.7188+0.0000+1.0000+0+1080123392+3244096132#4+3+-1+9+5+0.6250+-0.0000+0.1616+0.0000+0.2812+0.0000+1.0000+0+1080123392+3244096132#5+4+-1+9+5+0.5625+-0.0000+0.2721+-0.0112+0.2031+0.0000+1.0000+0+1080123392+3244096132#6+0+-1+1088+2+0.9277+-0.0000+-0.3648+-0.0963+0.4277+0.0000+1.0000+1014894446+1080123392+2714007141#7+2+-1+512+2+0.8607+-0.0000+-0.0390+0.0634+0.3607+0.0000+1.0000+0+1080123392+1729351260#8+3+-1+512+2+1.0779+-0.0000+0.1459+0.0480+0.5779+0.0000+1.0000+0+1080123392+1729351260#9+3+10+2304+2+1.0779+0.2018+0.0860+-0.1000+0.5779+0.0000+1.0000+0+1080123392+898723801#10+3+9+2048+2+1.0779+-0.2019+0.0860+-0.1000+0.5779+0.0000+1.0000+0+1080123392+2053914431"

	send(conn, "%xt%spore%rgr%user%pass%"+TEST_CELL)

	conn.Close()
}

func send(conn net.Conn, data string) {
	conn.Write([]byte(data + END_CHAR))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	log.Info(string(buf[:n]))
}
