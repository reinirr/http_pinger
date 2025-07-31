package cmd

import (
	"bufio"
	"log"
	"os"
)

func GetUrls() []string {
	file, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return urls
}
