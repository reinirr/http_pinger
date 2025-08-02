package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func PingUrl(ctx context.Context, url string, interval int, timeouts map[string]int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	client := &http.Client{
		Timeout: time.Duration(timeouts["http_timeout"]) * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Duration(timeouts["connection_timeout"]) * time.Second,
			}).DialContext,
			ResponseHeaderTimeout: time.Duration(timeouts["read_timeout"]) * time.Second,
			IdleConnTimeout:       30 * time.Second,
		},
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stopping ping for %s\n", url)
			return
		case <-ticker.C:
			reqCtx, cancel := context.WithTimeout(ctx, time.Duration(timeouts["http_timeout"])*time.Second)

			req, err := http.NewRequestWithContext(reqCtx, "GET", url, nil)
			if err != nil {
				fmt.Printf("Ошибка создания запроса для %s: %v\n", url, err)
				cancel()
				continue
			}

			resp, err := client.Do(req)
			cancel()
			if err != nil {
				fmt.Printf("Cannot ping %s: %v\n", url, err)
				WriteLog(url, fmt.Sprintf("ERROR: %v", err))
				continue
			}

			defer resp.Body.Close()
			status := resp.Status
			WriteLog(url, status)
		}
	}
}
