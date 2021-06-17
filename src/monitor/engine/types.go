package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan interface{}
	RequestProcessor Processor
}

type Processor func(Request) (Result, error)

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (engine *ConcurrentEngine) createWorker(
	in chan Request, out chan Result, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := engine.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

type Scheduler interface {
	ReadyNotifier
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
