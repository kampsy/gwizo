/*Package porter implement Porter, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/
*/
package porter

import (
	"fmt"
	"testing"

	"github.com/kampsy/gwizo/porter"
)

func TestStep1b(t *testing.T) {
	input := []string{
		"feed", "agreed", "plastered", "bled", "motoring", "sing", "conflated",
		"troubled", "sized", "hopping", "tanned", "falling", "hissing", "fizzed",
		"failing", "filing",
	}

	stem := []string{
		"feed", "agree", "plaster", "bled", "motor", "sing", "conflate", "trouble",
		"size", "hop", "tan", "fall", "hiss", "fizz", "fail", "file",
	}

	for i := 0; i < len(input); i++ {
		token := porter.Step1b(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAIL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s *PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
