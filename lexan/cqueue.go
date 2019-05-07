package lexan

// CharacterQueue holds a slice of Characters.
type CharacterQueue struct {
	s []Character
	head int
	tail int	
}

// Adds a Character to the end of a queue.
func (q * CharacterQueue) Add(char Character){
	q.s = append(q.s, char)
	q.tail += 1
}

// Removes a Character from the top of the queue
func (q * CharacterQueue) Remove() *Character {
	if len(q.s) <= q.head || q.Size() == 0 {
		return nil
	}
	
	char := q.s[q.head]
	q.head += 1

	if q.head > q.tail {
		q.s = []Character{}
		q.head = 0
		q.tail = 0
	}
	
	return &char
}

// Prepends a slice of Characters to the stack
func (q * CharacterQueue) Prepend(chars []Character){
	old := q.s[q.head:q.tail]
	new := chars
	
	for _, rep := range old {
		new = append(new, rep)
	}

	q.s = new
	q.head = 0
	q.tail = len(new)
}

// Get's the current size of the RuneStack
func (q * CharacterQueue) Size() int {
	return q.tail - q.head
}
