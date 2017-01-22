package gwizo

import (
	"fmt"
	"testing"

	"github.com/kampsy/gwizo"
)

func TestStep5b(t *testing.T) {
	input := []string{
		"controll", "roll",
	}

	stem := []string{
		"control", "roll",
	}

	for i := 0; i < len(input); i++ {
		token := gwizo.Step5b(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAIL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s *PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
