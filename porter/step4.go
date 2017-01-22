/*Package porter implement Porter, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/
*/
package porter

import "strings"

/*Step4 from "An algorithm for suffix stripping".
  From the paper:

        Step 4
            (m>1) AL    ->                  revival        ->  reviv
            (m>1) ANCE  ->                  allowance      ->  allow
            (m>1) ENCE  ->                  inference      ->  infer
            (m>1) ER    ->                  airliner       ->  airlin
            (m>1) IC    ->                  gyroscopic     ->  gyroscop
            (m>1) ABLE  ->                  adjustable     ->  adjust
            (m>1) IBLE  ->                  defensible     ->  defens
            (m>1) ANT   ->                  irritant       ->  irrit
            (m>1) EMENT ->                  replacement    ->  replac
            (m>1) MENT  ->                  adjustment     ->  adjust
            (m>1) ENT   ->                  dependent      ->  depend
            (m>1 and (*S or *T)) ION ->     adoption       ->  adopt
            (m>1) OU    ->                  homologou      ->  homolog
            (m>1) ISM   ->                  communism      ->  commun
            (m>1) ATE   ->                  activate       ->  activ
            (m>1) ITI   ->                  angulariti     ->  angular
            (m>1) OUS   ->                  homologous     ->  homolog
            (m>1) IVE   ->                  effective      ->  effect
            (m>1) IZE   ->                  bowdlerize     ->  bowdler
        The suffixes are now removed. All that remains is a little
        tidying up.
*/
func Step4(word string) string {
	// For AL suffix. AL ->
	al := strings.HasSuffix(word, "al")
	if al {
		pre := strings.TrimSuffix(word, "al")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ANCE suffix. ANCE ->
	ance := strings.HasSuffix(word, "ance")
	if ance {
		pre := strings.TrimSuffix(word, "ance")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ENCE suffix. ENCE ->
	ence := strings.HasSuffix(word, "ence")
	if ence {
		pre := strings.TrimSuffix(word, "ence")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ER suffix. ER ->
	er := strings.HasSuffix(word, "er")
	if er {
		pre := strings.TrimSuffix(word, "er")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For IC suffix. IC ->
	ic := strings.HasSuffix(word, "ic")
	if ic {
		pre := strings.TrimSuffix(word, "ic")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ABLE suffix. ABLE ->
	able := strings.HasSuffix(word, "able")
	if able {
		pre := strings.TrimSuffix(word, "able")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For IBLE suffix. IBLE ->
	ible := strings.HasSuffix(word, "ible")
	if ible {
		pre := strings.TrimSuffix(word, "ible")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ANT suffix. ANT ->
	ant := strings.HasSuffix(word, "ant")
	if ant {
		pre := strings.TrimSuffix(word, "ant")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For EMENT suffix. EMENT ->
	ement := strings.HasSuffix(word, "ement")
	if ement {
		pre := strings.TrimSuffix(word, "ement")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For MENT suffix. MENT ->
	ment := strings.HasSuffix(word, "ment")
	if ment {
		pre := strings.TrimSuffix(word, "ment")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ENT suffix. ENT ->
	ent := strings.HasSuffix(word, "ent")
	if ent {
		pre := strings.TrimSuffix(word, "ent")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// (m>1 and (*S or *T)) ION ->
	ion := strings.HasSuffix(word, "ion")
	if ion {
		pre := strings.TrimSuffix(word, "ion")
		if MeasureGreaterThan1(pre) && HasEndst(pre) {
			return pre
		}
	}

	// For OU suffix. OU ->
	ou := strings.HasSuffix(word, "ou")
	if ou {
		pre := strings.TrimSuffix(word, "ou")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ISM suffix. ISM ->
	ism := strings.HasSuffix(word, "ism")
	if ism {
		pre := strings.TrimSuffix(word, "ism")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ATE suffix. ATE ->
	ate := strings.HasSuffix(word, "ate")
	if ate {
		pre := strings.TrimSuffix(word, "ate")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For ITI suffix. ITI ->
	iti := strings.HasSuffix(word, "iti")
	if iti {
		pre := strings.TrimSuffix(word, "iti")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For OUS suffix. OUS ->
	ous := strings.HasSuffix(word, "ous")
	if ous {
		pre := strings.TrimSuffix(word, "ous")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For IVE suffix. IVE ->
	ive := strings.HasSuffix(word, "ive")
	if ive {
		pre := strings.TrimSuffix(word, "ive")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}

	// For IZE suffix. IZE ->
	ize := strings.HasSuffix(word, "ize")
	if ize {
		pre := strings.TrimSuffix(word, "ize")
		if MeasureGreaterThan1(pre) {
			return pre
		}
	}
	return word
}
