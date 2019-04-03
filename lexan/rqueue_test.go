package lexan

import "testing"

// Tests adding runes to the RuneQueue.
func TestAdd(t *testing.T) {
	queue := RuneQueue{}
	runes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
	
	for i, r := range runes {
		rep := RuneRep{r: r, size: 1}
		queue.Add(rep)

		if queue.head != 0 {
			t.Errorf("Head invalid. Should be 0")
		}

		if queue.tail != (i+1) {
			t.Errorf("Tail invalid. Should be %d", (i+1))
		}
	}
}

// Tests adding and then removing runes from the queue.
func TestRemove(t *testing.T) {
	queue := RuneQueue{}
	runes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
	
	for _, r := range runes {
		rep := RuneRep{r: r, size: 1}
		queue.Add(rep)
	}

	for i, r := range runes {
		top := queue.Remove()
		
		if top.r != r {
			t.Errorf("Runes does not equal.")
		}

		if queue.head != (i+1) {
			t.Errorf("Head(%d) invalid. Should be %d", queue.head, (i+1))
		}

		if queue.tail != len(runes) {
			t.Errorf("Tail invalid. Should be %d", len(runes))
		}
	}
}

// Tests prepending runes to the queue.
func TestPrepend(t *testing.T){
	queue := RuneQueue{}
	runes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
	
	for _, r := range runes {
		rep := []RuneRep{RuneRep{r: r, size: 1}}
		queue.Prepend(rep)		
	}

	if queue.Size() != len(runes) {
		t.Errorf("Queue is incorrect size. Size = %d, head = %d, tail = %d",
			queue.Size(), queue.head, queue.tail)
	}
	
	for i := len(runes); i > 0; i-- {
		rep := queue.Remove()
		
		if rep == nil {
			t.Errorf("rep is nil.")
		}
		
		if rep.r != runes[i-1] {
			t.Errorf("Runes are not equal. %c != %c", rep.r, runes[i])
		}
	}
}
