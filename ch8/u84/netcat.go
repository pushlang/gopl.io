package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	ra := &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 8000,
	}
	conn, err := net.DialTCP("tcp", nil, ra)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println()
		log.Println("Connection has been closed by server!")
		log.Println("Writing to 'Stdout' from 'conn' stopped!")
		done <- struct{}{}
	}()
	// mustCopy(conn, os.Stdin)
	testMess(conn, "AbCdEfjHiJkLmNoPqRsTuVyXyZ")
	
	conn.CloseWrite()
	log.Println("Writing side of connection closed by client!")
	time.Sleep(30*time.Second)
	<-done
	log.Println("Program stopped!")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	log.Println("Writing to 'conn' from 'Stdin' stopped!")
}

func testMess(dst io.Writer, mess string) {
	for i:=0; i<len(mess); i++ {
		dst.Write([]byte(mess[i:i+1]))
		dst.Write([]byte{' ', '\n'})
		time.Sleep(1*time.Second)
	}
	log.Println("Writing to 'conn' from 'Stdin' stopped!")
}
