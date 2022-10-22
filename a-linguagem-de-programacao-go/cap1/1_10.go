package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	f := createLogFile("fetchall.log")
	defer f.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch, f) // Inicia a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, f *os.File) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // Envia para o canal ch
		return
	}

	nbytes, err := io.Copy(f, resp.Body)

	if _, err := f.Write([]byte("\n\n")); err != nil { // Cria separação entre os logs
		f.Close()
		panic(err)
	}

	resp.Body.Close() // Evita o vazamento de bytes

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func createLogFile(filename string) *os.File {
	f, err := os.OpenFile("fetchall.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	return f
}
