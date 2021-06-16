package redis

import (
	"golang.org/x/sys/unix"
)

type apiState struct {
	epfd   int
	events []unix.EpollEvent
}
