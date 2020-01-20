package dispatcher

import "jobqueue/worker"

// disp ...
type disp struct {
	Workers []*worker.Worker
	JobChan worker.JobChannel
	Queue   worker.JobQueue
}

// New ...
func New(num int) *disp {
	return &disp{
		Workers: make([]*worker.Worker, num),
		JobChan: make(worker.JobChannel),
		Queue:   make(worker.JobQueue),
	}
}

func (d *disp) Start() *disp {
	l := len(d.Workers)
	for i := 1; i <= l; i++ {
		wrk := worker.New(i, make(worker.JobChannel), d.Queue, make(chan struct{}))
		wrk.Start()
		d.Workers = append(d.Workers, wrk)
	}
	return d
}

func (d *disp) Process() {
	for {
		select {
		case job := <-d.JobChan:
			jobChan := <-d.Queue
			jobChan <- job
		}
	}
}

func (d *disp) Submit(job worker.Job) {
	d.JobChan <- job
}
