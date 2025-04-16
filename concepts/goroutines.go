package concepts

import (
	"fmt"
	"time"
)

func GoRoutines() {
	go Say("world")
	Say("hello")
	go LogActivity("Customer Request")
	go SendNotification("Customer Request")
	LogActivity("Order Request")
	SendNotification("Order Request")
	time.Sleep(1000 * time.Millisecond)
}
func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func LogActivity(user string) {
	fmt.Printf("[%s] Logging user activity...\n", user)
	time.Sleep(1 * time.Second) // simulating DB wri te delay
	fmt.Printf("[%s] Activity logged successfully.\n", user)
}

func SendNotification(user string) {
	fmt.Printf("[%s] Sending notification...\n", user)
	time.Sleep(2 * time.Second) // simulating notification delay
	fmt.Printf("[%s] Notification sent.\n", user)
}

func WorkerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
