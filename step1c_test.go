package gwizo

import (
	"fmt"
	"testing"

	"github.com/kampsy/gwizo"
)

func TestStep1c(t *testing.T) {
	input := []string{
		"happy", "sky",
	}

	stem := []string{
		"happi", "sky",
	}

	for i := 0; i < len(input); i++ {
		token := gwizo.Step1c(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAiL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s **PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
