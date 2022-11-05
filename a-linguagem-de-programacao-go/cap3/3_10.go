package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("100"))
	fmt.Println(comma("1000"))
	fmt.Println(comma("10000"))
	fmt.Println(comma("100000"))
	fmt.Println(comma("1000000"))
	fmt.Println(comma("10000000"))
	fmt.Println(comma("100000000"))
}

func comma(s string) string {

	var buf bytes.Buffer

	if len(s) <= 3 {
		return s
	}

	n := len(s) / 3
	for i, v := range s {

		if i == len(s)-n*3 {
			n--
			if i == 0 {
				buf.WriteRune(v)
				continue
			}
			buf.WriteByte(',')
		}

		buf.WriteRune(v)
	}

	return buf.String()
}
