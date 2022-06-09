package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	id      string
	message Notification
}

type Notification struct {
	title       string
	description string
	sentTime    time.Time
}

type Ack struct {
	status int64
}

const (
	Unsuccessful = 0
	Successful   = 1
)

func worker(id int, jobs <-chan Event, results chan<- Ack) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- Ack{
			status: Successful,
		}
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan Event, numJobs)
	results := make(chan Ack, numJobs)
	event := Event{
		id: uuid.New().String(),
		message: Notification{
			title:       "Worker learning",
			description: "I'm learning!",
			sentTime:    time.Now(),
		},
	}

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- event
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
