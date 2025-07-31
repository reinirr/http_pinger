package main

import (
	"http_pinger/cmd"
	"os"
	"sync"
)

func main() {
	urls := cmd.GetUrls()

	if len(urls) == 0 {
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			cmd.PingUrl(url)
		}(url)
	}

	wg.Wait()
}
