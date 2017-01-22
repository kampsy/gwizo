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

func TestStep3(t *testing.T) {
	input := []string{
		"triplicate", "formative", "formalize", "electriciti", "electrical",
		"hopeful", "goodness",
	}

	stem := []string{
		"triplic", "form", "formal", "electric", "electric", "hope", "good",
	}

	for i := 0; i < len(input); i++ {
		token := gwizo.Step3(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAIL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s *PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
