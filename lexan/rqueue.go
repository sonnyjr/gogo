package lexan

// RuneRep holds a rune and it's size
type RuneRep struct {
	r rune	
	size int
}

// RuneQueue holds a slice of RuneReps.
type RuneQueue struct {
	s []RuneRep
	head int
	tail int	
}

// Adds a Runerep to the end of a queue.
func (r * RuneQueue) Add(rep RuneRep){
	r.s = append(r.s, rep)
	r.tail += 1
}

// Removes a RuneRep from the top of the queue
func (r * RuneQueue) Remove() *RuneRep {
	if len(r.s) <= r.head || r.Size() == 0 {
		return nil
	}
	
	rep := r.s[r.head]
	r.head += 1

	if r.head > r.tail {
		r.s = []RuneRep{}
		r.head = 0
		r.tail = 0
	}
	
	return &rep
}

// Prepends a slice of RuneReps to the stack
func (r * RuneQueue) Prepend(runes []RuneRep){
	old := r.s[r.head:r.tail]
	new := runes
	
	for _, rep := range old {
		new = append(new, rep)
	}

	r.s = new
	r.head = 0
	r.tail = len(new)
}

// Get's the current size of the RuneStack
func (r * RuneQueue) Size() int {
	return r.tail - r.head
}
