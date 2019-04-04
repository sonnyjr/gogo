package lexan

import (
	"bufio"
)

// The current read state of the Source file.
type SourceState struct {
	lineNumber int
	totalBytes int
	bytePosition int
}

// The source type represents the current file
// and read status.
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

	for i, r := range value {
		if r != reps[i].r {
			return false
		}
	}

	return true
}

// Peek's n runes ahead without changing the state
func (s * Source) Peek(n int) ([]RuneRep, error) {
	runes := []RuneRep{}
	totalSize := 0
	
	for i := 0; i < n; i++ {
		r, size, err := s.read(false)
			
		if err != nil {
			s.queue.Prepend(runes)
			return runes, err
		}
		
		totalSize += size
		rep := RuneRep{r: r, size: size}
		runes = append(runes, rep)
	}

	s.queue.Prepend(runes)
	return runes, nil
}

// Read's a rune and updates the current state
func (s * Source) Read() (rune, int, error) {	
	return s.read(true)
}

// Read's a rune from the source and updates the state
// based on the update value.
func (s * Source) read(update bool) (rune, int, error){
	rep := s.queue.Remove()
	
	if rep != nil {
		if update {
			s.updateState((*rep))
		}
		
		return (*rep).r, (*rep).size, nil
	}
	
	r, size, err := s.data.ReadRune()

	if update {
		s.updateState(RuneRep{r: r, size: size})
	}
	
	return r, size, err	
}

// Updates the state based on the rune that was read.
func (s * Source) updateState(rep RuneRep){
	s.state.totalBytes += rep.size
	
	if rep.r == '\n' {		
		s.state.lineNumber += 1
		s.state.bytePosition = 0
	} else {
		s.state.bytePosition += rep.size
	}
}
