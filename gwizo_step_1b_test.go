package gwizo

import (
  "testing"
  "fmt"
  "github.com/kampsy/gwizo"
)

func TestStep_1b(t *testing.T) {
  input := []string {
    "feed", "agreed", "plastered", "bled", "motoring", "sing", "conflated",
    "troubled", "sized", "hopping", "tanned", "falling", "hissing", "fizzed",
    "failing", "filing",
  }

  stem := []string {
    "feed","agree", "plaster", "bled", "motor", "sing", "conflate", "trouble",
    "size", "hop", "tan", "fall", "hiss", "fizz", "fail", "file",
  }

  for i := 0; i < len(input); i++ {
    octopus := gwizo.Ingest(input[i])
    if octopus.Step_1b() != stem[i] {
      t.Error(fmt.Sprintf("Test For %s ~~Failed~~ [%s != %s]", input[i], octopus.Step_1b(), stem[i]))
    }else {
      t.Log(fmt.Sprintf("Test For %s **Passed** [%s == %s]", input[i], octopus.Step_1b(), stem[i]))
    }
  }
}
