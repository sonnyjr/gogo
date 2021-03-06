package lexan

import (
	"io"
	"log"
)

// Parses a comment.
func (l * Analyzer) parseComment() *Token {
	comment := l.ParseSingleLineComment()

	if comment != nil {
		return comment
	}

	return l.ParseMultilineComment()	
}

// Attempts to parse a single line comment.
func (l * Analyzer) ParseSingleLineComment() * Token {
	if !l.source.Match ("//") {
		return nil
	}

	startingLine := l.source.state.lineNumber
	startingByte := l.source.state.bytePosition
	
	var size int = 2
	value := make([]rune, 2)
	
	for {
		rune, s, err := l.source.Read()
									
		if err == io.EOF || rune == '\n' {
			break
		}
		
		size += s
		value = append(value, rune)
	}

	return &Token{kind: comment,
		value: string(value),
		filename: l.filename,
		line: startingLine,
		byte: startingByte,
		size: size}	
}

// Attempts to parse a multiline comment.
func (l * Analyzer) ParseMultilineComment() * Token {
	if !l.source.Match ("/*") {
		return nil
	}
	
	var size int = 2
	value := make([]rune, 2)

	startingLine := l.source.state.lineNumber
	startingByte := l.source.state.bytePosition
	slash, _, _ := l.source.Read()
	star, _, _ := l.source.Read()	
	value = append(value, slash)
	value = append(value, star)	
	
	for {
		rune, s, err := l.source.Read()
									
		if err == io.EOF {
			return nil
		}
		
		size += s		
		value = append(value, rune)
		
		if rune == '*' {
			runes, err := l.source.Peek(1)
			
			if err != nil {
				log.Fatal(err)
			}

			if runes[0].char == '/' {
				rune, s, _ := l.source.Read()
				size += s
				value = append(value, rune)
				break
			}
		}
	}

	return &Token{kind: comment,
		value: string(value),
		filename: l.filename,
		line: startingLine,
		byte: startingByte, 
		size: size}	
}
