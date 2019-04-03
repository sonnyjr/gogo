// Package lexan
package lexan

import "fmt"

// TokenKind is just an int
type TokenKind int

type Token struct {
	kind TokenKind
	value string
	filename string	
	byte int
	size uint16
}

const (
	comment TokenKind = iota
	char TokenKind = iota
	newline TokenKind = iota
)

func (t TokenKind) String() string {
	return [...]string{
		"COMMENT",
		"CHAR",
		"NEWLINE"}[t]
}


func (t Token) String() string {
	return fmt.Sprintf("%s:%d:%d:%s:%s", t.filename, t.byte, t.size, t.kind, t.value)
}
