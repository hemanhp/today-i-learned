package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}

	//fmt.Println(strings.Join(os.Args[1:], " "))

	// Exercise 1.1
	//fmt.Println(strings.Join(os.Args[:], " "))

	//Exercise 1.2
	//s := ""
	//for index, arg := range os.Args[1:] {
	//	s += fmt.Sprintf("%d:%s \n", index, arg)
	//}

	//Exercise 1.3
	//start := time.Now()
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Printf("%q %f\n", s, time.Since(start).Seconds())
	//
	//start = time.Now()
	//fmt.Println(strings.Join(os.Args[1:], " "))
	//fmt.Printf("%q %f\n", "Join Function", time.Since(start).Seconds())
	//
	//fmt.Println(s)
}
