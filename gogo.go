package main

import (
	"bufio"
	"os"
	"log"
	"github.com/sonnyjr/gogo/LexicalAnalyzer"
)

func main(){
	args := os.Args[1:]	
	
	if len(args) > 0 {
		file, err := os.Open(args[0])
		reader := bufio.NewReader(file)
		
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}

		lex := lexicalAnalyzer.New(reader, args[0])
		lex.Analyze()		
	} else {
		reader := bufio.NewReader(os.Stdin)
		lex := lexicalAnalyzer.New(reader, "<STDIN>")
		lex.Analyze()
	}
}
