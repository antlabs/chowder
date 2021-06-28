package redis

import "errors"

var (
	OK  = errors.New("OK")
	ERR = errors.New("ERR")
)

type Action int8

const (
	NONE     Action = 0
	READABLE        = 1
	WRITABLE        = 2
)

type Event int8

const (
	FILE_EVENTS Event = 1 << iota
	TIME_EVENTS
	DONT_WAIT
	CALL_BEFORE_SLEEP
	CALL_AFTER_SLEEP
	ALL_EVENTS = (FILE_EVENTS | TIME_EVENTS)
)

const NOMORE = -1
const DELETED_EVENT_ID = -1
