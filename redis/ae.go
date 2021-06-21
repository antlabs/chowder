package redis

import (
	"errors"
	"time"

	"golang.org/x/sys/unix"
)

type FileProc func(eventLoop *EventLoop, fd int, clientData interface{}, mask int)
type TimeProc func(eventLoop *EventLoop, id int, clientData interface{})
type EventFinalizerProc func(eventLoop *EventLoop, clientData interface{})
type BeforeSleepProc func(eventLoop *EventLoop)

type FileEvent struct {
	mask       Action
	wfileProc  FileProc
	rfileProc  FileProc
	clientData interface{}
}

type TimeEvent struct {
	id            int
	when_sec      int
	when_ms       int
	timeProc      TimeProc
	finalizerProc EventFinalizerProc
	clientData    interface{}
	next          *TimeEvent
}

type FiredEvent struct {
	fd   int
	mask int
}

type EventLoop struct {
	maxFd           int // highest file descriptor currently registered
	setSize         int // max number of file descriptors tracked
	timeEventNextId int
	lastTime        int
	events          []FileEvent
	fired           []FiredEvent
	timeEventHead   *TimeEvent
	stop            bool
	apidata         string
	beforeSleep     BeforeSleepProc
	afterSleep      BeforeSleepProc
}

// 初始化函数
func Create(setSize int) {
	return &EventLoop{
		setSize: setSize,
		maxFd:   -1,
		events:  make([]FileEvent, setSize),
		fired:   make([]FiredEvent, setSize),
	}
}

func (el *EventLoop) Delete() {
	el.events = nil
	el.fired = nil

}

func (el *EventLoop) Stop() {
	el.stop = true

}

func (el *EventLoop) CreateFileEvent(fd int, mask int, proc FileProc, clientData interface{}) error {
	if fd >= el.setSize {
		return errors.New("create file event fail")
	}

	fe := &el.events[fd]

	if err := el.apiAddEvent(fd, mask); err != nil {
		return err
	}

	fe.mask |= mask

	if mask&READABLE > 0 {
		fe.rfileProc = proc
	}

	if mask&WRITABLE > 0 {
		fe.wfileProc = proc
	}

	fe.clientData = clientData

	if fd > el.maxFd {
		el.maxFd = fd
	}

	return nil
}

func (el *EventLoop) DeleteFileEvent(fd int, mask int) {
	if fd >= el.setSize {
		return
	}

	fe := el.events[fd]

	if fe.mask == NONE {
		return
	}

	el.apiDelEvent(fd, mask)
	fe.mask = fe.mask & (^mask)

	if fd == el.maxFd && fe.mask == NONE {

		j := 0
		for j = el.maxFd - 1; j >= 0; j-- {
			if el.events[j].mask != NONE {
				break
			}
		}

		el.maxFd = j
	}

	return
}

func (el *EventLoop) GetFileEvents() int {
	if fd >= el.setSize {
		return 0
	}

	fe := el.events[fd]
	return fe.mask
}

func GetTime() time.Time {
	return time.Now()
}

func (el *EventLoop) FileProc(fd int, clientData interface{}, mask int) {
}

func (el *EventLoop) TimeProc(id int, clientData interface{}) {
}

func (el *EventLoop) EventFinalizerProc(clientData interface{}) {
}

func (el *EventLoop) CreateTimeEvent(milliseconds int, proc TimeProc, clientData interface{}, finalizerProc EventFinalizerProc) {
}

func (el *EventLoop) DeleteTimeEvent(id int) {

}

func (el *EventLoop) ProcessEvents(flags EVENT) {
	processed, numevents := 0, 0

	if !(flags&TIME_EVENTS > 0) || !(flags&FILE_EVENTS > 0) {
		return
	}

	if el.maxFd != -1 || flags&TIME_EVENTS > 0 && flags&DONT_WAIT == 0 {
	}
}

func Wait(fd int, mask int, milliseconds int) {
	pfd := make([]unix.PollFd, 1)
	pfd[0].Fd = int32(fd)

	reMask := 0
	retVal := 0
	if mask&READABLE > 0 {
		reMask |= unix.POLLIN
	}

	if mask&WRITABLE > 0 {
		reMask |= unix.POLLOUT
	}

	retVal, _ := unix.Poll(pfd, milliseconds)
	if retVal == 1 {
		if pfd[0].Revents&unix.POLLIN > 0 {
			reMask |= READABLE
		}
		if pfd[0].Revents&unix.POLLOUT > 0 {
			reMask |= WRITABLE
		}
		if pfd[0].Revents&unix.POLLERR > 0 {
			reMask |= WRITABLE
		}
		if pfd[0].Revents&unix.POLLHUP > 0 {
			reMask |= WRITABLE
		}

		return reMask
	}

	return retVal
}

func (el *EventLoop) Main() {
	for !el.Stop {
		if el.beforeSleep != nil {
			el.beforeSleep(el)
		}
		el.ProcessEvents(ALL_EVENTS)
	}
}

func (el *EventLoop) GetApiName() string {
	return apiName()
}

func (el *EventLoop) SetBeforeSleepProc(beforeSleep BeforeSleepProc) {
	el.beforeSleep = beforesleep
}

func (el *EventLoop) SetBeforeSleepProc(beforeSleep BeforeSleepProc) {
	el.beforeSleep = beforesleep
}

func (el *EventLoop) GetSize() int {
	return el.setSize
}

func (el *EventLoop) Resize(setSize int) error {
	if setSize == el.setSize {
		return nil
	}

	if el.maxFd >= setSize {
		return errors.New("")
	}

	el.events = append(FileEvent{}, el.events[:setSize])
	el.fired = append(FiredEvent{}, el.fired[:setSize])

	return nil
}
