package stack

import (
	"fmt"
	"log"

	"github.com/fbac/calc-o-matic/pkg/grammar"
)

// Stack implementation
type Stack []grammar.Token

// NewStack returns the memory address of a new stack
func NewStack() Stack {
	return Stack{}
}

// isValid checks if a stack contains elements
func (s Stack) isValid() bool {
	return len(s) > 0
}

// First returns the first token in the stack and returns the updated stack
func (s Stack) First() (grammar.Token, Stack, error) {
	var token grammar.Token

	if s.isValid() {
		token = s[0]
		s = s[1:]
		return token, s, nil
	}

	return token, s, fmt.Errorf("invalid stack")
}

// Pop returns the last token in the Stack and returns the updated stack
func (s Stack) Pop() (grammar.Token, Stack, error) {
	var token grammar.Token

	if s.isValid() {
		token = s[len(s)-1]
		s = s[:len(s)-1]
		return token, s, nil
	}

	return token, s, fmt.Errorf("invalid stack")
}

// Shift handles stack.Pop()
func (s Stack) Shift() (grammar.Token, Stack) {
	tmpToken, tmpStack, err := s.Pop()
	if err != nil {
		log.Fatalln(err)
	}
	return tmpToken, tmpStack
}

// Merge merges two stacks and returns the merged
func (orig Stack) Merge(merge Stack) Stack {
	var mergedStack Stack
	for range orig {
		tmpToken, tmpStack, err := orig.Pop()
		if err != nil {
			log.Fatalln(err)
		}

		orig = tmpStack
		mergedStack = append(merge, tmpToken)
	}
	return mergedStack
}

// SetStackPriority returns the priority of the last element of the stack
func SetStackPriority(s Stack) int {
	if len(s) > 0 {
		return GetTokenPriority(s[len(s)-1].V)
	} else {
		return 0
	}
}

// GetTokenPriority returns the priority of a given symbol
func GetTokenPriority(s string) int {
	switch s {
	case "(":
		return 1
	case "+":
		return 2
	case "-":
		return 2
	case "*":
		return 3
	case "/":
		return 3
	case "^":
		return 4
	default:
		return 0
	}
}

// PrioIsBigger returns true if first priority provided is bigger than last priority
func PrioIsBigger(currPrio int, lastPrio int) bool {
	return currPrio > lastPrio
}
