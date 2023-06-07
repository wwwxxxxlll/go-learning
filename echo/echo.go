package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// var str, sep string
	// for _, s := range os.Args[1:] {
	// 	str += sep + s
	// 	sep = " "
	// }
	str := strings.Join(os.Args[1:], " ")
	fmt.Println(str)
}
