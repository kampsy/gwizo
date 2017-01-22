package gwizo

import "strings"

/*Step3 from "An algorithm for suffix stripping".

  From the paper:

        Step 3
            (m>0) ICATE ->  IC              triplicate     ->  triplic
            (m>0) ATIVE ->                  formative      ->  form
            (m>0) ALIZE ->  AL              formalize      ->  formal
            (m>0) ICITI ->  IC              electriciti    ->  electric
            (m>0) ICAL  ->  IC              electrical     ->  electric
            (m>0) FUL   ->                  hopeful        ->  hope
            (m>0) NESS  ->                  goodness       ->  good
*/
func Step3(word string) string {
	// For ICATE suffix. ICATE -> IC
	icate := strings.HasSuffix(word, "icate")
	if icate {
		pre := strings.TrimSuffix(word, "icate")
		if MeasureGreaterThan0(pre) {
			word = pre + "ic"
			return word
		}
	}

	// For ATIVE suffix. ATIVE ->
	ative := strings.HasSuffix(word, "ative")
	if ative {
		pre := strings.TrimSuffix(word, "ative")
		if MeasureGreaterThan0(pre) {
			word = pre
			return word
		}
	}

	// For ALIZE suffix. ALIZE -> AL
	alize := strings.HasSuffix(word, "alize")
	if alize {
		pre := strings.TrimSuffix(word, "alize")
		if MeasureGreaterThan0(pre) {
			word = pre + "al"
			return word
		}
	}

	// For ICITI suffix. ICITI -> IC
	iciti := strings.HasSuffix(word, "iciti")
	if iciti {
		pre := strings.TrimSuffix(word, "iciti")
		if MeasureGreaterThan0(pre) {
			word = pre + "ic"
			return word
		}
	}

	// For ICAL suffix. ICAL -> IC
	ical := strings.HasSuffix(word, "ical")
	if ical {
		pre := strings.TrimSuffix(word, "ical")
		if MeasureGreaterThan0(pre) {
			word = pre + "ic"
			return word
		}
	}

	// For FUL suffix. FUL ->
	ful := strings.HasSuffix(word, "ful")
	if ful {
		pre := strings.TrimSuffix(word, "ful")
		if MeasureGreaterThan0(pre) {
			word = pre
			return word
		}
	}

	// For NESS suffix. NESS ->
	ness := strings.HasSuffix(word, "ness")
	if ness {
		pre := strings.TrimSuffix(word, "ness")
		if MeasureGreaterThan0(pre) {
			word = pre
			return word
		}
	}
	return word
}
