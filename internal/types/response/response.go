package response

import (
	"strings"
)

const (
	S2C_CONNECT       = "cnc"
	S2C_DISCONNECT    = "dnc"
	S2C_CREATURE_DATA = "cda"
	S2C_LOGIN         = "lgn"
	S2C_REGISTER      = "rgr"
)

type Response interface {
	Args() []string
	Wrap() string
}

type ExtensionResponse struct {
	args []string
}

func NewExtensionResponse(args ...string) *ExtensionResponse {
	return &ExtensionResponse{args: args}
}

func (r *ExtensionResponse) Wrap() string {
	return "%xt%" + strings.Join(r.args, "%") + "%\x00"
}

func (r *ExtensionResponse) Args() []string {
	return r.args
}
