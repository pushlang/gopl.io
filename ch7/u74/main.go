// 7.4, 7.5

package main

import (
	"bufio"
	"fmt"
	"io"

	"gopl.io/ch5/outline"

	_ "gopl.io/ch5/outline"
)

type StrRead struct {
	str string
}
type StrLimit struct {
	io.Reader
	n int64
}

func (s StrRead) Read(p []byte) (n int, err error) {
	l := 0
	if lp, ls := len(p), len(s.str); ls > lp {
		l = lp
	} else {
		l = ls
	}
	copy(p, s.str[:l])
	//for i := 0; i < len(p) && i < len(s.str); i++ {
	//	p[i] = s.str[i]
	//}
	if len(p) < len(s.str) {
		return len(p), nil
	}
	return len(s.str), io.EOF
}
func (s StrLimit) Read(p []byte) (n int, err error) {
	n, _ = s.Reader.Read(p[:s.n])
	return n, io.EOF
}
func NewReader(s string) io.Reader {
	return &StrRead{s}
}
func LimitReader(r io.Reader, n int64) io.Reader {
	return &StrLimit{Reader: r, n: n}
}

func main() {
	s := `<html>
<head>
<title>
Title text
</title>
</head>
<body>
<a href="http://ya.ru">link ya.ru</a> 
</body>
</html>
`
	sr := NewReader(s)
	lr := LimitReader(sr, 20)

	outline.Run(sr)

	fmt.Println()

	sc := bufio.NewScanner(lr)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		fmt.Printf("%s\n", sc.Text())
	}
}
