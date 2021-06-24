package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	SaveChan         chan interface{}
	RequestProcessor Processor
}

type Processor func(Request) (ResultItems, error)

type Scheduler interface {
	WorkerReady(chan Request)
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type Request struct {
	Url           string
	Name          string
	PrimaryParser func([]byte) ResultItems
	PostParser    func(result ResultItems) SaveItems
}

type ResultItems struct {
	Items []interface{}
}

type SaveItems struct {
	Items []interface{}
}

type ParserFunc func(
	contents []byte, url string) ResultItems

type FuncParser struct {
	parser ParserFunc
	name   string
}

func NilResult() ResultItems {
	return ResultItems{}
}
