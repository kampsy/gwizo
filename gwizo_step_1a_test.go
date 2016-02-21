package gwizo

import (
  "testing"
  "fmt"
  "github.com/kampsy/gwizo"
)

func TestStep_1a(t *testing.T) {
  input := []string {
    "caresses", "ponies", "ties", "caress", "cats",
  }

  stem := []string {
    "caress", "poni", "ti", "caress", "cat",
  }

  for i := 0; i < len(input); i++ {
    octopus := gwizo.Ingest(input[i])
    if octopus.Step_1a() != stem[i] {
      t.Error(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], octopus.Step_1a(), stem[i]))
    }else {
      t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], octopus.Step_1a(), stem[i]))
    }
  }
}
