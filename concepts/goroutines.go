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

// Simulate the print job handling using a worker pool
func PrinterPool() {
	// Total number of print jobs to be processed
	const numJobs = 5

	// Channel to queue incoming print jobs
	printJobs := make(chan int, numJobs)

	// Channel to receive print confirmations (results)
	printConfirmations := make(chan int, numJobs)

	// Start 3 printers (workers)
	for printerID := 1; printerID <= 3; printerID++ {
		go Printer(printerID, printJobs, printConfirmations)
	}

	// Send 5 print jobs to the queue
	for jobID := 1; jobID <= numJobs; jobID++ {
		fmt.Printf("📝 Submitted Print Job #%d to the queue\n", jobID)
		printJobs <- jobID
	}

	// Close the job queue after all jobs are submitted
	close(printJobs)

	// Wait for all print confirmations
	for i := 1; i <= numJobs; i++ {
		<-printConfirmations
	}
	fmt.Println("\n✅ All print jobs completed.")
}

// printer simulates a worker (printer) that processes print jobs
func Printer(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		// Start printing
		fmt.Printf("🖨️ Printer %d started printing Job #%d\n", id, job)
		time.Sleep(1 * time.Second) // Simulate print time

		// Print complete
		fmt.Printf("✅ Printer %d finished Job #%d\n", id, job)

		// Send confirmation
		results <- job
	}
}
