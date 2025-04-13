package concepts

import (
	"fmt"
	"golangconcepts/models"
	"sync"
	"time"
)

func Channels() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
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
