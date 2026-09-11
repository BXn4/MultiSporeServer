package client

import (
	"bufio"
	"io"
	"multispore/types"
	"net"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type Client struct {
	ClientID int
	Conn     net.Conn
	Writer   *bufio.Writer
	Reader   *bufio.Reader

	TimeoutStamp    time.Time
	isDisconnecting bool

	RequestQueue  chan *types.Request
	ResponseQueue chan types.Response
}

func New(conn net.Conn) *Client {
	return &Client{
		ClientID: 0,
		Conn:     conn,
		Reader:   bufio.NewReader(conn),
		Writer:   bufio.NewWriter(conn),

		TimeoutStamp: time.Now(),

		RequestQueue:  make(chan *types.Request, 255),
		ResponseQueue: make(chan types.Response, 255),
	}
}

func (c *Client) ID() int {
	return c.ClientID
}

func (c *Client) SetClientID(id int) {
	c.ClientID = id
}

func (c *Client) GetIP() string {
	return strings.Split(c.Conn.RemoteAddr().String(), ":")[0]
}

func (c *Client) SendExtensionResponse(args ...string) {
	resp := types.NewExtensionResponse(args...)
	log.Logf(log.Level(-3), "%s", resp.Wrap())
	c.ResponseQueue <- resp
}

func (c *Client) receiveRequests() {
	defer close(c.RequestQueue)
	for {
		message, err := c.Reader.ReadString('\x00')
		if err != nil {
			if err == io.EOF {
				log.Infof("Client disconnected (EOF): %s", c.GetIP())
			} else if netErr, ok := err.(net.Error); ok {
				log.Infof("Network error from %s: %v", c.GetIP(), netErr)
			} else if strings.Contains(err.Error(), "use of closed network connection") {
				log.Infof("Connection closed: %s", c.GetIP())
			} else {
				log.Errorf("Read error from %s: %v", c.GetIP(), err)
			}
			c.Disconnect()
			return
		}
		log.Logf(log.Level(-5), "%s", message)

		req, err := types.ParseRequest(strings.Trim(message, "\x00"))
		if err != nil {
			log.Error("Failed to parse request: %v", err)
			continue
		}
		c.RequestQueue <- req
	}
}

func (c *Client) sendResponses() {
	defer close(c.ResponseQueue)
	for resp := range c.ResponseQueue {
		if resp == nil {
			return
		}
		c.Writer.Write([]byte(resp.Wrap()))
		c.Writer.Flush()
	}
}

func (c *Client) Disconnect() error {
	c.isDisconnecting = true

	log.Infof("Client being disconnected: %s", c.GetIP())

	c.Conn.Close()

	return nil
}

func (c *Client) IsDisconnecting() bool {
	return c.isDisconnecting
}
