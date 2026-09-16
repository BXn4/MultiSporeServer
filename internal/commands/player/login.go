package server

import (
	"fmt"
	"multispore/internal/client"
	"multispore/internal/commands"
	"multispore/internal/managers"
	"multispore/internal/types/request"
	"multispore/internal/types/response"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

func init() {
	commands.RegisterCommand(request.C2S_LOGIN,
		commands.CommandConfig{
			Name:        "Login",
			Description: "Handles login",
			Identifier:  response.S2C_LOGIN,
			MinArgs:     2,
			MaxArgs:     2,
		},
		nil,
		Login,
		nil,
	)
}

func Login(req *request.Request, c *client.Client, gm *managers.GameManager, cm *commands.CommandConfig) error {
	name := req.Args[2]
	password := req.Args[3]

	_, statusCode, err := c.DB.Authenticate(name, password)

	if err == nil {
		log.Infof("Checking if user %s is already logged in", name)
		searched, searchErr := gm.GetClientByName(name)
		if searched != nil && searchErr == nil {
			log.Infof("User %s found as already logged in on client %d, kicking existing session", name, searched.ClientID)
			statusCode = 3

			if c.GetIP() == searched.GetIP() {
				kickErr := searched.Disconnect()
				if kickErr != nil {
					log.Errorf("Failed to kick existing session for user %s: %v", name, kickErr)
				} else {
					log.Infof("Successfully kicked existing session for user %s", name)
				}
				time.Sleep(100 * time.Millisecond)
			}
		} else {
			log.Infof("User %s is not currently logged in (searchErr: %v)", name, searchErr)
		}
	}

	statusCodeStr := strconv.Itoa(statusCode)

	c.SendExtensionResponse(cm.Identifier, "1", statusCodeStr)
	if statusCode != 0 {
		if statusCode == 3 {
			return fmt.Errorf("Player %v is already logged in", name)
		} else {
			return fmt.Errorf("Access denied")
		}
	}

	return nil
}
