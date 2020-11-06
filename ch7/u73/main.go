// 7.3

package main

import (
	"bytes"
	"fmt"
	"gopl.io/ch4/treesort"
	"strconv"
)

type tree struct {
	treesort.Tree
}

func (t *tree) String() string {
	var buf bytes.Buffer

	var f func(t *treesort.Tree)

	f = func(t *treesort.Tree) {
		if t != nil {
			f(t.Left())
			buf.WriteString(strconv.Itoa(t.Value()))
			buf.WriteByte(' ')
			f(t.Right())
		}
	}
	f(&t.Tree)

	return buf.String()
}

func main() {
	t := treesort.Sort([]int{43, 23, 67, 12, 23, 76, 89, 5, 3, 8, 2, 1})

	fmt.Println(t)
}
