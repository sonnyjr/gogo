package lexan

import (
	"unicode"
	"log"
	"io"
)

func (l * Analyzer) parseWord() *Token {
	startingLine := l.source.state.lineNumber
	startingByte := l.source.state.bytePosition
	var size int = 0
	value := make([]rune, 2)
	
	for {
		reps, err := l.source.Peek(1)
	
		if err == io.EOF || unicode.IsSpace(reps[0].char){
			break
		}
		
		if err != nil {
			log.Fatal(err)
		}		

		rune, s, err := l.source.Read()

		if err != nil {
			log.Fatal(err)
		}
		
		size += s
		value = append(value, rune)
	}

	if size == 0 {
		return nil
	}	
	
	return &Token{
		kind: word,
		value: string(value),
		filename: l.source.filename,
		line: startingLine,
		byte: startingByte,
		size: size}
}

