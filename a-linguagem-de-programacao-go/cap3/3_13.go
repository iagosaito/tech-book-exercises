package main

import "fmt"

const (
	KB = "1000"
	MB = KB + "000"
	GB = MB + "000"
	TB = GB + "000"
	PB = TB + "000"
	EB = PB + "000"
	ZB = EB + "000"
	YB = ZB + "000"
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
