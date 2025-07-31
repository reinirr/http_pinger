package cmd

import (
	"fmt"
	"net/http"
	"time"
)

func PingUrl(url string) {
	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			resp, err := http.Get(url)

			if err != nil {
				fmt.Printf("Cannot ping %s", url)
			}

			defer resp.Body.Close()

			status := resp.Status
			WriteLog(url, status)
		}
	}

}
