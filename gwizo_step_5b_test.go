package gwizo

import (
  "testing"
  "fmt"
  "github.com/kampsy/gwizo"
)

func TestStep_5b(t *testing.T) {
  input := []string {
    "controll", "roll",
  }

  stem := []string {
    "control", "roll",
  }

  for i := 0; i < len(input); i++ {
    octopus := gwizo.Ingest(input[i])
    if octopus.Step_5b() != stem[i] {
      t.Error(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], octopus.Step_5b(), stem[i]))
    }else {
      t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], octopus.Step_5b(), stem[i]))
    }
  }
}
