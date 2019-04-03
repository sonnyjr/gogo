package lexan

import (
	"bufio"
	"fmt"
	"log"
	"io"
)

type Analyzer struct {
	source * Source
	filename string
}


func New(data * bufio.Reader, filename string) Analyzer {
	source := Source{data: data, filename: filename}
	return Analyzer{source: &source, filename: filename}
}

func (l * Analyzer) Analyze(){
	for {

		comment := l.parseComment()

		if comment != nil {
			fmt.Printf("%s\n", comment)
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
	}
}

func (l * Analyzer) readRune() error {
	r, size, err := l.source.Read()

	if err != nil {
		return err
	}
	
	if r == '\n' {			
		fmt.Printf("%s:%d:NEWLINE:%d\n", l.filename, l.source.bytePosition, size)
	} else {
		fmt.Printf("%s:%d:%c:%d\n", l.filename, l.source.bytePosition, r, size)
	}

	return nil
}

