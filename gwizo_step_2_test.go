package gwizo

import (
  "testing"
  "fmt"
  "github.com/kampsy/gwizo"
)

func TestStep_2(t *testing.T) {
  input := []string {
    "relational", "conditional", "rational", "valenci", "hesitanci", "digitizer",
    "conformabli", "radicalli", "differentli", "vileli", "analogousli",
    "vietnamization", "predication", "operator", "feudalism", "decisiveness",
    "hopefulness", "callousness", "formaliti", "sensitiviti", "sensibiliti",
  }

  stem := []string {
    "relate", "condition", "rational", "valence", "hesitance", "digitize",
    "conformable", "radical", "different", "vile", "analogous", "vietnamize",
    "predicate", "operate", "feudal", "decisive", "hopeful", "callous",
    "formal", "sensitive", "sensible",
  }

  for i := 0; i < len(input); i++ {
    octopus := gwizo.Ingest(input[i])
    if octopus.Step_2() != stem[i] {
      t.Error(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], octopus.Step_2(), stem[i]))
    }else {
      t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], octopus.Step_2(), stem[i]))
    }
  }
}
