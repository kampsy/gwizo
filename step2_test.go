package gwizo

import (
	"fmt"
	"testing"

	"github.com/kampsy/gwizo"
)

func TestStep2(t *testing.T) {
	input := []string{
		"relational", "conditional", "rational", "valenci", "hesitanci", "digitizer",
		"conformabli", "radicalli", "differentli", "vileli", "analogousli",
		"vietnamization", "predication", "operator", "feudalism", "decisiveness",
		"hopefulness", "callousness", "formaliti", "sensitiviti", "sensibiliti",
	}

	stem := []string{
		"relate", "condition", "rational", "valence", "hesitance", "digitize",
		"conformable", "radical", "different", "vile", "analogous", "vietnamize",
		"predicate", "operate", "feudal", "decisive", "hopeful", "callous",
		"formal", "sensitive", "sensible",
	}

	for i := 0; i < len(input); i++ {
		token := gwizo.Step2(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAIL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s *PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
