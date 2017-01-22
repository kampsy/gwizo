package gwizo

import (
	"fmt"
	"testing"

	"github.com/kampsy/gwizo"
)

func TestStep5a(t *testing.T) {
	input := []string{
		"probate", "rate", "cease",
	}

	stem := []string{
		"probat", "rate", "ceas",
	}

	for i := 0; i < len(input); i++ {
		token := gwizo.Step5a(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAIL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s *PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
