package main

type Stack []uint64

var LoopStack Stack

func (s *Stack) Push(data uint64) {
	tmp := append(*s, data)
	*s = tmp
}

func (s *Stack) Pull() uint64 {
	tmp1 := []uint64(*s)[len([]uint64(*s))-1]
	tmp2 := Stack([]uint64(*s)[0 : len([]uint64(*s))-1])
	*s = tmp2
	return tmp1
}
