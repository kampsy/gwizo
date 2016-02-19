package gwizo

import (
  "testing"
  "fmt"
  "github.com/kampsy/gwizo"
)

func TestStep_3(t *testing.T) {
  input := []string {
    "triplicate", "formative", "formalize", "electriciti", "electrical",
    "hopeful", "goodness",
  }

  stem := []string {
    "triplic", "form", "formal", "electric", "electric", "hope", "good",
  }

  for i := 0; i < len(input); i++ {
    analyse := gwizo.Ingest(input[i])
    if analyse.Step_3() != stem[i] {
      t.Error(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], analyse.Step_3(), stem[i]))
    }else {
      t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], analyse.Step_3(), stem[i]))
    }
  }
}