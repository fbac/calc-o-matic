package parser

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/fbac/calc-o-matic/pkg/grammar"
	"github.com/fbac/calc-o-matic/pkg/stack"
)

// functions to mutate strings
type stringMutator func(r rune) rune

// ParseExpr parsers the inputExpr and returns a valid Stack
func ParseExpr(inputExpr string) (stack.Stack, error) {
	// Trim whitespaces with strings.Map
	newExpr := mutateString(inputExpr, isWhitespace)

	// Check valid length
	if !isValidLength(newExpr) {
		return nil, fmt.Errorf("invalid expression [%v]: length not valid", inputExpr)
	}

	// Check valid characters
	if !isValidInput(newExpr) {
		return nil, fmt.Errorf("invalid expression [%v]: variables not allowed", inputExpr)
	}

	// Initialize rune and token stacks based on healthy math expression
	// Initialize a new token which will be holding temporary values
	runeStack := stack.NewRuneStack(newExpr)
	tokenStack := stack.NewStack()
	newToken := grammar.Token{}

	// Loop until runeStack is empty
	for len(runeStack) > 0 {
		// Pop the first rune, and update runeStack
		r, newRuneStack := runeStack.Shift()
		runeStack = newRuneStack

		// newToken is created based on currToken
		currToken := string(r)
		if isNum(currToken) {
			num, newRuneStack := runeStack.Read(runeIsNumber)
			runeStack = newRuneStack
			newToken = grammar.Token{T: grammar.DIGIT, V: currToken + num}
		} else if isOp(currToken) {
			newToken = grammar.Token{T: grammar.OPRTN, V: currToken}
		}

		// update tokenStack with newToken
		tokenStack = append(tokenStack, newToken)
	}

	return tokenStack, nil
}

// CreateAST creates the abstract syntax tree based on a Stack
// Not really a completely compliant AST, though.
// It maintains current and last priorities
func CreateAST(inputStack stack.Stack) (stack.Stack, error) {
	if len(inputStack) <= 0 {
		return nil, fmt.Errorf("invalid stack [%v]: empty stack", inputStack)
	}
	// Initialize operations stack and return stack
	opeStack := stack.NewStack()
	retStack := stack.NewStack()

	// tmpToken hold temporary values to append to retStack
	tmpToken := grammar.Token{}

	for len(inputStack) > 0 {
		// Get the current token and updated stack
		currToken, currStack, err := inputStack.First()
		if err != nil {
			return nil, fmt.Errorf("invalid stack [%v]: error reading first element", inputStack)
		}
		inputStack = currStack

		// Get priorities for current token and last operation
		currPrio := stack.GetTokenPriority(currToken.V)
		lastPrio := stack.SetStackPriority(opeStack)

		switch currToken.T {
		case grammar.DIGIT:
			retStack = append(retStack, currToken)
		case grammar.OPRTN:
			// If it's the first element, (, or currPrio is bigger than lastPrio, simply insert the element in stack
			if len(opeStack) == 0 || currToken.V == "(" || stack.PrioIsBigger(currPrio, lastPrio) {
				opeStack = append(opeStack, currToken)
			} else if currToken.V == ")" {
				// when the item is ), opeStack has to be shuffled
				// and look for the next ( or EOL
				tmpToken, opeStack = opeStack.Shift()
				for tmpToken.V != "(" && len(opeStack) != 0 {
					retStack = append(retStack, tmpToken)
					tmpToken, opeStack = opeStack.Shift()
				}
			} else {
				// when the last priority is bigger than the current one:
				// shuffle the stack, so we take the last token and return the opeStack
				// Then, reset priorities based on new opeStack
				// Finally, add the currToken to continue growing the stack.
				for !stack.PrioIsBigger(currPrio, lastPrio) {
					tmpToken, opeStack = opeStack.Shift()
					retStack = append(retStack, tmpToken)
					lastPrio = stack.SetStackPriority(opeStack)
				}
				opeStack = append(opeStack, currToken)
			}
		}
	}
	// Merge opeStack  with retStack and return it
	return opeStack.Merge(retStack), nil
}

/* Checker functions */

// stringIsNum checks if string is number
func isNum(s string) bool {
	return regexp.MustCompile(`\d+`).MatchString(s)
}

// isOp checks if string is a valid operation
func isOp(s string) bool {
	return regexp.MustCompile(`\(|\)|\-|\+|\/|\*|\^`).MatchString(s)
}

// runeIsNumber checks if rune is number
func runeIsNumber(r rune) bool {
	return unicode.IsDigit(r)
}

/* Validation checks */

// isValidLength returns if the input has a valid length
func isValidLength(s string) bool {
	return len(s) > 0
}

// isValidInput returns if the input contains any invalid characters
func isValidInput(s string) bool {
	return !regexp.MustCompile(`[a-zA-Z]`).MatchString(s)
}

/*
### Mutation functions
	strings.Map is really performant for this kind of processes
	Also, it's interesting to make this extensible in the future by adding funcs
*/

// mutateString mutates string based on stringMutator function
func mutateString(str string, f stringMutator) string {
	return strings.Map(f, str)
}

// isWhitespace is a stringMutator function
// used in conjunction with mutateString to remove all whitespaces
func isWhitespace(r rune) rune {
	if unicode.IsSpace(r) {
		return -1
	}
	return r
}
