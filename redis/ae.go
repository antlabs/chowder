package redis

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
	beforesleep     BeforeSleepProc
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
}

func (el *EventLoop) DeleteFileEvent(fd int, mask int) {
}

func (el *EventLoop) GetFileEvents() {
}

func (el *EventLoop) FileProc(fd int, clientData interface{}, mask int) {
}

func (el *EventLoop) TimeProc(id int, clientData interface{}) {
}

func (el *EventLoop) EventFinalizerProc(clientData interface{}) {
}

func (el *EventLoop) BeforeSleepProc() {
}

func (el *EventLoop) CreateTimeEvent(milliseconds int, proc TimeProc, clientData interface{}, finalizerProc EventFinalizerProc) {
}

func (el *EventLoop) DeleteTimeEvent(id int) {

}

func (el *EventLoop) ProcessEvents(flags int) {
}

func Wait(fd int, mask int, milliseconds int) {
}

func (el *EventLoop) Main() {
}

func (el *EventLoop) GetApiName() string {
}

func (el *EventLoop) SetBeforeSleepProc() {
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
