// bit project bit.go
package main

import (
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func main() {
	var t IntSet
	t.Add(33)
	t.Add(43)
	t.Add(1023)
	//t.Add(1025)
	//t.Add(1025)
	//fmt.Println(t.Has(1025))
	fmt.Println(32 % 64)
}
