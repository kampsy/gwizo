package gwizo

import (
	"fmt"
	"testing"

	"github.com/kampsy/gwizo"
)

func TestStep1b(t *testing.T) {
	input := []string{
		"feed", "agreed", "plastered", "bled", "motoring", "sing", "conflated",
		"troubled", "sized", "hopping", "tanned", "falling", "hissing", "fizzed",
		"failing", "filing",
	}

	stem := []string{
		"feed", "agree", "plaster", "bled", "motor", "sing", "conflate", "trouble",
		"size", "hop", "tan", "fall", "hiss", "fizz", "fail", "file",
	}

	for i := 0; i < len(input); i++ {
		token := gwizo.Step1b(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAIL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s *PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
