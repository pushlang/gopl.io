package main
import "fmt"

type mystring string

func (m *mystring) join() {
    fmt.Printf("%s\n", m)
    if *m != "" {
        mm:=*m
        //m0:=mm[0]
        *m=mm[1:]
        m.join()
        fmt.Printf("%s\n", mm[1:])
    }
}

func (m *mystring) String() string{
    return string(*m)
}

func main() {
    ms:=mystring("abcdefgh")
    
    ms.join()
}