/*Package porter implement Porter, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/
*/
package porter

import "strings"

/*Step2 from "An algorithm for suffix stripping"

From the paper:

        Step 2
            (m>0) ATIONAL ->  ATE       relational     ->  relate
            (m>0) TIONAL  ->  TION      conditional    ->  condition
                                        rational       ->  rational
            (m>0) ENCI    ->  ENCE      valenci        ->  valence
            (m>0) ANCI    ->  ANCE      hesitanci      ->  hesitance
            (m>0) IZER    ->  IZE       digitizer      ->  digitize
            (m>0) ABLI    ->  ABLE      conformabli    ->  conformable
            (m>0) ALLI    ->  AL        radicalli      ->  radical
            (m>0) ENTLI   ->  ENT       differentli    ->  different
            (m>0) ELI     ->  E         vileli        - >  vile
            (m>0) OUSLI   ->  OUS       analogousli    ->  analogous
            (m>0) IZATION ->  IZE       vietnamization ->  vietnamize
            (m>0) ATION   ->  ATE       predication    ->  predicate
            (m>0) ATOR    ->  ATE       operator       ->  operate
            (m>0) ALISM   ->  AL        feudalism      ->  feudal
            (m>0) IVENESS ->  IVE       decisiveness   ->  decisive
            (m>0) FULNESS ->  FUL       hopefulness    ->  hopeful
            (m>0) OUSNESS ->  OUS       callousness    ->  callous
            (m>0) ALITI   ->  AL        formaliti      ->  formal
            (m>0) IVITI   ->  IVE       sensitiviti    ->  sensitive
            (m>0) BILITI  ->  BLE       sensibiliti    ->  sensible
*/
func Step2(word string) string {
	// For ATIONAL suffix. ATIONAL -> ATE
	ational := strings.HasSuffix(word, "ational")
	if ational {
		pre := strings.TrimSuffix(word, "ational")
		if MeasureGreaterThan0(pre) {
			word = pre + "ate"
			return word
		}

	}

	// For TIONAL suffix. TIONAL -> TION
	tional := strings.HasSuffix(word, "tional")
	if tional {
		pre := strings.TrimSuffix(word, "tional")
		if MeasureGreaterThan0(pre) {
			word = pre + "tion"
			return word
		}

	}

	// For ENCI suffix. ENCI -> ENCE
	enci := strings.HasSuffix(word, "enci")
	if enci {
		pre := strings.TrimSuffix(word, "enci")
		if MeasureGreaterThan0(pre) {
			word = pre + "ence"
			return word
		}

	}

	// For ANCI suffix. ANCI -> ANCE
	anci := strings.HasSuffix(word, "anci")
	if anci {
		pre := strings.TrimSuffix(word, "anci")
		if MeasureGreaterThan0(pre) {
			word = pre + "ance"
			return word
		}

	}

	// For IZER suffix. IZER -> IZE
	izer := strings.HasSuffix(word, "izer")
	if izer {
		pre := strings.TrimSuffix(word, "izer")
		if MeasureGreaterThan0(pre) {
			word = pre + "ize"
			return word
		}

	}

	// For ABLI suffix. ABLI -> ABLE
	abli := strings.HasSuffix(word, "abli")
	if abli {
		pre := strings.TrimSuffix(word, "abli")
		if MeasureGreaterThan0(pre) {
			word = pre + "able"
			return word
		}

	}

	// For ALLI suffix. ALLI -> AL
	alli := strings.HasSuffix(word, "alli")
	if alli {
		pre := strings.TrimSuffix(word, "alli")
		if MeasureGreaterThan0(pre) {
			word = pre + "al"
			return word
		}

	}

	// For ENTLI suffix. ENTLI -> ENT
	entli := strings.HasSuffix(word, "entli")
	if entli {
		pre := strings.TrimSuffix(word, "entli")
		if MeasureGreaterThan0(pre) {
			word = pre + "ent"
			return word
		}

	}

	// For ELI suffix. ELI -> E
	eli := strings.HasSuffix(word, "eli")
	if eli {
		pre := strings.TrimSuffix(word, "eli")
		if MeasureGreaterThan0(pre) {
			word = pre + "e"
			return word
		}

	}

	// For OUSLI suffix. OUSLI -> OUS
	ousli := strings.HasSuffix(word, "ousli")
	if ousli {
		pre := strings.TrimSuffix(word, "ousli")
		if MeasureGreaterThan0(pre) {
			word = pre + "ous"
			return word
		}

	}

	// For IZATION suffix. IZATION -> IZE
	ization := strings.HasSuffix(word, "ization")
	if ization {
		pre := strings.TrimSuffix(word, "ization")
		if MeasureGreaterThan0(pre) {
			word = pre + "ize"
			return word
		}

	}

	// For ATION suffix. ATION -> ATE
	ation := strings.HasSuffix(word, "ation")
	if ation {
		pre := strings.TrimSuffix(word, "ation")
		if MeasureGreaterThan0(pre) {
			word = pre + "ate"
			return word
		}

	}

	// For ATOR suffix. ATOR -> ATE
	ator := strings.HasSuffix(word, "ator")
	if ator {
		pre := strings.TrimSuffix(word, "ator")
		if MeasureGreaterThan0(pre) {
			word = pre + "ate"
			return word
		}

	}

	// For ALISM suffix. ALISM -> AL
	alism := strings.HasSuffix(word, "alism")
	if alism {
		pre := strings.TrimSuffix(word, "alism")
		if MeasureGreaterThan0(pre) {
			word = pre + "al"
			return word
		}

	}

	// For IVENESS suffix. IVENESS -> IVE
	iveness := strings.HasSuffix(word, "iveness")
	if iveness {
		pre := strings.TrimSuffix(word, "iveness")
		if MeasureGreaterThan0(pre) {
			word = pre + "ive"
			return word
		}

	}

	// For FULNESS suffix. FULNESS -> FUL
	fulness := strings.HasSuffix(word, "fulness")
	if fulness {
		pre := strings.TrimSuffix(word, "fulness")
		if MeasureGreaterThan0(pre) {
			word = pre + "ful"
			return word
		}

	}

	// For OUSNESS suffix. OUSNESS -> OUS
	ousness := strings.HasSuffix(word, "ousness")
	if ousness {
		pre := strings.TrimSuffix(word, "ousness")
		if MeasureGreaterThan0(pre) {
			word = pre + "ous"
			return word
		}

	}

	// For ALITI suffix. ALITI -> AL
	aliti := strings.HasSuffix(word, "aliti")
	if aliti {
		pre := strings.TrimSuffix(word, "aliti")
		if MeasureGreaterThan0(pre) {
			word = pre + "al"
			return word
		}

	}

	// For IVITI suffix. IVITI -> IVE
	iviti := strings.HasSuffix(word, "iviti")
	if iviti {
		pre := strings.TrimSuffix(word, "iviti")
		if MeasureGreaterThan0(pre) {
			word = pre + "ive"
			return word
		}

	}

	// For BILITI suffix. BILITI -> BLE
	biliti := strings.HasSuffix(word, "biliti")
	if biliti {
		pre := strings.TrimSuffix(word, "biliti")
		if MeasureGreaterThan0(pre) {
			word = pre + "ble"
			return word
		}

	}

	return word
}
