package eval

import (
	"strconv"

	"github.com/fbac/calc-o-matic/pkg/grammar"
	"github.com/fbac/calc-o-matic/pkg/stack"
)

// Evaluate postfix expression
func Calculate(inputStack stack.Stack) int {
	// Initialize variables
	retStack := stack.NewIntStack()
	currToken := grammar.Token{}
	x, y := 0, 0

	// Loop over inputStack to get values and operations
	for len(inputStack) > 0 {
		// Get first Token in stack
		currToken, inputStack, _ = inputStack.First()

		// Switch based on type
		switch currToken.T {

		case grammar.DIGIT:
			// If it's a number, convert it to int and add it to the stack
			newValue, _ := strconv.Atoi(currToken.V)
			retStack = append(retStack, newValue)
		case grammar.OPRTN:
			// If it's an operation:
			// Pop x, Pop y, and run DoFunc(OPE, NUM, NUM)
			// Update the current value by appending result to retStack
			x, retStack = retStack.Pop()
			y, retStack = retStack.Pop()
			retStack = append(retStack, DoFunc(currToken.V, x, y))
		}
	}

	// 
	return retStack.UpdateResult()
}
