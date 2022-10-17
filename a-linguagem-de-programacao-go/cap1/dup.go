package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%d\t%s\n", v, k)
	}
}
