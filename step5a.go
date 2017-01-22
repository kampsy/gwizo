package gwizo

import "strings"

/*Step5a from "An algorithm for suffix stripping".

  From the paper:

        Step 5a
            (m>1) E     ->                  probate        ->  probat
                                            rate           ->  rate
            (m=1 and not *o) E ->           cease          ->  ceas
*/
func Step5a(word string) string {
	// E suffix. E ->
	e := strings.HasSuffix(word, letterE)
	if e {
		pre := strings.TrimSuffix(word, letterE)
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// (m=1 and not *o) E ->
	pre := strings.TrimSuffix(word, letterE)
	if MeasureEqualTo1(pre) && !HascvcEndLastNotwxy(pre) {
		return pre
	}
	return word
}
