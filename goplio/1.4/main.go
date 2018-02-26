package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupFiles := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, dupFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, dupFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for name, _ := range dupFiles[line] {
				fmt.Printf("%s ", name)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, fileName string, counts map[string]int, dupFiles map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if _, ok := dupFiles[input.Text()]; !ok {
			dupFiles[input.Text()] = make(map[string]bool)
		}
		dupFiles[input.Text()][fileName] = true
	}
}
