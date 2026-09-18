package player

import (
	"fmt"
	"multispore/internal/client"
	"multispore/internal/commands"
	"multispore/internal/managers"
	"multispore/internal/models/creature"
	"multispore/internal/types/request"
	"multispore/internal/types/response"
	"strings"
)

func init() {
	commands.RegisterCommand(request.C2S_REGISTER,
		commands.CommandConfig{
			Name:        "Register",
			Identifier:  response.S2C_REGISTER,
			Description: "Register",
			Args:        "{}",
			MinArgs:     4,
			MaxArgs:     4,
		},
		RegisterValidator,
		Register,
		nil,
	)
}

var invalidChars = "+%&*/()[]{}\"'\\´`^°§€²³,;:?µ$"

// username+pass+creature
func Register(req *request.Request, c *client.Client, gm *managers.GameManager, cm *commands.CommandConfig) error {
	c.SendExtensionResponse(cm.Identifier, "1", "statusCodeStr")
	return nil
}

func RegisterValidator(req *request.Request, c *client.Client, gm *managers.GameManager, cm *commands.CommandConfig) (error, commands.ErrorCodes) {
	if len(req.Args) < cm.MinArgs {
		return fmt.Errorf("Not enough args. NEEDED/GOT: %d/%d", cm.MinArgs, len(req.Args)), commands.MIN_ARGS
	}

	if cm.MinArgs > 0 {
		if len(req.Args) > cm.MaxArgs {
			return fmt.Errorf("Too much args. NEEDED/GOT: %d/%d", cm.MaxArgs, len(req.Args)), commands.MAX_ARGS
		}
	}

	username := req.Args[1]
	password := req.Args[2]
	creatureData := req.Args[3]

	if len(username) < 4 {
		return fmt.Errorf("Can't register the player, because the username is short!"), -1
	}

	if strings.ContainsAny(username, invalidChars) {
		return fmt.Errorf("Can't register the player, because the username is wrong / contains not allowed chars!"), -1
	}

	if len(password) < 4 {
		return fmt.Errorf("Can't register the player, because the password is short!"), -1
	}

	if strings.ContainsAny(password, invalidChars) {
		return fmt.Errorf("Can't register the player, because the password is containts invalid chars!"), -1
	}

	r, err := creature.NewRigblockFromString(creatureData)
	if err != nil {
		return err, -1
	}

	creature.String(r)

	_, err = c.DB.GetPlayerByName(username)
	if err == nil {
		return err, -1
	}

	return nil, commands.SUCCESS
}
