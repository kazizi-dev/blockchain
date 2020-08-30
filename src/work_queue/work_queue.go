package work_queue

// jobs should be started calling .Run()
type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

//var wg sync.WaitGroup

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	queue := new(WorkQueue)

	// make channels for jobs going into the queue and results coming out
	queue.Jobs = make(chan Worker, maxJobs)
	queue.Results = make(chan interface{}, maxJobs)

	// create worker goroutines to watch the Jobs queue for work
	for i := 0; i < int(nWorkers); i++{
		go queue.worker()
	}

	return queue
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	for job := range queue.Jobs {
		// runs jobs and return values back to the channel
		queue.Results <- job.Run()
	}
//	close(queue.Results)
}

// Put the work into the Jobs channel so a worker can find it and start the task.
func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
}

// Close .Jobs and remove all remaining jobs from the channel.
func (queue WorkQueue) Shutdown() {
	close(queue.Jobs)
}
