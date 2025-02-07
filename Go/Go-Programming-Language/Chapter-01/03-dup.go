package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// 01 : Read from stdin
	//counts := make(map[string]int)
	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	counts[input.Text()]++
	//}
	//for line, n := range counts {
	//	if n > 1 {
	//		fmt.Printf("%d\t%s\n", n, line)
	//	}
	//}

	//02: Read from stdin or files
	//counts := make(map[string]int)
	//files := os.Args[1:]
	//if len(files) == 0 {
	//	countLines(os.Stdin, counts)
	//} else {
	//	for _, f := range files {
	//		file, err := os.Open(f)
	//		if err != nil {
	//			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	//			continue
	//		}
	//		countLines(file, counts)
	//		file.Close()
	//	}
	//}
	//for line, n := range counts {
	//	if n > 1 {
	//		fmt.Printf("%d\t%s\n", n, line)
	//	}
	//}

	// 03: Simplified
	counts := make(map[string]int)
	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file) // Deprecated, use os.ReadFile instead
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

//func countLines(f *os.File, counts map[string]int) {
//	input := bufio.NewScanner(f)
//	for input.Scan() {
//		counts[input.Text()]++
//	}
//}
