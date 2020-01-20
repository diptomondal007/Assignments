package worker

// Job ...
type Job struct {
	ID   int
	term int
}

// JobChannel ...
type JobChannel chan Job

// JobQueue ...
type JobQueue chan chan Job

// Worker ...
type Worker struct {
	ID      int
	JobChan JobChannel
	Queue   JobQueue
	Quit    chan struct{}
}

// New ...
func New(ID int, JobChan JobChannel, Queue JobQueue, Quit chan struct{}) *Worker {
	return &Worker{
		ID:      ID,
		JobChan: JobChan,
		Queue:   Queue,
		Quit:    Quit,
	}
}

// Start ...
func (wr *Worker) Start() {
	//c := &http.Client{}
	go func() {
		for {

			wr.Queue <- wr.JobChan
			select {
			case _ = <-wr.JobChan:
				//asdasd
			case <-wr.Quit:
				close(wr.JobChan)
				return
			}
		}
	}()
}

// Stop ...
func (wr *Worker) Stop() {
	close(wr.Quit)
}
