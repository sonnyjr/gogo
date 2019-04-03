package lexan

import (
	"io"
	"log"
)

func (l * Analyzer) parseComment() *Token {
	comment := l.ParseSingleLineComment()

	if comment != nil {
		return comment
	}

	return l.ParseMultilineComment()	
}

func (l * Analyzer) ParseSingleLineComment() * Token {
	if !l.source.Match ("//") {
		return nil
	}

	startingByte := l.source.bytePosition
	
	var size uint16 = 2
	value := make([]rune, 2)
	
	for {
		rune, _, err := l.source.Read()
									
		if err == io.EOF || rune == '\n' {
			break
		}

		value = append(value, rune)
	}

	return &Token{kind: comment,
		value: string(value),
		filename: l.filename,
		byte: startingByte,
		size: size}	
}

func (l * Analyzer) ParseMultilineComment() * Token {
	if !l.source.Match ("/*") {
		return nil
	}
	
	var size uint16 = 2
	value := make([]rune, 2)

	startingByte := l.source.bytePosition
	slash, _, _ := l.source.Read()
	star, _, _ := l.source.Read()	
	value = append(value, slash)
	value = append(value, star)	
	
	for {
		rune, _, err := l.source.Read()
									
		if err == io.EOF {
			return nil
		}
		
		value = append(value, rune)
		
		if rune == '*' {
			runes, err := l.source.Peek(1)
			
			if err != nil {
				log.Fatal(err)
			}

			if runes[0].r == '/' {
				value = append(value, runes[0].r)
				break
			}
		}
	}

	return &Token{kind: comment,
		value: string(value),
		filename: l.filename,
		byte: startingByte, 
		size: size}	
}
