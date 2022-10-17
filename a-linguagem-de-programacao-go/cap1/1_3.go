package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startSimpleEcho := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	endSimpleEcho := time.Now()
	resultSimpleEcho := endSimpleEcho.Sub(startSimpleEcho)

	fmt.Println("Simple echo time:", resultSimpleEcho)

	startBetterEcho := time.Now()

	fmt.Println(strings.Join(os.Args, " "))

	endBetterEcho := time.Now()

	resultBetterEcho := endBetterEcho.Sub(startBetterEcho)

	fmt.Println("Better Echo time:", resultBetterEcho)
}
