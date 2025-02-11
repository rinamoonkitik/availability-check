package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//Проверка доступности тип

func checkStatus(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Ошибка при запросе к сайту")
		return false
	}
	defer resp.Body.Close()

	//Если доступен
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("Сайт доступен!")
		return true
	}
	//Если недоступен
	fmt.Printf("Сайт недоступн! Статус: %d\n", resp.StatusCode)
	return false
}

func main() {

	siteURL := "***"

	//Интервал
	checkInterval := 60 * time.Hour

	//Бесконечный цикл для проверка
	//go func() {
	for {
		isCheck := checkStatus(siteURL)

		if !isCheck {
			log.Println("Походу упал...")
		}
		time.Sleep(checkInterval)

	}
}
