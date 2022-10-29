package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/iagosaito/converter/src/converter"
)

func main() {

	if len(os.Args[1:]) > 0 {
		for _, v := range os.Args[1:] {

			r, err := strconv.ParseFloat(v, 64)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro to convert: %v\n", err)
				os.Exit(1)
			}

			converter.ShowAllConvertions(r)
		}
	} else {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			r, err := strconv.ParseFloat(s.Text(), 64)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro to convert: %v\n", err)
				os.Exit(1)
			}

			converter.ShowAllConvertions(r)
		}
	}
}
