package lexan

import (
	"bufio"
	"fmt"
	"io"
)

type Analyzer struct {
	source * Source
	filename string
}


func New(data * bufio.Reader, filename string) Analyzer {
	state := SourceState{previous: nil, next: nil}
	source := Source{data: data, filename: filename, state: &state}
	return Analyzer{source: &source, filename: filename}
}

func (l * Analyzer) Analyze(){
	for {
		comment := l.parseComment()

		if comment != nil {
			fmt.Printf("%s\n", comment)
		} else {
			word := l.parseWord()

			if word != nil {
				fmt.Printf("%s\n", word)
			} else {
				tkn, err := l.readRune()

				if tkn != nil {
					fmt.Printf("%s\n", tkn)
				}
				
				if err == io.EOF {
					fmt.Printf("END-OF-FILE\n")
					return
				}			
			}
		}
	}
}

func (l * Analyzer) readRune() (* Token, error) {
	startingLine := l.source.state.lineNumber
	startingByte := l.source.state.bytePosition	
	r, size, err := l.source.Read()

	if err != nil {
		return nil, err
	}
	
	if r == '\n' {
		return &Token{kind: newline,
			value: "\n",
			filename: l.filename,
			line: startingLine,
			byte: startingByte, 
			size: size}, nil	
	}

	return nil, nil
}

