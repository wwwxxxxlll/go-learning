// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	filename := os.Args[1]
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Printf("%s", err.Error())
// 	} else {
// 		defer file.Close()
// 		for k, v := range countlines(file) {
// 			fmt.Printf("%d\t%s\n", v, k)
// 		}

// 	}
// }
// func countlines(f *os.File) map[string]int {
// 	counts := make(map[string]int)
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		str := scanner.Text()
// 		counts[str]++
// 	}
// 	return counts
// }
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		str := input.Text()
		counts[str]++
		if counts[str] > 1 {
			fmt.Println(f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
