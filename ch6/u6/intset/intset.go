//rm u6; goimports -v -w *.go; gofmt -w *.go; go build; ./u6
// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

func init() {

}

var bitness = 32<<(^uint(0)>>63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

func wordBit(value, bit int) (int, uint) {
	return value / bit, uint(value % bit)
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(value int) bool {
	word, bit := wordBit(value, bitness)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(value int) {
	word, bit := wordBit(value, bitness)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith1(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	s.With(t, func(i int, tword uint){
		s.words[i] |= tword
	})
}

func (s *IntSet) IntersectionWith(t *IntSet) {
	s.With(t, func(i int, tword uint){
		s.words[i] &= tword
	})
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	s.With(t, func(i int, tword uint){
		s.words[i] &^= tword
	})
}

func (s *IntSet) SymmDifferenceWith(t *IntSet) {
	s.With(t, func(i int, tword uint){
		s.words[i] ^= tword
	})
}

func (s *IntSet) With(t *IntSet, do func(int, uint)) {
	for i, tword := range t.words {
		if i < len(s.words) {
			do(i, tword)
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *IntSet) Copy() *IntSet {
	n := new(IntSet)
	s.iterate(func(p int, v int) (ok bool) {
		n.Add(v)
		return
	})
	return n
}

func (s *IntSet) Len() (l int) {
	s.iterate(func(p int, v int) (ok bool) {
		l = p
		return
	})
	return
}

func (s *IntSet) GetValue(position int) (value int) {
	s.iterate(func(p int, v int) (ok bool) {
		if p == position {
			value = v
			ok = true
		}
		return
	})
	return
}

func (s *IntSet) Elems() (elems []int) {
	s.iterate(func(p int, v int) (ok bool) {
		elems = append(elems, v)
		return
	})
	return
}

func (s *IntSet) GetPosition(value int) (position int) {
	s.iterate(func(p int, v int) (ok bool) {
		if value == v {
			position = p
			ok = true
		}
		return
	})
	return
}

func (s *IntSet) Remove(value int) {
	word, bit := wordBit(value, bitness)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() (value int) {
	s.iterate(func(p int, v int) (ok bool) {
		s.Remove(v)
		return
	})
	return
}

func (s *IntSet) iterate(do func(int, int) bool) {
	counter := 0
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				counter++
				if do(counter, 64*i+j) {
					break
				}
			}
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	s.iterate(func(p int, v int) (ok bool) {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", v)

		return
	})

	buf.WriteByte('}')
	return buf.String()
}

