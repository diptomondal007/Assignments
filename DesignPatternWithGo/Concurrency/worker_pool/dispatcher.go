package worker_pool

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(request Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(id int,w WorkerLauncher){
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop(){
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request){
	d.inCh <- r
}