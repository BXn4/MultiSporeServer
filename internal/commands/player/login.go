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
			MinArgs:     3,
			MaxArgs:     3,
		},
		LoginValidator,
		Login,
		nil,
	)
}

func Login(req *request.Request, c *client.Client, gm *managers.GameManager, cm *commands.CommandConfig) error {
	name := req.Args[1]
	password := req.Args[2]

	p, statusCode, err := c.DB.Authenticate(name, password)

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

	if p != nil {
		c.Player = p
		id := c.Player.ID

		room, err := gm.AddRoom(p.Stage) // return the room if its found. if not, creates new

		if err != nil {
			return fmt.Errorf("Failed to load location %d: %v", id, err)
		}
		c.Location = room
	}

	return nil
}

func LoginValidator(req *request.Request, c *client.Client, gm *managers.GameManager, cm *commands.CommandConfig) (string, commands.ErrorCodes) {
	println(req.Args[1])
	println(req.Args[2])
	if len(req.Args) < cm.MinArgs {
		return fmt.Sprintf("Not enough args. NEEDED/GOT: %d/%d", cm.MinArgs, len(req.Args)), commands.MIN_ARGS
	}

	if cm.MinArgs > 0 {
		if len(req.Args) > cm.MaxArgs {
			return fmt.Sprintf("Too much args. NEEDED/GOT: %d/%d", cm.MaxArgs, len(req.Args)), commands.MAX_ARGS
		}
	}

	return "Command ran without any errors.", commands.SUCCESS
}
