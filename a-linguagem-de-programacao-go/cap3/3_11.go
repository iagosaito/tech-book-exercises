package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("100000"))
	fmt.Println(comma("1000.00"))
	fmt.Println(comma("1000000000.00000000"))

}

func comma(s string) string {

	var buf bytes.Buffer

	i, d := splitDecimal(s)

	if len(i) <= 3 {
		return s
	}

	n := len(i) / 3
	for ind, v := range i {

		if ind == len(i)-n*3 {
			n--
			if ind == 0 {
				buf.WriteRune(v)
				continue
			}
			buf.WriteByte(',')
		}

		buf.WriteRune(v)
	}

	buf.WriteString(d)

	return buf.String()
}

func splitDecimal(s string) (i string, d string) {
	if index := decimalIndex(s); index != -1 {
		return s[:index], s[index:]
	}

	return s, ""
}

func decimalIndex(s string) int {
	return strings.Index(s, ".")
}
