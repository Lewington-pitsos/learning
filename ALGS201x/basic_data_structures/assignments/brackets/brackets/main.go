package main

import (
	"fmt"
	"strconv"
)

func main() {
	mapping := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	fmt.Println(getStatus("{}[]", mapping))
	fmt.Println(getStatus("[()]", mapping))
	fmt.Println(getStatus("(())", mapping))
	fmt.Println(getStatus("{[]}()", mapping))

	fmt.Println(getStatus("{", mapping))
	fmt.Println(getStatus("{[}", mapping))
	fmt.Println(getStatus("foo(bar);", mapping))
	fmt.Println(getStatus("foo(bar[i);", mapping))
}

type indexedElement struct {
	index int
	value string
}

type stack struct {
	pile []indexedElement
}

func (s *stack) add(str indexedElement) {
	s.pile = append(s.pile, str)
}

func (s *stack) isEmpty() bool {
	return len(s.pile) <= 0
}

func (s *stack) take() indexedElement {
	str := s.read()

	s.pile = s.pile[:len(s.pile)-1]

	return str
}

func (s *stack) read() indexedElement {
	return s.pile[len(s.pile)-1]
}

func newStack() *stack {
	return &stack{
		[]indexedElement{},
	}
}

func getStatus(code string, mapping map[string]string) string {

	stk := newStack()

	for index, char := range code {
		for openBrace, closeBrace := range mapping {
			element := indexedElement{index, string(char)}
			if element.value == closeBrace {
				if !stk.isEmpty() && stk.read().value == openBrace {
					stk.take()
				} else {
					return strconv.Itoa(index + 1)
				}
			} else if element.value == openBrace {
				stk.add(element)
			}
		}
	}

	if stk.isEmpty() {
		return "Success"
	}

	for {
		var unmatched indexedElement
		if stk.isEmpty() {
			return strconv.Itoa(unmatched.index + 1)
		}
		unmatched = stk.take()
	}
}
