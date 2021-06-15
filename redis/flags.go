package redis

type Option int8

const (
	OK  Option = 0
	ERR        = -1
)

type Action int8

const (
	NONE     Action = 0
	READABLE        = 1
	WRITABLE        = 2
)

type Event int8

const (
	FILE_EVENTS Event = 1
	TIME_EVENTS       = 2
	ALL_EVENTS        = (FILE_EVENTS | TIME_EVENTS)
	DONT_WAIT         = 4
)

const NOMORE = -1
const DELETED_EVENT_ID = -1
