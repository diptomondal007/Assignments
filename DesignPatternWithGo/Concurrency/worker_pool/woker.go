package worker_pool

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}
