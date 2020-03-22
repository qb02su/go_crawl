package engine

type ReadyNotifier interface {
	WorkerReady(chan Request)
}
type Scheduler interface {
	Submit(request Request)
	WorkerChan()chan Request
	//ConfigureWorkerChan(chan Request)
	ReadyNotifier
	Run()
}



type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

func (e *ConcurrentEngine)Run(seeds ...Request){
	//in:=make(chan Request)
	out:=make(chan ParseResult)
	//e.Scheduler.ConfigureWorkerChan(in)
	e.Scheduler.Run()
	for i:=0;i<e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}
	for _,r:=range(seeds){
		e.Scheduler.Submit(r)
	}

	//itemCount:=0
	for{
		result:=<-out
		for _,item:=range result.Item{
			//log.Printf("Got %d Item %v ",itemCount,item)
			//itemCount++
			go func() {
				e.ItemChan<-item
			}()
		}
		for _,request:=range result.Request{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifier){
	go func(){
		for {
			//s.WorkerReady(in)
			ready.WorkerReady(in)
			request:=<-in
			result,err:=Worker(request)
			if err!=nil{
				continue
			}
			out<-result
		}
	}()
}