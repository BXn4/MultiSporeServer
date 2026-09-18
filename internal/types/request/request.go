package request

import (
	"errors"
	"strings"
)

type RequestKind int

const (
	UNKNOWN RequestKind = iota
	C2S_CONNECT
	C2S_DISCONNECT
	C2S_CREATURE_DATA
	C2S_LOGIN
	C2S_REGISTER
)

func LookupRequestKind(kindStr string) RequestKind {
	kindLookup := map[string]RequestKind{
		"cnc": C2S_CONNECT,
		"dnc": C2S_DISCONNECT,
		"cda": C2S_CREATURE_DATA,
		"lgn": C2S_LOGIN,
		"rgr": C2S_REGISTER,
	}

	if val, ok := kindLookup[kindStr]; ok {
		return val
	}

	return UNKNOWN
}

type Request struct {
	Kind RequestKind
	Args []string
}

func ParseRequest(raw_request string) (*Request, error) {
	trimmed_req := strings.TrimSpace(raw_request)

	if strings.HasPrefix(trimmed_req, "%xt") {
		return ParseExtensionRequest(trimmed_req)
	}

	return nil, errors.New("Can not parse request!")
}

func ParseExtensionRequest(req string) (*Request, error) {

	trimmed_req := strings.Trim(req, "%")
	args := strings.Split(trimmed_req, "%")
	args = args[2:]

	kind := LookupRequestKind(args[0])

	return &Request{
		Kind: kind,
		Args: args,
	}, nil
}
