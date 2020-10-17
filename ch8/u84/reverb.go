package main

import (
	"bufio"	
	"strings"
	"fmt"
	"log"
	"net"
	"time"
	"sync"
)

func echo(c *net.TCPConn, delay time.Duration, done chan struct{}, shout string, wg sync.WaitGroup) {
		c.Write([]byte("\t" + strings.ToUpper(shout) + "\n"))
		time.Sleep(delay)
		c.Write([]byte("\t" + shout + "\n"))
		time.Sleep(delay)
		c.Write([]byte("\t" + strings.ToLower(shout) + "\n"))
		//done <- struct{}{}
		wg.Done()
}

func testMessage(c *net.TCPConn) {
	i := 0
	for {
		i++
		str:=fmt.Sprintf("Test string (%d) from server!\n", i)
		c.Write([]byte(str))
		time.Sleep(1*time.Second)
	}
}

func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)

	done := make(chan struct{})
	
	var wg sync.WaitGroup
	
	for input.Scan(){
		wg.Add(1)
		go echo(c, 5*time.Second, done, input.Text(), wg)	
	}
	
	//<-done
	wg.Wait()
	c.Close()	
	fmt.Println("Disconnected from client!")
}

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
		if err != nil {
			log.Print(err) 
			continue
		}
		fmt.Println("Connected to client!")
		go handleConn(conn)
	}
}
