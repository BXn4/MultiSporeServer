package commands

type ErrorCodes int

const (
	SUCCESS = 0
)

const (
	NOT_IMPLEMENTED = 1000 + iota
	MIN_ARGS
	MAX_ARGS
	INVALID_ARGS

	CONVERT_ERROR
	INVALID_VALUE
	NOT_DECLARED
)
