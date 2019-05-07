// Package lexan
package lexan

import "fmt"

// TokenKind is just an int
type TokenKind int

// The Token type holds the TokenKind that was found,
// it's value, and location information within the
// source files.
type Token struct {
	kind TokenKind
	value string
	filename string
	line int
	byte int
	size int
}

// Define the TokenKind constants
const (
	comment TokenKind = iota
	char
	word
)

// Convert's the TokenKind to it's string representation.
func (t TokenKind) String() string {
	return [...]string{
		"COMMENT",
		"CHAR",
		"WORD"}[t]
}

// Convert's the Token to it's string representation.
func (t Token) String() string {
	return fmt.Sprintf("%s:%d:%d:%s(%s)", t.filename, t.line, t.byte, t.kind, t.value)
}
