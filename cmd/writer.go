package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var logMutex sync.Mutex

func WriteLog(url, status string) {
	logMutex.Lock()
	defer logMutex.Unlock()
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}
	defer file.Close()

	entry := fmt.Sprintf("%s [%s] %s\n", time.Now().Format(time.RFC3339), url, status)
	if _, err := file.WriteString(entry); err != nil {
		log.Fatalf("error while writing to file: %v", err)
	}
	fmt.Println("Writing success!")
}
