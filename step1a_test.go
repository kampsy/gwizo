/*Package gwizo implement Porter, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/
*/
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
		token := gwizo.Step1a(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s FAIL [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s -PASS- [%s == %s]", input[i], token, stem[i]))
		}
	}
}
