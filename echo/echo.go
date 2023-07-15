package main

import (
	"flag"
	"fmt"
	"strings"
)

// func main() {
// 	// var str, sep string
// 	// for _, s := range os.Args[1:] {
// 	// 	str += sep + s
// 	// 	sep = " "
// 	// }
// 	str := strings.Join(os.Args[1:], " ")
// 	fmt.Println(str)
// }
// Echo4 prints its command-line arguments.

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	var unicode = '去'
	//[1] announce use the first operate number, # announce generate prefix like 0X,0x
	fmt.Printf("%d %[1]c %#[1]x", unicode)
}
