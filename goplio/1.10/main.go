// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	fileName := fmt.Sprintf("fetchall-%d.txt", time.Now().UnixNano())
	out, err := os.Create(filepath.Join(wd, "1.10", fileName))
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer func() {
		err := out.Close()
		if err != nil {
			ch <- fmt.Sprint(err)
		}
	}()

	nbytes, err := io.Copy(out, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
