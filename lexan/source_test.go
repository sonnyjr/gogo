package lexan

import (
	"testing"
	"bufio"
	"strings"
)

type MatchTest struct {
	match string
	value string
	success bool
}

func TestMatch(t *testing.T){
	matches := []MatchTest{
		MatchTest{match: "//", value: "//hello", success: true},
		MatchTest{match: "//", value: "/ /hello", success: false},
		MatchTest{match: "/h", value: "//hello", success: false},
		MatchTest{match: "smile", value: "smile", success: true},
	}

	for _, m := range matches {
		reader := bufio.NewReader(strings.NewReader(m.value))
		analyzer := New(reader, "stdio")
		result := analyzer.source.Match(m.match)
		
		if result != m.success {
			t.Errorf("Match failed when trying to match %s with %s", m.match, m.value)
		}
	}
}

func TestPeek(t *testing.T){
	
}

func TestRead(t *testing.T){
	
}
