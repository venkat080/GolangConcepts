package concepts

import (
	"fmt"
	"time"
)

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
