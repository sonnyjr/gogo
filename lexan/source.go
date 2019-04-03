package lexan

import (
	"bufio"
)

type SourceState struct {
	lineNumber int
	totalBytes int
	bytePosition int
	
	previous * SourceState
	next * SourceState
}

type Source struct {
	data * bufio.Reader
	filename string
	state * SourceState
	queue RuneQueue
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
			s.revertState(len(runes))			
			s.queue.Prepend(runes)
			return runes, err
		}
		
		totalSize += size
		rep := RuneRep{r: r, size: size}
		runes = append(runes, rep)
	}

	s.revertState(len(runes))
	s.queue.Prepend(runes)
	return runes, nil
}

func (s * Source) Read() (rune, int, error) {	
	rep := s.queue.Remove()
	
	if rep != nil {
		s.updateState((*rep))
		return (*rep).r, (*rep).size, nil
	}
	
	r, size, err := s.data.ReadRune()
	s.updateState(RuneRep{r: r, size: size})	
	
	return r, size, err
}

func (s * Source) revertState(n int){
	for i := 0; i < n; i++ {
		previous := s.state.previous
		s.state = previous
		previous.next = nil
	}
}

func (s * Source) updateState(rep RuneRep){
	newLine := s.state.lineNumber
	newBytes := s.state.bytePosition
	newTotalBytes := s.state.totalBytes + rep.size
	
	if rep.r == '\n' {
		newLine += 1
		newBytes = 0
	} else {
		newBytes += rep.size
	}

	newState := &SourceState{lineNumber: newLine,
		bytePosition: newBytes,
		totalBytes: newTotalBytes,
		previous: s.state}

	s.state.next = newState

	s.state = newState
}
