package lexan

import "testing"

func TestCharacterNewline(t *testing.T){
	c := NewCharacter('\n', 1)

	if c.char != '\n' {
		t.Errorf("New Character is invalid. Should be newline. %c", c.char)
	}

	if c.class != newline {
		t.Errorf("Invalid class: <newline> != %s", c.class)
	}

	if c.kind != none {
		t.Errorf("Invalid kind: <none> != %s", c.kind)
	}	
}

func TestCharacterChar(t *testing.T){
	runes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
	
	for _, r := range runes {
		c := NewCharacter(r, 1)

		if c.char != r {
			t.Errorf("New Character is invalid. Should be %c not %c", r, c.char)
		}

		
		if c.class != unicode_char {
			t.Errorf("Invalid class: <unicode_char> != %s",  c.class)
		}

		if (c.kind & letter) == 0 {
			t.Errorf("Invalid kind: <none> != %s", c.kind)
		}
	}
}

