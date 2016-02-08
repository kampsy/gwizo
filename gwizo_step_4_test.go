package gwizo

import (
  "testing"
  "fmt"
  "github.com/kampsy/gwizo"
)

func TestStep_4(t *testing.T) {
  input := []string {
    "revival", "allowance", "inference", "airliner", "gyroscopic", "adjustable",
    "defensible", "irritant", "replacement", "adjustment", "dependent",
    "adoption", "homologou", "communism", "activate", "angulariti", "homologous",
    "effective", "bowdlerize",
  }

  stem := []string {
    "reviv", "allow", "infer", "airlin", "gyroscop", "adjust", "defens",
    "irrit", "replac", "adjust", "depend", "adopt", "homolog", "commun", "activ",
    "angular", "homolog", "effect", "bowdler",
  }

  for i := 0; i < len(input); i++ {
    analyse := gwizo.Ingest(input[i])
    if analyse.Step_4() != stem[i] {
      t.Error(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], analyse.Step_4(), stem[i]))
    }else {
      t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], analyse.Step_4(), stem[i]))
    }
  }
}
