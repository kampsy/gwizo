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

func TestStep1c(t *testing.T) {
	input := []string{
		"happy", "sky",
	}

	stem := []string{
		"happi", "sky",
	}

	for i := 0; i < len(input); i++ {
		token := porter.Step1c(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAiL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s **PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
