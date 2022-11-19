package stack

// RuneStack type
type RuneStack []rune

// nextRune are functions to look for runes in strings
type nextRune func(rune) bool

// NewRuneStack returns a new runeStack based on an input string
func NewRuneStack(s string) RuneStack {
	return []rune(s)
}

// Read reads a RuneStack based on a nextRune function
// Returns the generated string and the updated RuneStack
func (r RuneStack) Read(f nextRune) (string, RuneStack) {
	var s = ""
	for len(r) > 0 && f(r[0]) {
		c, newRuneStack := r.Shift()
		r = newRuneStack
		s += string(c)
	}
	return s, r
}

// Shift returns the first rune in the stack and the updated RuneStack
func (r RuneStack) Shift() (rune, []rune) {
	return r[0], r[1:]
}
