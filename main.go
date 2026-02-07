package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://www.youtube.com",
		"https://linkedin.com",
		"https://github.com",
		"https://microsoft.com",
	}

	var wg sync.WaitGroup
	inicio := time.Now()

	fmt.Println("Inciando processamento")

	for _, url := range urls {
		wg.Add(1)
		go verifyServiceStatus(url, &wg)
	}

	wg.Wait()

	fmt.Printf("Processamento finalizado. Tempo: %.2f", time.Since(inicio).Seconds())
}

func verifyServiceStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	_, err := http.Get(url)

	periodo := time.Since(start).Seconds()
	if err != nil {
		fmt.Printf("Erro ao consultar serviço [%s] [%.2fs]\n", url, periodo)
	} else {
		fmt.Printf("Sucesso ao consultar serviço [%s] [%.2fs]\n", url, periodo)
	}

}
