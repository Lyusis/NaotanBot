package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan interface{}
	RequestProcessor Processor
}

type Processor func(Request) (Result, error)

type Scheduler interface {
	WorkerReady(chan Request)
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type Request struct {
	Url    string
	Name   string
	Parser func([]byte) Result
}

type Result struct {
	Items []interface{}
}

type ParserFunc func(
	contents []byte, url string) Result

type FuncParser struct {
	parser ParserFunc
	name   string
}

func NilResult() Result {
	return Result{}
}
