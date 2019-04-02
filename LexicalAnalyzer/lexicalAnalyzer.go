package lexicalAnalyzer

import (
	"bufio"
	"fmt"
	"log"
	"io"
)

type Analyzer struct {
	data * bufio.Reader
	filename string
	currByte uint16
}


func New(data * bufio.Reader, filename string) Analyzer {
	return Analyzer{data: data, filename: filename}
}

func (l * Analyzer) Analyze(){
	for {
		byte, error := l.data.Peek(1)

		if error != nil {
			log.Fatal(error)
			break
		}

		if byte[0] == '/' {
			token, err := l.parseComment()

			if err != nil {
				log.Fatal(err)
			}
			
			if token == nil {
				l.readRune()
			}
			
			fmt.Printf("%s\n", token)
		} else {
			err := l.readRune()

			if err == io.EOF {
				fmt.Printf("END-OF-FILE\n")
				return
			}
			
			if err != nil {
				log.Fatal(err)
			}
		}

		l.currByte += 1
	}
}

func (l * Analyzer) readRune() error {
	r, size, err := l.data.ReadRune()
	count := 0
	
	if err != nil {
		return err
	}
	
	if r == '\n' {			
		fmt.Printf("%s:%d:NEWLINE:%d\n", l.filename, count, size)
	} else {
		fmt.Printf("%s:%d:%c:%d\n", l.filename, count, r, size)
	}

	return nil
}

func (l * Analyzer) parseComment() (*Token, error) {
	bytes, err := l.data.Peek(2)

	if err != nil {
		return nil, err
	}

	if bytes[0] == '/' && bytes[1] == '/' {
		var size uint16
		value := make([]byte, 2)
		
		for {
			byte, err := l.data.ReadByte()
									
			if err == io.EOF || byte == '\n' {
				break
			}

			value = append(value, byte)
		}

		return &Token{kind: comment,
			value: string(value),
			filename: l.filename,
			byte: l.currByte,
			size: size}, nil;
	}

	return nil, nil
}
