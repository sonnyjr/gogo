package main

import (
	"bufio"
	"fmt"
	"log"
	"io"
)

type LexicalAnalyzer struct {
	data * bufio.Reader
	filename string
	
}

func (l *LexicalAnalyzer) Analyze(){
	count := 0
	
	for {
		r, size, err := l.data.ReadRune()

		if err == io.EOF {
			fmt.Printf("END-OF-FILE\n")
			return
		}
		
		if err != nil {
			log.Fatal(err)
			break
		}

		if r == '\n' {
			fmt.Printf("%s:%d:NEWLINE:%d\n", l.filename, count, size)
		} else {
			fmt.Printf("%s:%d:%c:%d\n", l.filename, count, r, size)
		}
		
		count += 1
	}
}
