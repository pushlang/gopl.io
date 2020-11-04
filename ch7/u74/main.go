package main

import (
	"gopl.io/ch5/outline"
	"io"
)

type StrRead struct {
	str string
}

func (s StrRead) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p) && i < len(s.str); i++ {
		p[i] = s.str[i]
	}
	if len(p) < len(s.str) {
		return len(p), nil
	}
	return len(s.str), io.EOF
}
func NewReader(s string) *StrRead {
	return &StrRead{s}
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

	outline.Run(sr)

	//sc := bufio.NewScanner(sr)
	//sc.Split(bufio.ScanWords)
	//for sc.Scan() {
	//	fmt.Printf("%s ", sc.Text())
	//}
}
