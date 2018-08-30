package utils

const (
	STRING_TYPE  byte = 0
	SET_TYPE     byte = 1
	ZSET_TYPE    byte = 2
	ZSET_SCORE   byte = 3
	TTL_TYPE     byte = 109
	EXPTIME_TYPE byte = 110
)
const (
	FLAG_NORMAL byte = iota
	FLAG_DELETED
)

var (
	EmptyListInterfaces []interface{} = make([]interface{}, 0)
)
