package creture

import (
	"multispore/internal/client"
	"multispore/internal/commands"
	"multispore/internal/managers"
	"multispore/internal/types/request"
	"multispore/internal/types/response"
)

func init() {
	commands.RegisterCommand(request.C2S_CREATURE_DATA,
		commands.CommandConfig{
			Name:        "Creature data",
			Description: "Sends creature data to client",
			Identifier:  response.S2C_CREATURE_DATA,
			MinArgs:     0,
			MaxArgs:     0,
		},
		nil,
		CreatureData,
		nil,
	)
}

func CreatureData(req *request.Request, c *client.Client, gm *managers.GameManager, cm *commands.CommandConfig) error {
	return nil
}
