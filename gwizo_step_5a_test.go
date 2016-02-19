package gwizo

import (
  "testing"
  "fmt"
  "github.com/kampsy/gwizo"
)

func TestStep_5a(t *testing.T) {
  input := []string {
    "probate", "rate", "cease",
  }

  stem := []string {
    "probat", "rate", "ceas",
  }

  for i := 0; i < len(input); i++ {
    analyse := gwizo.Ingest(input[i])
    if analyse.Step_5a() != stem[i] {
      t.Error(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], analyse.Step_5a(), stem[i]))
    }else {
      t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], analyse.Step_5a(), stem[i]))
    }
  }
}