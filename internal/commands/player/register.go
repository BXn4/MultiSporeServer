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

	"github.com/charmbracelet/log"
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
	username := req.Args[1]
	password := req.Args[2]
	rigblocks := req.Args[3]

	log.Debug("Everything is fine! Player register should start")

	creature := creature.NewCreatureFromString(rigblocks)
	creature.Health = 6
	creature.MaxHealth = creature.Health
	creature.Hunger = 0

	log.Debugf("Player creature: %s", creature.Rigblocks.String())

	player, err := c.DB.CreateAccount(username, password, *creature)
	if err != nil {
		return err
	}
	c.Player = player
	room, err := gm.AddRoom(c.Player.GetStage())
	if err != nil {
		return fmt.Errorf("Failed to load room for player %d: %v", c.Player.GetID(), err)
	}

	c.Location = room

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
	rigblocks := req.Args[3]

	if len(username) < 4 {
		return fmt.Errorf("Can't register the player, because the username is short!"), commands.USERNAME_SHORT
	}

	if len(username) > 32 {
		return fmt.Errorf("Can't register the player, because the username is long!"), commands.USERNAME_LONG
	}

	if strings.ContainsAny(username, invalidChars) {
		return fmt.Errorf("Can't register the player, because the username is wrong / contains not allowed chars!"), commands.USERNAME_WRONG
	}

	if len(password) < 4 {
		return fmt.Errorf("Can't register the player, because the password is short!"), commands.PASSWORD_SHORT
	}

	if len(username) > 64 {
		return fmt.Errorf("Can't register the player, because the password is long!"), commands.PASSWORD_LONG
	}

	if strings.ContainsAny(password, invalidChars) {
		return fmt.Errorf("Can't register the player, because the password is containts invalid chars!"), commands.PASSWORD_WRONG
	}

	creature := creature.NewCreatureFromString(rigblocks)
	if creature == nil {
		return fmt.Errorf("Failed to create a new creature from string!"), commands.INVALID_CREATURE_DATA
	}

	_, err := c.DB.GetPlayerByName(username)
	if err == nil {
		return fmt.Errorf("Account exist!"), commands.ACCOUNT_EXIST
	}

	return nil, commands.SUCCESS
}
