package main

import (
	"fmt"
	"golangconcepts/concepts"
	"golangconcepts/models"
	"sync"
	"time"
)

func main() {

	// Embedding //
	// concepts.Embedding()

	////Goroutines//
	// go concepts.Say("world")
	// concepts.Say("hello")
	// go concepts.LogActivity("Customer Request")
	// go concepts.SendNotification("Customer Request")
	// concepts.LogActivity("Order Request")
	// concepts.SendNotification("Order Request")
	//time.Sleep(1000 * time.Millisecond)

	//// channels: Example 1//
	user := "Sandeep"
	ch := make(chan string)

	// Start the receiver first
	go concepts.ReceiveConfirmation(ch)

	// Send message in a separate goroutine
	go concepts.SendMessage(user, ch)

	// Log while waiting
	fmt.Println("[Main] Waiting for goroutines to finish...")
	time.Sleep(3 * time.Second)
	fmt.Println("[Main] Messaging workflow completed.")

	//// channels with waitgroup: Example 2//

	fmt.Println("[Main] All messages sent and confirmed.")
	users := []string{"Sandeep", "Anjali", "Rahul", "Maya", "Zara"}
	var wg sync.WaitGroup

	for _, user := range users {
		ch := make(chan string)

		wg.Add(2) // one for sender, one for receiver

		go concepts.SendMessageWithWaitGroup(user, ch, &wg)
		go concepts.ReceiveConfirmationWithWaitGroup(user, ch, &wg)

	}

	fmt.Println("[Main] Waiting for all message operations to complete...")
	wg.Wait()
	fmt.Println("[Main] All messages sent and confirmed.")
	fmt.Println("[Main] Total messages sent: %s", models.TotalMessages) //// Example with Mutex//

	//// Example with Mutex//
	var counter int
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Final count:", counter)

}
