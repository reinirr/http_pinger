package cmd

import (
	"bufio"
	"log"
	"net/url"
	"os"
	"strings"
)

func GetUrls() []string {
	file, err := os.Open("urls.txt")
	if err != nil {
		log.Printf("Ошибка открытия файла urls.txt: %v", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var urls []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if _, err := url.Parse(line); err != nil {
			log.Printf("Некорректный URL: %s - %v", line, err)
			continue
		}

		urls = append(urls, line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
	}

	if len(urls) == 0 {
		log.Println("Предупреждение: не найдено валидных URL в файле urls.txt")
	}

	return urls
}
