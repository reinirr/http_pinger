package main

import (
	"context"
	"fmt"
	"http_pinger/cmd"
	"http_pinger/interfaces"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gopkg.in/yaml.v3"
)

func main() {
	urls := cmd.GetUrls()

	if len(urls) == 0 {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nПолучен сигнал завершения, останавливаю пингеры...")
		cancel()
	}()

	interval := 5
	timeouts := map[string]int{
		"http_timeout":       30,
		"connection_timeout": 10,
		"read_timeout":       10,
		"write_timeout":      10,
	}

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Ошибка чтения config.yaml: %v", err)
	} else {
		var yamlCfg interfaces.Config
		err = yaml.Unmarshal(yamlFile, &yamlCfg)
		if err != nil {
			log.Printf("Ошибка парсинга config.yaml: %v", err)
		} else {
			interval = yamlCfg.IntervalPulling
			if yamlCfg.HTTPTimeout > 0 {
				timeouts["http_timeout"] = yamlCfg.HTTPTimeout
			}
			if yamlCfg.ConnectionTimeout > 0 {
				timeouts["connection_timeout"] = yamlCfg.ConnectionTimeout
			}
			if yamlCfg.ReadTimeout > 0 {
				timeouts["read_timeout"] = yamlCfg.ReadTimeout
			}
			if yamlCfg.WriteTimeout > 0 {
				timeouts["write_timeout"] = yamlCfg.WriteTimeout
			}
		}
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			cmd.PingUrl(ctx, u, interval, timeouts)
		}(url)
	}

	wg.Wait()
	fmt.Println("Все пингеры остановлены")
}
