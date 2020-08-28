package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	workQueue := new(WorkQueue)
	workQueue.Jobs = make(chan Worker, maxJobs)
	workQueue.Results = make(chan interface{}, maxJobs)

	for i := 0; i < int(nWorkers); i++{
		go workQueue.worker()
	}

	return workQueue
}
