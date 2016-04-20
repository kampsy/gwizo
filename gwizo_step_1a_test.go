package gwizo

import (
	"fmt"
	"testing"

	"github.com/kampsy/gwizo"
)

func TestStep1a(t *testing.T) {
	input := []string{
		"caresses", "ponies", "ties", "caress", "cats",
	}

	stem := []string{
		"caress", "poni", "ti", "caress", "cat",
	}

	for i := 0; i < len(input); i++ {
		octopus := gwizo.Ingest(input[i])
		if octopus.Step1a() != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], octopus.Step1a(), stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], octopus.Step1a(), stem[i]))
		}
	}
}
