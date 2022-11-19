package stack

// RuneStack type
type IntStack []int

// NewIntStack returns a new IntStack
func NewIntStack() IntStack {
	return []int{}
}

// isValid checks if a stack contains elements
func (i IntStack) isValid() bool {
	return len(i) > 0
}

// Pop returns the last token in the Stack and returns the updated stack
func (i IntStack) Pop() (int, IntStack) {
	var x int
	if i.isValid() {
		x = i[len(i)-1]
		i = i[:len(i)-1]
	}
	return x, i
}

// UpdateResult is a copy of Pop, it only does not return the updated Stack
func (i IntStack) UpdateResult() int {
	var x int
	if i.isValid() {
		x = i[len(i)-1]
	}
	return x
}
