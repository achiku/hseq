package hseq

import (
	"fmt"
	"sync"
)

// Seq Sequence
type Seq struct {
	Name string
	nums []int64
}

// String returns max value
func (s *Seq) String() string {
	return fmt.Sprintf("%d", s.nums[len(s.nums)-1])
}

// Int returns max value
func (s *Seq) Int() int {
	return int(s.nums[len(s.nums)-1])
}

// Int64 returns max value
func (s *Seq) Int64() int64 {
	return s.nums[len(s.nums)-1]
}

var seqList []*Seq
var mu sync.Mutex

// Get returns sequence
func Get(name string) *Seq {
	mu.Lock()
	defer mu.Unlock()
	for _, s := range seqList {
		if s.Name == name {
			s.nums = append(s.nums, s.nums[len(s.nums)-1]+1)
			return s
		}
	}
	seq := &Seq{
		Name: name,
		nums: []int64{1},
	}
	seqList = append(seqList, seq)
	return seq
}
