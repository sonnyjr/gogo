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
	queue CharacterQueue
}

// Checks if the value is the next set of runes.
func (s * Source) Match(value string) bool {
	var size = len(value)
	reps, err := s.Peek(size)
	
	if err != nil || size != len(reps) {
		return false
	}

	for i, r := range value {
		if r != reps[i].char {
			return false
		}
	}

	return true
}

// Peek's n runes ahead without changing the state
func (s * Source) Peek(n int) ([]Character, error) {
	runes := []Character{}
	totalSize := 0
	
	for i := 0; i < n; i++ {
		r, size, err := s.read(false)
			
		if err != nil {
			s.queue.Prepend(runes)
			return runes, err
		}
		
		totalSize += size
		rep := NewCharacter(r, size)
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
		
		return (*rep).char, (*rep).size, nil
	}
	
	r, size, err := s.data.ReadRune()

	if update {
		s.updateState(NewCharacter(r,size))
	}
	
	return r, size, err	
}

// Updates the state based on the rune that was read.
func (s * Source) updateState(c Character){
	s.state.totalBytes += c.size
	
	if c.char == '\n' {		
		s.state.lineNumber += 1
		s.state.bytePosition = 0
	} else {
		s.state.bytePosition += c.size
	}
}
