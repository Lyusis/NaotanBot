package engine

const (
	DelayOp = "DELAY OPERATION"
)

type ConcurrentEngine struct {
	Scheduler    Scheduler
	WorkerCount  int
	SaveChan     chan interface{}
	RequestChan  chan Request
	Workers      func(Request, chan ResultItems)
	RoutineCount int64
}

type Scheduler interface {
	WorkerReady(chan Request)
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type Request struct {
	Url    string
	Name   string
	Method string
	Body   []byte
	Parser func([]byte) ResultItems
}

type ResultItems struct {
	Items []interface{}
}

func NilResult() ResultItems {
	return ResultItems{}
}
