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
		octopus := gwizo.Ingest(input[i])
		if octopus.Step5a() != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], octopus.Step5a(), stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], octopus.Step5a(), stem[i]))
		}
	}
}
