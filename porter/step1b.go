/*Package porter implement Porter, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/
*/
package porter

import "strings"

/*Step1b from "An algorithm for suffix stripping".
  Deals with plurals and past participles. The subsequent steps
  are much more straightforward.

  From the paper:

      (m>0) EED -> EE                    feed      ->  feed
                                         agreed    ->  agree
      (*v*) ED  ->                       plastered ->  plaster
                                         bled      ->  bled
      (*v*) ING ->                       motoring  ->  motor
                                         sing      ->  sing

  If the second or third of the rules in Step 1b is successful,
  the following is done:
      AT -> ATE                       conflat(ed)  ->  conflate
      BL -> BLE                       troubl(ed)   ->  trouble
      IZ -> IZE                       siz(ed)      ->  size
      (*d and not (*L or *S or *Z))
         -> single letter
                                      hopp(ing)    ->  hop
                                      tann(ed)     ->  tan
                                      fall(ing)    ->  fall
                                      hiss(ing)    ->  hiss
                                      fizz(ed)     ->  fizz
      (m=1 and *o) -> E               fail(ing)    ->  fail
                                      fil(ing)     ->  file
  The rule to map to a single letter causes the removal of one of
  the double letter pair. The -E is put back on -AT, -BL and -IZ,
  so that the suffixes -ATE, -BLE and -IZE can be recognised
  later. This E may be removed in step 4.
*/
func Step1b(word string) string {
	// Word Measure (m > 0) and EED suffix. EED -> EE
	eed := strings.HasSuffix(word, "eed")
	if eed {
		pre := strings.TrimSuffix(word, "eed")
		if MeasureNum(pre) > 0 {
			str := pre + "ee"
			return str
		}
		return word
	}

	// Word has Vowel and ED suffix. ED ->
	ed := strings.HasSuffix(word, "ed")
	if ed {
		pre := strings.TrimSuffix(word, "ed")
		if !HasVowel(pre) {
			return word
		}
		word = pre
	}

	// Word has Vowel and ING suffix. ING ->
	ing := strings.HasSuffix(word, "ing")
	if ing {
		pre := strings.TrimSuffix(word, "ing")
		if !HasVowel(pre) {
			return word
		}
		word = pre
	}

	/*If the second or third of the rules in Step 1b is successful,
	the following is done
	*/
	if ed || ing {
		// Word has AT suffix. AT -> ATE
		at := strings.HasSuffix(word, "at")
		if at {
			pre := strings.TrimSuffix(word, "at")
			word = pre + "ate"
		}

		// Word has BL suffix. BL -> BLE
		bl := strings.HasSuffix(word, "bl")
		if bl {
			pre := strings.TrimSuffix(word, "bl")
			word = pre + "ble"
		}

		// Word has IZ suffix. IZ -> IZE
		iz := strings.HasSuffix(word, "iz")
		if iz {
			pre := strings.TrimSuffix(word, "iz")
			word = pre + "ize"
		}

		// (*d and not (*L or *S or *Z)) -> single letter
		if HasSameDoubleConsonant(word) {
			ll := strings.HasSuffix(word, "ll")
			ss := strings.HasSuffix(word, "ss")
			zz := strings.HasSuffix(word, "zz")
			if ll || ss || zz {
				return word
			}
			wordLen := len(word)
			lastLetter := word[(wordLen - 1):]
			pre := strings.TrimSuffix(word, lastLetter)
			return pre
		}
		// (m=1 and *o) -> E
		if MeasureEqualTo1(word) && HascvcEndLastNotwxy(word) {
			word = word + letterE
		}
	}
	return word
}
