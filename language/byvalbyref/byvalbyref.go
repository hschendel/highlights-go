package main

import (
	"fmt"
)

func main() {
	var s SomeStruct
	s.println()

	byVal(s)
	s.println()

	byRef(&s)
	s.println()

	s.setA("new")
	s.println()

	s.setB("new")
	s.println()
}

type BSetter interface {
	setB(b string)
}

type SomeStruct struct {
	A string
	B string
}

func byVal(s SomeStruct) {
	s.A = "val"
}

func byRef(s *SomeStruct) {
	s.A = "ref"
}

// Methods

func (s SomeStruct) setA(a string) {
	s.A = a
}

func (s *SomeStruct) setB(b string) {
	s.B = b
}

func (s *SomeStruct) println() {
	fmt.Printf("A: %q, B: %q\n\n", s.A, s.B)
}
