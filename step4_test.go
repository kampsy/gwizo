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

func TestStep4(t *testing.T) {
	input := []string{
		"revival", "allowance", "inference", "airliner", "gyroscopic", "adjustable",
		"defensible", "irritant", "replacement", "adjustment", "dependent",
		"adoption", "homologou", "communism", "activate", "angulariti", "homologous",
		"effective", "bowdlerize",
	}

	stem := []string{
		"reviv", "allow", "infer", "airlin", "gyroscop", "adjust", "defens",
		"irrit", "replac", "adjust", "depend", "adopt", "homolog", "commun", "activ",
		"angular", "homolog", "effect", "bowdler",
	}

	for i := 0; i < len(input); i++ {
		token := gwizo.Step4(input[i])
		if token != stem[i] {
			t.Errorf(fmt.Sprintf("Test For %s -FAIL- [%s != %s]", input[i], token, stem[i]))
		} else {
			t.Log(fmt.Sprintf("Test For %s *PASS* [%s == %s]", input[i], token, stem[i]))
		}
	}
}
