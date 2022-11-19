/*
grammar package is independent from the rest of packages
to make it easy to extend in the future
*/
package grammar

// TokenType represents an operation or a digit
type TokenType int

// Constant values are needed to define the grammar
const (
	OPRTN TokenType = iota // 0
	DIGIT                  // 1
)

// Token holds the value of a given operation/number
type Token struct {
	T TokenType
	V string
}
