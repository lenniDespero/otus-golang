package hw8

//WorkerFunction keep functions to workers
type WorkerFunction func() error

//Run slice of functions on N workers while get M errors
func Run(functions []WorkerFunction, workersCount int, maxErrors int) ([]error, int) {
	tasksChan := make(chan WorkerFunction)
	resultsChan := make(chan error, workersCount-1)
	closeChan := make(chan bool)
	closedChan := make(chan bool)
	var counter int
	var errorsSlice []error

	for i := 0; i < workersCount; i++ {
		go startWorker(tasksChan, resultsChan, closeChan, closedChan)
	}

	func(counter *int, errorsSlice *[]error) {
		var errors int
		var inProgress int

		for i := 0; i < workersCount; i++ {
			inProgress++
			tasksChan <- functions[i]
		}

		for {
			err := <-resultsChan
			inProgress--
			*counter++

			if err != nil {
				*errorsSlice = append(*errorsSlice, err)
				errors++
			}

			if *counter == len(functions) || errors == maxErrors {
				close(closeChan)
				return
			} else if len(functions)-*counter-inProgress > 0 {
				inProgress++
				tasksChan <- functions[workersCount-1+*counter]
			}
		}
	}(&counter, &errorsSlice)

	for i := 0; i < workersCount; i++ {
		<-closedChan
	}
	return errorsSlice, counter
}

func startWorker(tasksChan <-chan WorkerFunction, resultsChan chan<- error, closeChan <-chan bool, closedChan chan<- bool) {
	defer func() {
		closedChan <- true
	}()

	for {
		select {
		case task := <-tasksChan:
			resultsChan <- task()
		case <-closeChan:
			return
		}
	}
}
