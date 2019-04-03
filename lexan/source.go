package lexan

import (
	"bufio"
)

type Source struct {
	data * bufio.Reader
	filename string
	bytePosition int
	stack RuneStack
}

// Checks if the value is the next set of runes.
func (s * Source) Match(value string) bool {
	var size = len(value)
	reps, err := s.Peek(size)
	
	if err != nil || size != len(reps) {
		return false
	}

	runes := []rune{}
	for _, rep := range reps {
		runes = append(runes, rep.r)
	}
	
	s.stack.Prepend(reps)
	
	if string(runes) != value {
		return false
	}

	return true
}

// Peek's n runes ahead.
func (s * Source) Peek(n int) ([]RuneRep, error) {
	runes := []RuneRep{}
	totalSize := 0
	
	for i := 0; i < n; i++ {
		r, size, err := s.Read()			
			
		if err != nil {
			s.stack.Prepend(runes)
			return runes, err
		}
		
		totalSize += size
		rep := RuneRep{r: r, size: size}
		runes = append(runes, rep)
	}

	s.bytePosition -= totalSize
	s.stack.Prepend(runes)
	return runes, nil
}

func (s * Source) Read() (rune, int, error) {	
	rep := s.stack.Pop()
	
	if rep != nil {	
		s.bytePosition += (*rep).size		
		return (*rep).r, (*rep).size, nil
	}

	r, size, err := s.data.ReadRune()
	s.bytePosition += size
	
	return r, size, err
}
