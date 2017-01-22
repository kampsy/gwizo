/*Package porter implement Porter, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/
*/
package porter

import (
	"strings"
	"unicode"
)

/*Stem from "An algorithm for suffix stripping".
 */
func Stem(word string) string {
	if len(word) <= 2 {
		return word
	}
	var str string
	for _, char := range word {
		if unicode.IsLetter(rune(char)) {
			str = str + string(char)
		}
	}
	if strings.TrimSpace(str) == "" {
		return word
	}
	word = Step1a(str)
	word = Step1b(word)
	word = Step1c(word)
	word = Step2(word)
	word = Step3(word)
	word = Step4(word)
	word = Step5a(word)
	word = Step5b(word)
	return word
}
