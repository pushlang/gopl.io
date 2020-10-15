// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c *net.TCPConn, delay time.Duration, shout string, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Write([]byte("\t" + strings.ToUpper(shout) + "\n"))
	time.Sleep(delay)
	c.Write([]byte("\t" + shout + "\n"))
	time.Sleep(delay)
	c.Write([]byte("\t" + strings.ToLower(shout) + "\n"))
}

//!+
func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	
	var wg sync.WaitGroup
	
	start := make(chan struct{})
	for input.Scan(){
                wg.Add(1)
		start<-struct{}{}
		go echo(c, 1*time.Second, input.Text(), &wg)
	}
	

	go func() {
		<-start
		wg.Wait()
		c.CloseRead()
		fmt.Println("Disconnected!")
	}()

	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	tcpAddr := &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 8000,
	}

	l, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.AcceptTCP()
		fmt.Println("Connected!")
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
