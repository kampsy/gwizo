/*Package gwizo is Next generation Go implementation of the Porter Stemmer algorithm. Specificaly
the suffix stripping. M.F.Porter 1980.
It is designed to be extensible so that developers can easily create new experiences
*/
package gwizo

import (
	"fmt"
	"strings"
)

// list of vowels and y.
const (
	y = "y"
	a = "a"
	e = "e"
	i = "i"
	o = "o"
	u = "u"
)

// Octopus collects information about the word.
type Octopus struct {
	Word    string // The word to be stemmed.
	VowCon  string // example vcvcvc. Where v = vowel and c = consonant.
	Measure int    // Number of times the pair vc appears.
}

// Ingest Returns the Octopus type
func Ingest(w string) Octopus {
	// Collection of vowels and consonants
	var collection []string
	// Change the word to lowercase letters.
	wordLower := strings.ToLower(w)
	for num := 0; num < len(wordLower); num++ {

		// Check for vowels at index 0 y included.
		if num == 0 {
			if string(wordLower[num]) == y || string(wordLower[num]) == a || string(wordLower[num]) == e ||
				string(wordLower[num]) == i || string(wordLower[num]) == o ||
				string(wordLower[num]) == u {
				collection = append(collection, "v")
			} else {
				collection = append(collection, "c")
			}
			continue
		}

		// If Y is preceded by a vowel it becomes a consonant and if Y is preceded
		// by a consonant it becomes a vowel.
		if collection[num-1] == "v" && string(wordLower[num]) == y {
			collection = append(collection, "c")
			continue
		} else if collection[num-1] == "c" && string(wordLower[num]) == y {
			collection = append(collection, "v")
			continue
		}

		if string(wordLower[num]) == a || string(wordLower[num]) == e ||
			string(wordLower[num]) == i || string(wordLower[num]) == o ||
			string(wordLower[num]) == u {
			collection = append(collection, "v")
		} else {
			collection = append(collection, "c")
		}

	}

	str := strings.Join(collection, "")
	var anl Octopus // Instance of Octopus.
	anl.Word = wordLower
	anl.VowCon = str
	anl.Measure = strings.Count(str, "vc")

	return anl
}

// This method remakes the VowCon and measure of the word for each step.
func (oct *Octopus) thinian(w string) {
	// Collection of vowels and consonants
	var collection []string
	// Change the word to lowercase letters.
	wordLower := strings.ToLower(w)
	for num := 0; num < len(wordLower); num++ {

		// Check for y at the beginning.
		if num == 0 {
			if string(wordLower[num]) == y || string(wordLower[num]) == a || string(wordLower[num]) == e ||
				string(wordLower[num]) == i || string(wordLower[num]) == o ||
				string(wordLower[num]) == u {
				collection = append(collection, "v")
			} else {
				collection = append(collection, "c")
			}
			continue
		}

		// If Y is preceded by a vowel Y becomes a consonant and if Y is preceded
		// by a consonant Y becomes a vowel.
		if collection[num-1] == "v" && string(wordLower[num]) == y {
			collection = append(collection, "c")
			continue
		} else if collection[num-1] == "c" && string(wordLower[num]) == y {
			collection = append(collection, "v")
			continue
		}

		if string(wordLower[num]) == a || string(wordLower[num]) == e ||
			string(wordLower[num]) == i || string(wordLower[num]) == o ||
			string(wordLower[num]) == u {
			collection = append(collection, "v")
		} else {
			collection = append(collection, "c")
		}

	}

	str := strings.Join(collection, "")
	oct.Word = wordLower
	oct.VowCon = str
	oct.Measure = strings.Count(str, "vc")
}

// implementation of String Method and so Stringer interface
func (oct *Octopus) String() string {
	return fmt.Sprintf("%s %s %d", oct.Word, oct.VowCon, oct.Measure)
}

// HasVowel returns bool of (*v*)
func (oct *Octopus) HasVowel() bool {
	return strings.Contains(oct.VowCon, "v")
}

// HasConsonant returns bool of (*c*)
func (oct *Octopus) HasConsonant() bool {
	return strings.Contains(oct.VowCon, "c")
}

// MeasureGreaterThan0 checks if measure value is grater than 0
func (oct *Octopus) MeasureGreaterThan0() bool {
	if oct.Measure > 0 {
		return true
	}
	return false
}

// MeasureGreaterThan1 checks if measure value is grater than 1
func (oct *Octopus) MeasureGreaterThan1() bool {
	if oct.Measure > 1 {
		return true
	}
	return false
}

// HascvcEndLastNotwxy checks if VowCon pattern ends with cvc, where second
// c is not W, X, Y
func HascvcEndLastNotwxy(str string) bool {
	nest := Ingest(str)
	cvc := strings.HasSuffix(nest.VowCon, "cvc")
	wlen := len(nest.Word)
	lastLetter := nest.Word[(wlen - 1)]
	word := string(lastLetter)
	w := strings.Contains(word, "w")
	x := strings.Contains(word, "x")
	y := strings.Contains(word, y)

	if cvc == true && w == false && x == false && y == false {
		return true
	}
	return false
}

// HasMeasureEqualTo1 checks if measure value = 1
func HasMeasureEqualTo1(str string) bool {
	nest := Ingest(str)
	if nest.Measure == 1 {
		return true
	}
	return false
}

// HasSameDoubleConsonant accepts a string as an argument, checks if it has
// double consonant as suffix and returns a boolean
func HasSameDoubleConsonant(str string) bool {
	nest := Ingest(str)
	cc := strings.HasSuffix(nest.VowCon, "cc")
	if cc == true {
		wlen := (len(str) - 1)
		letr := string(str[wlen])
		letr2 := string(str[(wlen - 1)])
		if letr == letr2 {
			return true
		}
	}
	return false
}

// HasEndst checks if word has suffix S or T
func (oct *Octopus) HasEndst() bool {
	s := strings.HasSuffix(oct.Word, "s")
	t := strings.HasSuffix(oct.Word, "t")

	if s == true || t == true {
		return true
	}
	return false
}

// HasEndl checks if word has suffix L
func (oct *Octopus) HasEndl() bool {
	l := strings.HasSuffix(oct.Word, "l")
	if l == true {
		return true
	}
	return false
}

/*Step1a deals with plurals and past participles. The subsequent steps are
much more straightforward.
Step 1A according the stemmer doc.*/
func (oct *Octopus) Step1a() string {
	str := oct.Word
	oct.thinian(str) // remake for the word

	// For SSES suffix. SSES -> SS
	sses := strings.HasSuffix(oct.Word, "sses")
	if sses == true {
		pre := strings.TrimSuffix(oct.Word, "sses")
		str = pre + "ss"
		return str
	}

	// For IES suffix. IES -> I
	ies := strings.HasSuffix(oct.Word, "ies")
	if ies == true {
		pre := strings.TrimSuffix(oct.Word, "ies")
		str = pre + i
		return str
	}

	// For SS suffix. SS -> SS
	ss := strings.HasSuffix(oct.Word, "ss")
	if ss == true {
		pre := strings.TrimSuffix(oct.Word, "ss")
		str = pre + "ss"
		return str
	}

	// For S suffix. S ->
	s := strings.HasSuffix(oct.Word, "s")
	if s == true {
		pre := strings.TrimSuffix(oct.Word, "s")
		str = pre
		return str
	}

	return str
}

/*Step1b according the stemmer doc.*/
func (oct *Octopus) Step1b() string {
	str := oct.Word
	oct.thinian(str) // remake for the word

	// Word Measure (m > 0) and EED suffix. EED -> EE
	eed := strings.HasSuffix(oct.Word, "eed")
	if eed == true && oct.Measure > 0 {
		if len(oct.Word) == 4 {
			str = oct.Word
		} else {
			pre := strings.TrimSuffix(oct.Word, "eed")
			str = pre + "ee"
		}
		return str
	}

	// Word has Vowel and ED suffix. ED ->
	ed := strings.HasSuffix(oct.Word, "ed")
	if oct.HasVowel() == true && ed == true {
		// word exception
		if len(oct.Word) == 4 {
			str = oct.Word
		} else {
			pre := strings.TrimSuffix(oct.Word, "ed")
			str = pre
		}
	}

	// Word has Vowel and ING suffix. ING ->
	ing := strings.HasSuffix(oct.Word, "ing")
	if oct.HasVowel() == true && ing == true {
		if len(str) == 4 {
			return str
		}
		pre := strings.TrimSuffix(oct.Word, "ing")
		str = pre
	}

	/* If the second or third of the above rules is successful the following
	is done.*/

	if ed == true {
		// Word has AT suffix. AT -> ATE
		at := strings.HasSuffix(str, "at")
		if at == true {
			pre := strings.TrimSuffix(str, "at")
			str = pre + "ate"
			return str
		}

		// Word has BL suffix. BL -> BLE
		bl := strings.HasSuffix(str, "bl")
		if bl == true {
			pre := strings.TrimSuffix(str, "bl")
			str = pre + "ble"
			return str
		}

		// Word has IZ suffix. IZ -> IZE
		iz := strings.HasSuffix(str, "iz")
		if iz == true {
			pre := strings.TrimSuffix(str, "iz")
			str = pre + "ize"
			return str
		}

	}

	// (*d and not (*L or *S or *Z)) -> single letter at the end
	if ed == true || ing == true {
		if HasSameDoubleConsonant(str) == true {
			ll := strings.HasSuffix(str, "ll")
			ss := strings.HasSuffix(str, "ss")
			zz := strings.HasSuffix(str, "zz")
			if ll == true || ss == true || zz == true {
				return str
			}
			strlen := len(str)
			lastLetter := str[:(strlen - 1)]
			str = string(lastLetter)

			return str
		}
	}

	// (m=1 and *o) -> E
	if ing == true {
		if HasMeasureEqualTo1(str) == true && HascvcEndLastNotwxy(str) == true {
			str = str + e
			return str
		}
	}

	return str
}

/*Step1c according the stemmer doc.*/
func (oct *Octopus) Step1c() string {
	str := oct.Word
	oct.thinian(str) // remake for the word

	// (*v*) Y -> I
	// Word has Vowel and Y suffix. Y -> I
	y := strings.HasSuffix(oct.Word, y)
	if oct.HasVowel() == true && y == true {
		if len(oct.Word) <= 3 {
			str = oct.Word
		} else {
			pre := strings.TrimSuffix(oct.Word, "y")
			str = pre + i
		}
	}

	return str
}

/*Step2 Rules according to the stemmer doc.
=========================================*/
func (oct *Octopus) Step2() string {
	str := oct.Word
	oct.thinian(str) // remake for the word
	// word exception
	rational := "rational"

	if oct.MeasureGreaterThan0() == true {

		// For ATIONAL suffix. ATIONAL -> ATE
		ational := strings.HasSuffix(oct.Word, "ational")
		if ational == true {
			if oct.Word == rational {
				str = rational
			} else {
				pre := strings.TrimSuffix(oct.Word, "ational")
				str = pre + "ate"
			}
			return str
		}

		// For TIONAL suffix. TIONAL -> TION
		tional := strings.HasSuffix(oct.Word, "tional")
		if tional == true {
			if oct.Word == rational {
				str = rational
			} else {
				pre := strings.TrimSuffix(oct.Word, "tional")
				str = pre + "tion"
			}
			return str
		}

		// For ENCI suffix. ENCI -> ENCE
		enci := strings.HasSuffix(oct.Word, "enci")
		if enci == true {
			pre := strings.TrimSuffix(oct.Word, "enci")
			str = pre + "ence"
			return str
		}

		// For ANCI suffix. ANCI -> ANCE
		anci := strings.HasSuffix(oct.Word, "anci")
		if anci == true {
			pre := strings.TrimSuffix(oct.Word, "anci")
			str = pre + "ance"
			return str
		}

		// For IZER suffix. IZER -> IZE
		izer := strings.HasSuffix(oct.Word, "izer")
		if izer == true {
			pre := strings.TrimSuffix(oct.Word, "izer")
			str = pre + "ize"
			return str
		}

		// For ABLI suffix. ABLI -> ABLE
		abli := strings.HasSuffix(oct.Word, "abli")
		if abli == true {
			pre := strings.TrimSuffix(oct.Word, "abli")
			str = pre + "able"
			return str
		}

		// For ALLI suffix. ALLI -> AL
		alli := strings.HasSuffix(oct.Word, "alli")
		if alli == true {
			pre := strings.TrimSuffix(oct.Word, "alli")
			str = pre + "al"
			return str
		}

		// For ENTLI suffix. ENTLI -> ENT
		entli := strings.HasSuffix(oct.Word, "entli")
		if entli == true {
			pre := strings.TrimSuffix(oct.Word, "entli")
			str = pre + "ent"
			return str
		}

		// For ELI suffix. ELI -> E
		eli := strings.HasSuffix(oct.Word, "eli")
		if eli == true {
			pre := strings.TrimSuffix(oct.Word, "eli")
			str = pre + e
			return str
		}

		// For OUSLI suffix. OUSLI -> OUS
		ousli := strings.HasSuffix(oct.Word, "ousli")
		if ousli == true {
			pre := strings.TrimSuffix(oct.Word, "ousli")
			str = pre + "ous"
			return str
		}

		// For IZATION suffix. IZATION -> IZE
		ization := strings.HasSuffix(oct.Word, "ization")
		if ization == true {
			pre := strings.TrimSuffix(oct.Word, "ization")
			str = pre + "ize"
			return str
		}

		// For ATION suffix. ATION -> ATE
		ation := strings.HasSuffix(oct.Word, "ation")
		if ation == true {
			pre := strings.TrimSuffix(oct.Word, "ation")
			str = pre + "ate"
			return str
		}

		// For ATOR suffix. ATOR -> ATE
		ator := strings.HasSuffix(oct.Word, "ator")
		if ator == true {
			pre := strings.TrimSuffix(oct.Word, "ator")
			str = pre + "ate"
			return str
		}

		// For ALISM suffix. ALISM -> AL
		alism := strings.HasSuffix(oct.Word, "alism")
		if alism == true {
			pre := strings.TrimSuffix(oct.Word, "alism")
			str = pre + "al"
			return str
		}

		// For IVENESS suffix. IVENESS -> IVE
		iveness := strings.HasSuffix(oct.Word, "iveness")
		if iveness == true {
			pre := strings.TrimSuffix(oct.Word, "iveness")
			str = pre + "ive"
			return str
		}

		// For FULNESS suffix. FULNESS -> FUL
		fulness := strings.HasSuffix(oct.Word, "fulness")
		if fulness == true {
			pre := strings.TrimSuffix(oct.Word, "fulness")
			str = pre + "ful"
			return str
		}

		// For OUSNESS suffix. OUSNESS -> OUS
		ousness := strings.HasSuffix(oct.Word, "ousness")
		if ousness == true {
			pre := strings.TrimSuffix(oct.Word, "ousness")
			str = pre + "ous"
			return str
		}

		// For ALITI suffix. ALITI -> AL
		aliti := strings.HasSuffix(oct.Word, "aliti")
		if aliti == true {
			pre := strings.TrimSuffix(oct.Word, "aliti")
			str = pre + "al"
			return str
		}

		// For IVITI suffix. IVITI -> IVE
		iviti := strings.HasSuffix(oct.Word, "iviti")
		if iviti == true {
			pre := strings.TrimSuffix(oct.Word, "iviti")
			str = pre + "ive"
			return str
		}

		// For BILITI suffix. BILITI -> BLE
		biliti := strings.HasSuffix(oct.Word, "biliti")
		if biliti == true {
			pre := strings.TrimSuffix(oct.Word, "biliti")
			str = pre + "ble"
			return str
		}

	}

	return str
}

/*Step3 rules according the stemmer doc.
=====================================*/
func (oct *Octopus) Step3() string {
	str := oct.Word
	oct.thinian(str) // remake for the word

	if oct.MeasureGreaterThan0() == true {
		// For ICATE suffix. ICATE -> IC
		icate := strings.HasSuffix(oct.Word, "icate")
		if icate == true {
			pre := strings.TrimSuffix(oct.Word, "icate")
			str = pre + "ic"
			return str
		}

		// For ATIVE suffix. ATIVE ->
		ative := strings.HasSuffix(oct.Word, "ative")
		if ative == true {
			pre := strings.TrimSuffix(oct.Word, "ative")
			str = pre
			return str
		}

		// For ALIZE suffix. ALIZE -> AL
		alize := strings.HasSuffix(oct.Word, "alize")
		if alize == true {
			pre := strings.TrimSuffix(oct.Word, "alize")
			str = pre + "al"
			return str
		}

		// For ICITI suffix. ICITI -> IC
		iciti := strings.HasSuffix(oct.Word, "iciti")
		if iciti == true {
			pre := strings.TrimSuffix(oct.Word, "iciti")
			str = pre + "ic"
			return str
		}

		// For ICAL suffix. ICAL -> IC
		ical := strings.HasSuffix(oct.Word, "ical")
		if ical == true {
			pre := strings.TrimSuffix(oct.Word, "ical")
			str = pre + "ic"
			return str
		}

		// For FUL suffix. FUL ->
		ful := strings.HasSuffix(oct.Word, "ful")
		if ful == true {
			pre := strings.TrimSuffix(oct.Word, "ful")
			str = pre
			return str
		}

		// For NESS suffix. NESS ->
		ness := strings.HasSuffix(oct.Word, "ness")
		if ness == true {
			pre := strings.TrimSuffix(oct.Word, "ness")
			str = pre
			return str
		}

	}

	return str
}

/*Step4 rules according the stemmer doc.
The suffixes will now be removed
=====================================*/
func (oct *Octopus) Step4() string {
	str := oct.Word
	oct.thinian(str) // remake for the word

	if oct.MeasureGreaterThan1() == true {
		// For AL suffix. AL ->
		al := strings.HasSuffix(oct.Word, "al")
		if al == true {
			pre := strings.TrimSuffix(oct.Word, "al")
			str = pre
			return str
		}

		// For ANCE suffix. ANCE ->
		ance := strings.HasSuffix(oct.Word, "ance")
		if ance == true {
			pre := strings.TrimSuffix(oct.Word, "ance")
			str = pre
			return str
		}

		// For ENCE suffix. ENCE ->
		ence := strings.HasSuffix(oct.Word, "ence")
		if ence == true {
			pre := strings.TrimSuffix(oct.Word, "ence")
			str = pre
			return str
		}

		// For ER suffix. ER ->
		er := strings.HasSuffix(oct.Word, "er")
		if er == true {
			pre := strings.TrimSuffix(oct.Word, "er")
			str = pre
			return str
		}

		// For IC suffix. IC ->
		ic := strings.HasSuffix(oct.Word, "ic")
		if ic == true {
			pre := strings.TrimSuffix(oct.Word, "ic")
			str = pre
			return str
		}

		// For ABLE suffix. ABLE ->
		able := strings.HasSuffix(oct.Word, "able")
		if able == true {
			pre := strings.TrimSuffix(oct.Word, "able")
			str = pre
			return str
		}

		// For IBLE suffix. IBLE ->
		ible := strings.HasSuffix(oct.Word, "ible")
		if ible == true {
			pre := strings.TrimSuffix(oct.Word, "ible")
			str = pre
			return str
		}

		// For ANT suffix. ANT ->
		ant := strings.HasSuffix(oct.Word, "ant")
		if ant == true {
			pre := strings.TrimSuffix(oct.Word, "ant")
			str = pre
			return str
		}

		// For EMENT suffix. EMENT ->
		ement := strings.HasSuffix(oct.Word, "ement")
		if ement == true {
			pre := strings.TrimSuffix(oct.Word, "ement")
			str = pre
			return str
		}

		// For MENT suffix. MENT ->
		ment := strings.HasSuffix(oct.Word, "ment")
		if ment == true {
			pre := strings.TrimSuffix(oct.Word, "ment")
			str = pre
			return str
		}

		// For ENT suffix. ENT ->
		ent := strings.HasSuffix(oct.Word, "ent")
		if ent == true {
			pre := strings.TrimSuffix(oct.Word, "ent")
			str = pre
			return str
		}

		if oct.HasEndst() == false { // Made Personal correction from true to false.
			// For ION suffix. ION ->
			ion := strings.HasSuffix(oct.Word, "ion")
			if ion == true {
				pre := strings.TrimSuffix(oct.Word, "ion")
				str = pre
				return str
			}
		}

		// For OU suffix. OU ->
		ou := strings.HasSuffix(oct.Word, "ou")
		if ou == true {
			pre := strings.TrimSuffix(oct.Word, "ou")
			str = pre
			return str
		}

		// For ISM suffix. ISM ->
		ism := strings.HasSuffix(oct.Word, "ism")
		if ism == true {
			pre := strings.TrimSuffix(oct.Word, "ism")
			str = pre
			return str
		}

		// For ATE suffix. ATE ->
		ate := strings.HasSuffix(oct.Word, "ate")
		if ate == true {
			pre := strings.TrimSuffix(oct.Word, "ate")
			str = pre
			return str
		}

		// For ITI suffix. ITI ->
		iti := strings.HasSuffix(oct.Word, "iti")
		if iti == true {
			pre := strings.TrimSuffix(oct.Word, "iti")
			str = pre
			return str
		}

		// For OUS suffix. OUS ->
		ous := strings.HasSuffix(oct.Word, "ous")
		if ous == true {
			pre := strings.TrimSuffix(oct.Word, "ous")
			str = pre
			return str
		}

		// For IVE suffix. IVE ->
		ive := strings.HasSuffix(oct.Word, "ive")
		if ive == true {
			pre := strings.TrimSuffix(oct.Word, "ive")
			str = pre
			return str
		}

		// For IZE suffix. IZE ->
		ize := strings.HasSuffix(oct.Word, "ize")
		if ize == true {
			pre := strings.TrimSuffix(oct.Word, "ize")
			str = pre
			return str
		}

	}

	return str
}

/*Step5a rules according the stemmer doc.
little tidying up
=====================================*/
func (oct *Octopus) Step5a() string {
	str := oct.Word
	oct.thinian(str) // remake for the word

	if oct.MeasureGreaterThan1() == true {
		// E suffix. E ->
		e := strings.HasSuffix(oct.Word, e)
		if e == true {
			if len(oct.Word) > 4 {
				pre := strings.TrimSuffix(oct.Word, "e")
				str = pre
				return str
			}
		}
	}

	if HasMeasureEqualTo1(oct.Word) == true && HascvcEndLastNotwxy(oct.Word) == false {
		// (m=1 and not *o) E ->
		e := strings.HasSuffix(oct.Word, e)
		if e == true {
			if len(oct.Word) > 4 { // this helps the above function
				pre := strings.TrimSuffix(oct.Word, "e")
				str = pre
				return str
			}
		}
	}

	return str
}

/*Step5b according the stemmer doc.*/
func (oct *Octopus) Step5b() string {
	str := oct.Word
	oct.thinian(str) // remake for the word

	if oct.MeasureGreaterThan1() == true {
		if HasSameDoubleConsonant(oct.Word) == true && oct.HasEndl() == true {
			w := oct.Word
			strlen := len(w)
			lastLetter := w[:(strlen - 1)]
			str = string(lastLetter)
		}
	} else {
		str = oct.Word
	}

	return str
}

/*ShallowStem Returns the stem of the word. This method bells out when the word is Changed
==============================*/
func (oct *Octopus) ShallowStem() string {
	str := oct.Word
	if len(oct.Word) > 4 {
		var slice []string

		mytype := oct

		mystr := mytype.Step1a()
		if mystr != mytype.Word {
			slice = append(slice, mystr)
		}

		mytype2 := oct
		mystr2 := mytype2.Step1b()
		if mystr2 != mytype2.Word {
			slice = append(slice, mystr2)
		}

		mytype3 := oct
		mystr3 := mytype3.Step1c()
		if mystr3 != mytype3.Word {
			slice = append(slice, mystr3)
		}

		mytype4 := oct
		mystr4 := mytype4.Step2()
		if mystr4 != mytype4.Word {
			slice = append(slice, mystr4)
		}

		mytype5 := oct
		mystr5 := mytype5.Step3()
		if mystr5 != mytype5.Word {
			slice = append(slice, mystr5)
		}

		mytype6 := oct
		mystr6 := mytype6.Step4()
		if mystr6 != mytype6.Word {
			slice = append(slice, mystr6)
		}

		mytype7 := oct
		mystr7 := mytype7.Step5a()
		if mystr7 != mytype7.Word {
			slice = append(slice, mystr7)
		}

		mytype8 := oct
		mystr8 := mytype8.Step5b()
		if mystr8 != mytype8.Word {
			slice = append(slice, mystr8)
		}

		if len(slice) > 0 {
			str = slice[0]
		}

		return str
	}

	return str
}

/*ShallowStemmed Returns the Step that was used to stem the word
=========================================*/
func (oct *Octopus) ShallowStemmed() string {
	stepUsed := fmt.Sprintf("None %s is a stem", oct.Word)
	if len(oct.Word) > 4 {
		mytype := oct

		mystr := mytype.Step1a()
		if mystr != mytype.Word {
			stepUsed = "Step1a()"
			return stepUsed
		}

		mytype.Word = mystr
		mystr2 := mytype.Step1b()
		if mystr2 != mytype.Word {
			stepUsed = "Step1b()"
			return stepUsed
		}

		mytype.Word = mystr2
		mystr3 := mytype.Step1c()
		if mystr3 != mytype.Word {
			stepUsed = "Step1c()"
			return stepUsed
		}

		mytype.Word = mystr3
		mystr4 := mytype.Step2()
		if mystr4 != mytype.Word {
			stepUsed = "Step2()"
			return stepUsed
		}

		mytype.Word = mystr4
		mystr5 := mytype.Step3()
		if mystr5 != mytype.Word {
			stepUsed = "Step3()"
			return stepUsed
		}

		mytype.Word = mystr5
		mystr6 := mytype.Step4()
		if mystr6 != mytype.Word {
			stepUsed = "Step4()"
			return stepUsed
		}

		mytype.Word = mystr6
		mystr7 := mytype.Step5a()
		if mystr7 != mytype.Word {
			stepUsed = "Step5a()"
			return stepUsed
		}

		mytype.Word = mystr7
		mystr8 := mytype.Step5b()
		if mystr8 != mytype.Word {
			stepUsed = "Step5b()"
			return stepUsed
		}
	}

	return stepUsed
}

/*DeepStem makes sure that the word goes through every step in the algorithm.
  Any change made by one step is passed on to the next step
================================================================*/
func (oct *Octopus) DeepStem() string {
	str := oct.Word
	if len(oct.Word) > 4 {
		octopus := oct

		tentacle := octopus.Step1a()

		octopus.Word = tentacle
		tentacle2 := octopus.Step1b()

		octopus.Word = tentacle2
		tentacle3 := octopus.Step1c()

		octopus.Word = tentacle3
		tentacle4 := octopus.Step2()

		octopus.Word = tentacle4
		tentacle5 := octopus.Step3()

		octopus.Word = tentacle5
		tentacle6 := octopus.Step4()

		octopus.Word = tentacle6
		tentacle7 := octopus.Step5a()

		octopus.Word = tentacle7
		tentacle8 := octopus.Step5b()

		str = tentacle8
		return str
	}

	return str
}
