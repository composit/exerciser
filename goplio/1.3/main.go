package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	run1 := time.Since(start1).Seconds()
	fmt.Printf("echo1: %f seconds\n", run1)

	start2 := time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	run2 := time.Since(start2).Seconds()
	fmt.Printf("echo2: %f seconds\n", run2)

	start3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	run3 := time.Since(start3).Seconds()
	fmt.Printf("echo3: %f seconds\n", run3)
}
