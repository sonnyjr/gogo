// The lexan package is a Lexical Analyzer for the go programming language
package lexan

import (
	"bufio"
	"fmt"
	"io"
)

// Analyzer represents the analysis of the program.
type Analyzer struct {
	source * Source
	filename string
}

// New creates a new Analyzer.
func New(data * bufio.Reader, filename string) Analyzer {
	state := SourceState{}
	source := Source{data: data, filename: filename, state: &state}
	return Analyzer{source: &source, filename: filename}
}

// Analyzer method will start the lexical analysis of the program
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
				_, _, err := l.source.Read()
				
				if err == io.EOF {
					fmt.Printf("END-OF-FILE\n")
					return
				}			
			}
		}
	}
}
