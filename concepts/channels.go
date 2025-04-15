package concepts

import (
	"fmt"
	"golangconcepts/models"
	"math/rand"
	"sync"
	"time"
)

func Channels() {
	//Example 1 for Channels
	user := "Sandeep"
	ch := make(chan string)

	// Start the receiver first
	go ReceiveConfirmation(ch)

	// Send message in a separate goroutine
	go SendMessage(user, ch)

	// Log while waiting
	fmt.Println("[Main] Waiting for goroutines to finish...")
	time.Sleep(3 * time.Second)
	fmt.Println("[Main] Messaging workflow completed.")

}

func ChannelsWithWaitGroup() {
	//// channels with waitgroup: Example 2//

	fmt.Println("[Main] All messages sent and confirmed.")
	users := []string{"Sandeep", "Anjali", "Rahul", "Maya", "Zara"}
	var wg sync.WaitGroup

	for _, user := range users {
		ch := make(chan string)

		wg.Add(2) // one for sender, one for receiver

		go SendMessageWithWaitGroup(user, ch, &wg)
		go ReceiveConfirmationWithWaitGroup(user, ch, &wg)

	}

	fmt.Println("[Main] Waiting for all message operations to complete...")
	wg.Wait()
	fmt.Println("[Main] All messages sent and confirmed.")
	fmt.Println("[Main] Total messages sent: %s", models.TotalMessages)

}

func ChannelsWithMutexExampleClearUnderstanding() {
	//// Example with Mutex//
	var counter int
	// var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			models.Mu.Lock()
			counter++
			models.Mu.Unlock()
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Final count:", counter)
}

func ChannelsWithRWMutex() {
	//// channels with RWMutex: Example 2//

	users := []string{"Sandeep", "Anjali", "Rahul", "Maya", "Zara"}
	var wg sync.WaitGroup

	for _, user := range users {
		ch := make(chan string)
		wg.Add(2)

		go SendMessageWithRwMutex(user, ch, &wg)
		go ReceiveConfirmationWithRwMutex(user, ch, &wg)
	}

	wg.Wait()
	fmt.Println("[Main] All messages processed.")

}

func ChannelsWithSelectStatement() {
	// Create buffered channels
	normalCh := make(chan string, 5)
	highCh := make(chan string, 5)

	// Launch sender
	go SendMessagesWithSelectStatement(normalCh, highCh)

	// Handle messages
	HandleMessagesWithSelectStatement(normalCh, highCh)

	fmt.Println("\n✅ All messages processed.")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func Sender(ch chan string) {
	ch <- "sending message"

}

func Receiver(ch chan string) {

	message := <-ch
	fmt.Println(message)
}

// Simulate sending a message to a user
func SendMessage(to string, ch chan string) {
	fmt.Println("[Sender] Preparing to send message...")
	time.Sleep(2 * time.Second) // simulate network delay
	ch <- fmt.Sprintf("Message sent to %s", to)
	fmt.Println("[Sender] Message has been pushed to channel.")
}

// Simulate receiving confirmation from the channel
func ReceiveConfirmation(ch chan string) {
	fmt.Println("[Receiver] Waiting for message...")
	msg := <-ch // blocking wait for message
	fmt.Println("[Receiver] Received:", msg)
}

func SendMessageWithWaitGroup(user string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[Sender] Sending message to %s...\n", user)
	time.Sleep(1 * time.Second) // simulate network delay
	ch <- fmt.Sprintf("Message sent to %s", user)
	// mu.Lock()
	models.TotalMessages++
	// mu.Unlock()
}

func ReceiveConfirmationWithWaitGroup(user string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	fmt.Printf("[Receiver] Confirmation for %s: %s\n", user, msg)
}

func SendMessageWithRwMutex(user string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[Sender] Sending message to %s...\n", user)
	time.Sleep(1 * time.Second)

	message := fmt.Sprintf("Message sent to %s", user)
	ch <- message

	// 🔒 Write to shared log (exclusive lock)
	models.LogMutex.Lock()
	models.MessageLog[user] = message
	models.LogMutex.Unlock()
	fmt.Printf("[Sender] Message recorded in log for %s\n", user)
}

func ReceiveConfirmationWithRwMutex(user string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	fmt.Printf("[Receiver] Confirmation for %s: %s\n", user, msg)

	// 🔓 Read from shared log (non-exclusive lock)
	models.LogMutex.RLock()
	logEntry, exists := models.MessageLog[user]
	models.LogMutex.RUnlock()

	if exists {
		fmt.Printf("[Receiver] Read from log: %s\n", logEntry)
	} else {
		fmt.Printf("[Receiver] No log found for %s\n", user)
	}
}

// sendMessages uses a custom random generator to avoid global rand.Seed
func SendMessagesWithSelectStatement(normalCh chan string, highCh chan string) {
	users := []string{"Sandeep", "Anjali", "Rahul", "Maya", "Zara"}

	// Create a new source and generator
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for _, user := range users {
		time.Sleep(time.Millisecond * 300) // simulate delay

		// Use custom rand instance
		if r.Intn(2) == 0 {
			highCh <- fmt.Sprintf("🚨 High-priority message for %s", user)
			// 👉 Example: 🚨 High-priority message for Sandeep
		} else {
			normalCh <- fmt.Sprintf("📩 Normal message for %s", user)
			// 👉 Example: 📩 Normal message for Anjali
		}
	}

	time.Sleep(time.Second) // wait before closing to allow message pickup
	close(highCh)
	close(normalCh)
}

func HandleMessagesWithSelectStatement(normalCh chan string, highCh chan string) {
	for {
		select {
		case msg, ok := <-highCh:
			if ok {
				fmt.Println("[HIGH] ", msg)
			} else {
				highCh = nil // disable select case
			}
		case msg, ok := <-normalCh:
			if ok {
				fmt.Println("[NORMAL]", msg)
			} else {
				normalCh = nil
			}
		}

		if highCh == nil && normalCh == nil {
			break
		}
	}
}
