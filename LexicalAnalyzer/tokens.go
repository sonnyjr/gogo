package lexicalAnalyzer

import "fmt"

type TokenKind int

type Token struct {
	kind TokenKind
	value string
	filename string
	byte uint16
	size uint16
}

const (
	comment TokenKind = iota
	rune TokenKind = iota
	newline TokenKind = iota
)

func (t TokenKind) String() string {
	return [...]string{
		"COMMENT",
		"RUNE",
		"NEWLINE"}[t]
}


func (t Token) String() string {
	return fmt.Sprintf("%s:%d:%d:%s:%s", t.filename, t.byte, t.size, t.kind, t.value)
}
