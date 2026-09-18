package commands

import (
	"fmt"
	"multispore/internal/client"
	"multispore/internal/managers"
	"multispore/internal/types/request"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Command struct {
	Config    CommandConfig
	Validator func(req *request.Request, c *client.Client, gm *managers.GameManager, cm *CommandConfig) (error, ErrorCodes)
	Handler   func(req *request.Request, c *client.Client, gm *managers.GameManager, cm *CommandConfig) error
	DBSaver   func(c *client.Client) error
}

type CommandConfig struct {
	Name        string // Command name to us to easier to known which command is that (CreatureData)
	Identifier  string // 3 letter identifier to the command (CDA) CreatureData
	Description string // Whats the command will do.
	Args        string // What args needed to be send back.

	MinArgs int  // if the args less than X (needed) len, then dont allow command to run. Probaly will missing some values.
	MaxArgs int  // if the args more than X (needed) len, then dont allow command to run.
	IsBool  bool // 1 or 0

	Category string // Categories.
}

var Commands = map[request.RequestKind]Command{}

func RegisterCommand(
	kind request.RequestKind,
	commandConfig CommandConfig,
	commandValidator func(*request.Request, *client.Client, *managers.GameManager, *CommandConfig) (error, ErrorCodes),
	commandHandler func(*request.Request, *client.Client, *managers.GameManager, *CommandConfig) error,
	commandDBSaver func(*client.Client) error,
) {
	Commands[kind] = Command{
		Config:    commandConfig,
		Validator: commandValidator,
		Handler:   commandHandler,
		DBSaver:   commandDBSaver,
	}
}

func ErrorHandler(req *request.Request, c *client.Client, cm *CommandConfig, err error, errc ErrorCodes) error {
	c.SendExtensionResponse(cm.Identifier, "-1", strconv.Itoa(int(errc)), strings.Join(req.Args[2:], "%"))
	return fmt.Errorf("Command %s failed with error code: %d.\n---> Reason: %s", cm.Name, errc, err.Error())
}

func HandleClient(c *client.Client, gm *managers.GameManager) {
	for req := range c.RequestQueue {
		if c.IsDisconnecting() {
			log.Debug("Client is disconnecting, stopping request handler")
			return
		}
		if req == nil {
			return
		}

		err := HandleRequest(req, c, gm)
		if err != nil {
			log.Errorf("Error during request handling: %v", err.Error())
			continue
		}
	}

}

func HandleRequest(req *request.Request, c *client.Client, gm *managers.GameManager) error {
	command, implemented := Commands[req.Kind]

	if !implemented {
		cm := CommandConfig{
			Name:       req.Args[0],
			Identifier: req.Args[0],
		}
		return ErrorHandler(req, c, &cm, fmt.Errorf("The command is not implemented"), NOT_IMPLEMENTED)
	}

	// log.Debugf("Handling command: %s", command.Config.Name)

	// command error = int
	// all error codes what the cafe having in int
	if command.Validator != nil {
		err, commandError := command.Validator(req, c, gm, &command.Config)

		// If theres an ANY error, then dont run the handler.
		if commandError != SUCCESS {
			return ErrorHandler(req, c, &command.Config, err, commandError)
		}
	}

	err := command.Handler(req, c, gm, &command.Config)
	if err != nil {
		return fmt.Errorf("Error during command handling: %w", err)
	}

	if command.DBSaver != nil {
		err = command.DBSaver(c)
		if err != nil {
			return fmt.Errorf("Error during command DB saving: %w", err)
		}
	}

	return nil
}
