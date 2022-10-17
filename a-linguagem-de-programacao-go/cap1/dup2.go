package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		readFromStdin()
	} else {
		for _, filename := range files {
			readFromFile(filename)
		}
	}
}

func readFromStdin() {
	countLines(os.Stdin)
}

func readFromFile(filename string) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	countLines(file)
}

func countLines(f *os.File) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%d\t%s\n", v, k)
	}
}
