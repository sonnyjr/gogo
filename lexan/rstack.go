package lexan

type RuneRep struct {
	r rune	
	size int
}

type RuneStack struct {
	s []RuneRep
	head int
	tail int	
}

func (r * RuneStack) Push(rep RuneRep){
	r.s = append(r.s, rep)
	r.tail += 1
}

func (r * RuneStack) Pop() *RuneRep {
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

func (r * RuneStack) Prepend(runes []RuneRep){
	old := r.s[r.head:r.tail]
	new := runes
	
	for _, rep := range old {
		new = append(new, rep)
	}

	r.s = new
	r.head = 0
	r.tail = len(new) - 1
}

func (r * RuneStack) Size() int {
	return r.tail - r.head
}
