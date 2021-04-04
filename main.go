package main

import "fmt"

func main() {
	var i uint8
	s := newStack(1)
	s = s.put(2)
	s = s.put(3)
	s = s.put(4)
	s, i = s.pop()
	fmt.Println(i)
	s, i = s.pop()
	fmt.Println(i)
	s, i = s.pop()
	fmt.Println(i)
	s, i = s.pop()
	fmt.Println(i)
	s, i = s.pop()
	fmt.Println(i)
	s, i = s.pop()
	fmt.Println(i)
	s, i = s.pop()
	fmt.Println(i)
	s, i = s.pop()
	fmt.Println(i)
	_, i = s.pop()
	fmt.Println(i)
}

type stack struct {
	value uint8
	next  *stack
}

func (s *stack) pop() (*stack, uint8) {
	if s == nil {
		return nil, 0
	}

	return s.next, s.value
}

func (s *stack) put(i uint8) *stack {
	ns := newStack(i)
	ns.next = s

	fmt.Println("a", ns)
	fmt.Println("a", ns.next.value)

	return ns
}

func newStack(v uint8) *stack {
	return &stack{value: v}
}
