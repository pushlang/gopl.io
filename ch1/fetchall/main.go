// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func readFromFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\r\n"), nil
}

func main() {
	start := time.Now()
	ch := make(chan string)

	var urlList []string

	if len(os.Args) == 2 {
		var err error
		urlList, err = readFromFile(os.Args[1]) //urls.txt
		fmt.Fprintf(os.Stderr, "readFromFile: %v\n", err)
	} else {
		urlList = []string{"http://yaaaaaaa.ru", "http://mail.ru", "http://google.ru", "http://rambler.ru"}
	}

	//

	for _, url := range urlList { //
		go fetch("http://"+url, ch) // start a goroutine
	}
	for range urlList {
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

	// nbytes, err := io.Copy(ioutil.Discard, resp.Body) // after calling, resp.Body is empty
	text, err := ioutil.ReadAll(resp.Body)
	nbytes, err := writeToFile(text, url)
	if err != nil {
		fmt.Println(err)
	}

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func writeToFile(b []byte, fn string) (int, error) {
	filename := strings.Split(fn, "//")

	f, err := os.Create(filename[1] + ".html")
	if err != nil {
		return 0, err
	}

	n, err := f.Write(b)
	if err != nil {
		f.Close()
		return 0, err
	}

	err = f.Close()
	if err != nil {
		return 0, err
	}

	return n, nil
}
