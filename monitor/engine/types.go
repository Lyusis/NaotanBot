package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	SaveChan    chan interface{}
	Workers     func(Request) (ResultItems, error)
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
	Parser func([]byte) ResultItems
}

type ResultItems struct {
	Items []interface{}
}

func NilResult() ResultItems {
	return ResultItems{}
}
