package redis

type TimeEvent struct {
	id            int
	when          time.Durtion
	timeProc      TimeProc
	finalizerProc EventFinalizerProc
	clientData    interface{}
}
