package qkverror

import (
	"errors"
)

var (
	ErrorServerNoAuthNeed = errors.New("client sent auth, but server no password")
	ErrorAuthFailed       = errors.New("client sent a invalid password ")
	ErrorNoAuth           = errors.New("client no authentication")
	ErrorCommand          = errors.New("command invalid")
	ErrorCommandParams    = errors.New("command params invalid")
	ErrorUnknownType      = errors.New("unknown response data type")
	ErrorKeyEmpty         = errors.New("key can't be empty")
	ErrorServerInternal   = errors.New("server internal error")
	ErrorTypeNotMatch     = errors.New("key type not match")
	ErrorInvalidMeta      = errors.New("invalid key meta")
	ErrorInvalidRawData   = errors.New("invalid raw data")
	ErrorWrongType        = errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
	ErrorNotInteger       = errors.New("value is not an integer or out of range")
	ErrorOutOfRange       = errors.New("index out of range")
)
