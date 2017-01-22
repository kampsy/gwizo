package gwizo

import (
	"strings"
)

/*Step1c from "An algorithm for suffix stripping".
  Deals with plurals and past participles. The subsequent steps
  are much more straightforward.

  From the paper:

        Step 1c
            (*v*) Y -> I                    happy        ->  happi
                                            sky          ->
*/
func Step1c(word string) string {
	// (*v*) Y -> I
	// Word has Vowel and Y suffix. Y -> I
	y := strings.HasSuffix(word, letterY)
	if y {
		pre := strings.TrimSuffix(word, letterY)
		if HasVowel(pre) {
			return pre + letterI
		}
	}
	return word
}
