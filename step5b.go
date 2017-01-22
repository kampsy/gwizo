package gwizo

import "strings"

/*Step5b from "An algorithm for suffix stripping".

  From the paper:

        Step 5b
            (m > 1 and *d and *L) -> single letter
                                    controll       ->  control
                                    roll           ->  roll
*/
func Step5b(word string) string {
	// (m > 1 and *d and *L) -> single letter
	if MeasureGreaterThan1(word) && HasSameDoubleConsonant(word) && HasEndl(word) {
		wordLen := len(word)
		lastLetter := word[(wordLen - 1):]
		pre := strings.TrimSuffix(word, lastLetter)
		return pre
	}
	return word
}
