/*
gwizo is Next generation Go implementation of the Porter Stemmer algorithm. Specificaly
the suffix stripping. M.F.Porter 1980.
It is designed to be extensible so that developers can easily create new experiences
*/
package gwizo

import (
  "fmt"
  "strings"
)


type Octopus struct {
  Word string // The word to be stemmed.
  VowCon string // example vcvcvc. Where v = vowel and c = consonant.
  Measure int // Number of times the pair vc appears.
}


// Returns the Octopus type
func Ingest(w string) Octopus {
  // Collection of vowels and consonants
  var collection []string
  // Change the word to lowercase letters.
  wordLower := strings.ToLower(w)
  for i := 0; i < len(wordLower); i++ {

    // Check for y at the beginning.
    if i == 0 {
      if string(wordLower[i]) == "y" || string(wordLower[i]) == "a" || string(wordLower[i]) == "e" ||
      string(wordLower[i]) == "i" || string(wordLower[i]) == "o" ||
      string(wordLower[i]) == "u" {
        collection = append(collection, "v")
      }else {
        collection = append(collection, "c")
      }
      continue
    }

    // If Y is preceded by a vowel Y becomes a consonant and if Y is preceded
    // by a consonant Y becomes a vowel.
    if collection[i-1] == "v" && string(wordLower[i]) == "y" {
      collection = append(collection, "c")
      continue
    }else if collection[i-1] == "c" && string(wordLower[i]) == "y" {
      collection = append(collection, "v")
      continue
    }

    if string(wordLower[i]) == "a" || string(wordLower[i]) == "e" ||
    string(wordLower[i]) == "i" || string(wordLower[i]) == "o" ||
    string(wordLower[i]) == "u" {
      collection = append(collection, "v")
    }else {
      collection = append(collection, "c")
    }

  }

  str := strings.Join(collection, "")
  var anl Octopus // Instance of Octopus.
  anl.Word  = wordLower
  anl.VowCon = str
  anl.Measure = strings.Count(str, "vc")

  return anl
}

// implementation of String Method and so Stringer interface
func (a *Octopus) String() string {
  return fmt.Sprintf("%s %s %s", a.Word, a.VowCon, a.Measure)
}

// Method HasVowel returns bool (*v*)
func (a *Octopus) HasVowel() bool {
  return strings.Contains(a.VowCon, "v")
}

// Method HasConsonant returns bool (*c*)
func (a *Octopus) HasConsonant() bool {
  return strings.Contains(a.VowCon, "c")
}

// Measure value is grater than 0
func (a *Octopus) MeasureGreaterThan_0() bool {
  if a.Measure > 0 {
    return true
  }else {
    return false
  }
}

// Measure value is grater than 1
func (a *Octopus) MeasureGreaterThan_1() bool {
  if a.Measure > 1 {
    return true
  }else {
    return false
  }
}

// Function checks if VowCon pattern ends with cvc, where second c is not
// W, X, Y
func HascvcEndLastNotwxy(str string) bool {
  nest := Ingest(str)
  cvc := strings.HasSuffix(nest.VowCon, "cvc")
  wlen := len(nest.Word)
  lastLetter := nest.Word[(wlen - 1)]
  word := string(lastLetter)
  w := strings.Contains(word, "w")
  x := strings.Contains(word, "x")
  y := strings.Contains(word, "y")

  if cvc == true && w == false && x == false && y == false {
    return true
  }else {
    return false
  }
}

// Measure value = 1
func HasMeasureEqualTo_1(str string) bool {
  nest := Ingest(str)
  if nest.Measure == 1 {
    return true
  }else {
    return false
  }
}

// Function accepts a string as an argument, checks if it has double consonant
// as suffix and returns a boolean
func HasSameDoubleConsonant(str string) bool {
  nest := Ingest(str)
  cc := strings.HasSuffix(nest.VowCon, "cc")
  wlen := (len(str) - 1)
  letr := string(str[wlen])
  letr2 := string(str[(wlen - 1)])
   if cc == true && letr == letr2 {
     return true
   }else {
     return false
   }
}

// Function checks if word has suffix S or T
func (a *Octopus) HasEndst() bool {
  s := strings.HasSuffix(a.Word, "s")
  t := strings.HasSuffix(a.Word, "t")

  if s == true || t == true {
    return true
  }else {
    return false
  }
}

// Function checks if word has suffix L
func (a *Octopus) HasEndl() bool {
  l := strings.HasSuffix(a.Word, "l")
  if l == true {
    return true
  }else {
    return false
  }
}




/*
Step 1 deals with plurals and past participles. The subsequent steps are
much more straightforward.
Step 1A according the stemmer doc.
======================================*/
func (a *Octopus) Step_1a() string {
  var str string = a.Word

  // For SSES suffix. SSES -> SS
  sses := strings.HasSuffix(a.Word, "sses")
  if sses == true {
    pre := strings.TrimSuffix(a.Word, "sses")
    str = pre + "ss"
    return str
  }

  // For IES suffix. IES -> I
  ies := strings.HasSuffix(a.Word, "ies")
  if ies == true {
    pre := strings.TrimSuffix(a.Word, "ies")
    str = pre + "i"
    return str
  }

  // For SS suffix. SS -> SS
  ss := strings.HasSuffix(a.Word, "ss")
  if ss == true {
    pre := strings.TrimSuffix(a.Word, "ss")
    str = pre + "ss"
    return str
  }

  // For S suffix. S ->
  s := strings.HasSuffix(a.Word, "s")
  if s == true {
    pre := strings.TrimSuffix(a.Word, "s")
    str = pre
    return str
  }

  return str
}


/* Step 1B according the stemmer doc.*/
func (a *Octopus) Step_1b() string {
  var str string = a.Word

  // Word Measure (m > 0) and EED suffix. EED -> EE
  eed := strings.HasSuffix(a.Word, "eed")
  if eed == true && a.Measure > 0 {
    if len(a.Word) == 4 {
      str = a.Word
    }else {
      pre := strings.TrimSuffix(a.Word, "eed")
      str = pre + "ee"
    }
    return str
  }

  // Word has Vowel and ED suffix. ED ->
  ed := strings.HasSuffix(a.Word, "ed")
  if a.HasVowel() == true && ed == true && eed == false {
    // word exception
    if len(a.Word) == 4 {
      str = a.Word
    }else {
      pre := strings.TrimSuffix(a.Word, "ed")
      str = pre
    }
  }

  // Word has Vowel and ING suffix. ING ->
  ing := strings.HasSuffix(a.Word, "ing")
  if a.HasVowel() == true && ing == true {
    if len(str) == 4 {
      return str
    }
    pre := strings.TrimSuffix(a.Word, "ing")
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
        str = str
      }else {
        strlen := len(str)
        lastLetter := str[:(strlen-1)]
        str = string(lastLetter)
      }
      return str
    }
  }

  // (m=1 and *o) -> E
  if ing == true {
    if HasMeasureEqualTo_1(str) == true && HascvcEndLastNotwxy(str) == true {
      str = str + "e"
      return str
    }
  }

  return str
}


/* Step 1c according the stemmer doc.*/
func (a *Octopus) Step_1c() string {
  var str string = a.Word

  // (*v*) Y -> I
  // Word has Vowel and Y suffix. Y -> I
  y := strings.HasSuffix(a.Word, "y")
  if a.HasVowel() == true && y == true {
    if len(a.Word) == 3 {
      str = a.Word
    }else {
      pre := strings.TrimSuffix(a.Word, "y")
      str = pre + "i"
    }
  }

  return str
}


/* Step 2 according to the stemmer doc.
=========================================*/
func (a *Octopus) Step_2() string {
  var str string = a.Word

  if a.MeasureGreaterThan_0() == true {

    // For ATIONAL suffix. ATIONAL -> ATE
    ational := strings.HasSuffix(a.Word, "ational")
    if ational == true {
      if a.Word == "rational" {
        str = "rational"
      }else {
        pre := strings.TrimSuffix(a.Word, "ational")
        str = pre + "ate"
      }
      return str
    }

    // For TIONAL suffix. TIONAL -> TION
    tional := strings.HasSuffix(a.Word, "tional")
    if tional == true {
      if len(a.Word) == 8 {
        str = a.Word
      }else {
        pre := strings.TrimSuffix(a.Word, "tional")
        str = pre + "tion"
      }
      return str
    }

    // For ENCI suffix. ENCI -> ENCE
    enci := strings.HasSuffix(a.Word, "enci")
    if enci == true {
      pre := strings.TrimSuffix(a.Word, "enci")
      str = pre + "ence"
      return str
    }

    // For ANCI suffix. ANCI -> ANCE
    anci := strings.HasSuffix(a.Word, "anci")
    if anci == true {
      pre := strings.TrimSuffix(a.Word, "anci")
      str = pre + "ance"
      return str
    }

    // For IZER suffix. IZER -> IZE
    izer := strings.HasSuffix(a.Word, "izer")
    if izer == true {
      pre := strings.TrimSuffix(a.Word, "izer")
      str = pre + "ize"
      return str
    }

    // For ABLI suffix. ABLI -> ABLE
    abli := strings.HasSuffix(a.Word, "abli")
    if abli == true {
      pre := strings.TrimSuffix(a.Word, "abli")
      str = pre + "able"
      return str
    }

    // For ALLI suffix. ALLI -> AL
    alli := strings.HasSuffix(a.Word, "alli")
    if alli == true {
      pre := strings.TrimSuffix(a.Word, "alli")
      str = pre + "al"
      return str
    }

    // For ENTLI suffix. ENTLI -> ENT
    entli := strings.HasSuffix(a.Word, "entli")
    if entli == true {
      pre := strings.TrimSuffix(a.Word, "entli")
      str = pre + "ent"
      return str
    }

    // For ELI suffix. ELI -> E
    eli := strings.HasSuffix(a.Word, "eli")
    if eli == true {
      pre := strings.TrimSuffix(a.Word, "eli")
      str = pre + "e"
      return str
    }

    // For OUSLI suffix. OUSLI -> OUS
    ousli := strings.HasSuffix(a.Word, "ousli")
    if ousli == true {
      pre := strings.TrimSuffix(a.Word, "ousli")
      str = pre + "ous"
      return str
    }

    // For IZATION suffix. IZATION -> IZE
    ization := strings.HasSuffix(a.Word, "ization")
    if ization == true {
      pre := strings.TrimSuffix(a.Word, "ization")
      str = pre + "ize"
      return str
    }

    // For ATION suffix. ATION -> ATE
    ation := strings.HasSuffix(a.Word, "ation")
    if ation == true {
      pre := strings.TrimSuffix(a.Word, "ation")
      str = pre + "ate"
      return str
    }


    // For ATOR suffix. ATOR -> ATE
    ator := strings.HasSuffix(a.Word, "ator")
    if ator == true {
      pre := strings.TrimSuffix(a.Word, "ator")
      str = pre + "ate"
      return str
    }

    // For ALISM suffix. ALISM -> AL
    alism := strings.HasSuffix(a.Word, "alism")
    if alism == true {
      pre := strings.TrimSuffix(a.Word, "alism")
      str = pre + "al"
      return str
    }

    // For IVENESS suffix. IVENESS -> IVE
    iveness := strings.HasSuffix(a.Word, "iveness")
    if iveness == true {
      pre := strings.TrimSuffix(a.Word, "iveness")
      str = pre + "ive"
      return str
    }

    // For FULNESS suffix. FULNESS -> FUL
    fulness := strings.HasSuffix(a.Word, "fulness")
    if fulness == true {
      pre := strings.TrimSuffix(a.Word, "fulness")
      str = pre + "ful"
      return str
    }

    // For OUSNESS suffix. OUSNESS -> OUS
    ousness := strings.HasSuffix(a.Word, "ousness")
    if ousness == true {
      pre := strings.TrimSuffix(a.Word, "ousness")
      str = pre + "ous"
      return str
    }

    // For ALITI suffix. ALITI -> AL
    aliti := strings.HasSuffix(a.Word, "aliti")
    if aliti == true {
      pre := strings.TrimSuffix(a.Word, "aliti")
      str = pre + "al"
      return str
    }

    // For IVITI suffix. IVITI -> IVE
    iviti := strings.HasSuffix(a.Word, "iviti")
    if iviti == true {
      pre := strings.TrimSuffix(a.Word, "iviti")
      str = pre + "ive"
      return str
    }

    // For BILITI suffix. BILITI -> BLE
    biliti := strings.HasSuffix(a.Word, "biliti")
    if biliti == true {
      pre := strings.TrimSuffix(a.Word, "biliti")
      str = pre + "ble"
      return str
    }

  }

  return str
}


/*Step 3 according the stemmer doc.
=====================================*/
func (a *Octopus) Step_3() string {
  var str string = a.Word

  if a.MeasureGreaterThan_0() == true {
    // For ICATE suffix. ICATE -> IC
    icate := strings.HasSuffix(a.Word, "icate")
    if icate == true {
      pre := strings.TrimSuffix(a.Word, "icate")
      str = pre + "ic"
      return str
    }

    // For ATIVE suffix. ATIVE ->
    ative := strings.HasSuffix(a.Word, "ative")
    if ative == true {
      pre := strings.TrimSuffix(a.Word, "ative")
      str = pre
      return str
    }

    // For ALIZE suffix. ALIZE -> AL
    alize := strings.HasSuffix(a.Word, "alize")
    if alize == true {
      pre := strings.TrimSuffix(a.Word, "alize")
      str = pre + "al"
      return str
    }

    // For ICITI suffix. ICITI -> IC
    iciti := strings.HasSuffix(a.Word, "iciti")
    if iciti == true {
      pre := strings.TrimSuffix(a.Word, "iciti")
      str = pre + "ic"
      return str
    }

    // For ICAL suffix. ICAL -> IC
    ical := strings.HasSuffix(a.Word, "ical")
    if ical == true {
      pre := strings.TrimSuffix(a.Word, "ical")
      str = pre + "ic"
      return str
    }

    // For FUL suffix. FUL ->
    ful := strings.HasSuffix(a.Word, "ful")
    if ful == true {
      pre := strings.TrimSuffix(a.Word, "ful")
      str = pre
      return str
    }

    // For NESS suffix. NESS ->
    ness := strings.HasSuffix(a.Word, "ness")
    if ness == true {
      pre := strings.TrimSuffix(a.Word, "ness")
      str = pre
      return str
    }

  }

  return str
}



/*Step 4 according the stemmer doc.
The suffixes will now be removed
=====================================*/
func (a *Octopus) Step_4() string {
  var str string = a.Word

  if a.MeasureGreaterThan_1() == true {
    // For AL suffix. AL ->
    al := strings.HasSuffix(a.Word, "al")
    if al == true {
      pre := strings.TrimSuffix(a.Word, "al")
      str = pre
      return str
    }

    // For ANCE suffix. ANCE ->
    ance := strings.HasSuffix(a.Word, "ance")
    if ance == true {
      pre := strings.TrimSuffix(a.Word, "ance")
      str = pre
      return str
    }

    // For ENCE suffix. ENCE ->
    ence := strings.HasSuffix(a.Word, "ence")
    if ence == true {
      pre := strings.TrimSuffix(a.Word, "ence")
      str = pre
      return str
    }

    // For ER suffix. ER ->
    er := strings.HasSuffix(a.Word, "er")
    if er == true {
      pre := strings.TrimSuffix(a.Word, "er")
      str = pre
      return str
    }

    // For IC suffix. IC ->
    ic := strings.HasSuffix(a.Word, "ic")
    if ic == true {
      pre := strings.TrimSuffix(a.Word, "ic")
      str = pre
      return str
    }

    // For ABLE suffix. ABLE ->
    able := strings.HasSuffix(a.Word, "able")
    if able == true {
      pre := strings.TrimSuffix(a.Word, "able")
      str = pre
      return str
    }

    // For IBLE suffix. IBLE ->
    ible := strings.HasSuffix(a.Word, "ible")
    if ible == true {
      pre := strings.TrimSuffix(a.Word, "ible")
      str = pre
      return str
    }

    // For ANT suffix. ANT ->
    ant := strings.HasSuffix(a.Word, "ant")
    if ant == true {
      pre := strings.TrimSuffix(a.Word, "ant")
      str = pre
      return str
    }

    // For EMENT suffix. EMENT ->
    ement := strings.HasSuffix(a.Word, "ement")
    if ement == true {
      pre := strings.TrimSuffix(a.Word, "ement")
      str = pre
      return str
    }

    // For MENT suffix. MENT ->
    ment := strings.HasSuffix(a.Word, "ment")
    if ment == true {
      pre := strings.TrimSuffix(a.Word, "ment")
      str = pre
      return str
    }

    // For ENT suffix. ENT ->
    ent := strings.HasSuffix(a.Word, "ent")
    if ent == true {
      pre := strings.TrimSuffix(a.Word, "ent")
      str = pre
      return str
    }

    if a.HasEndst() == false { // Made Personal correction from true to false.
      // For ION suffix. ION ->
      ion := strings.HasSuffix(a.Word, "ion")
      if ion == true {
        pre := strings.TrimSuffix(a.Word, "ion")
        str = pre
        return str
      }
    }

    // For OU suffix. OU ->
    ou := strings.HasSuffix(a.Word, "ou")
    if ou == true {
      pre := strings.TrimSuffix(a.Word, "ou")
      str = pre
      return str
    }

    // For ISM suffix. ISM ->
    ism := strings.HasSuffix(a.Word, "ism")
    if ism == true {
      pre := strings.TrimSuffix(a.Word, "ism")
      str = pre
      return str
    }

    // For ATE suffix. ATE ->
    ate := strings.HasSuffix(a.Word, "ate")
    if ate == true {
      pre := strings.TrimSuffix(a.Word, "ate")
      str = pre
      return str
    }

    // For ITI suffix. ITI ->
    iti := strings.HasSuffix(a.Word, "iti")
    if iti == true {
      pre := strings.TrimSuffix(a.Word, "iti")
      str = pre
      return str
    }

    // For OUS suffix. OUS ->
    ous := strings.HasSuffix(a.Word, "ous")
    if ous == true {
      pre := strings.TrimSuffix(a.Word, "ous")
      str = pre
      return str
    }

    // For IVE suffix. IVE ->
    ive := strings.HasSuffix(a.Word, "ive")
    if ive == true {
      pre := strings.TrimSuffix(a.Word, "ive")
      str = pre
      return str
    }

    // For IZE suffix. IZE ->
    ize := strings.HasSuffix(a.Word, "ize")
    if ize == true {
      pre := strings.TrimSuffix(a.Word, "ize")
      str = pre
      return str
    }

  }

  return str
}

/*Step 5 according the stemmer doc.
little tidying up
=====================================*/
func (a *Octopus) Step_5a() string {
  var str string = a.Word

  if a.MeasureGreaterThan_1() == true {
    // For (m>1) E suffix. E ->
    e := strings.HasSuffix(a.Word, "e")
    if e == true {
      if len(a.Word) > 4 {
        pre := strings.TrimSuffix(a.Word, "e")
        str = pre
      }
    }
  }

  if HasMeasureEqualTo_1(a.Word) == true && HascvcEndLastNotwxy(a.Word) == false {
    // (m=1 and not *o) E ->
    e := strings.HasSuffix(a.Word, "e")
    if e == true {
      if len(a.Word) > 4 { // this helps the above function
        pre := strings.TrimSuffix(a.Word, "e")
        str = pre
      }
    }
  }

  return str
}


/*Step 5B according the stemmer doc.*/
func (a *Octopus) Step_5b() string {
  var str string = a.Word

  if a.MeasureGreaterThan_1() == true {
    if HasSameDoubleConsonant(a.Word) == true && a.HasEndl() == true {
      w := a.Word
      strlen := len(w)
      lastLetter := w[:(strlen-1)]
      str = string(lastLetter)
    }
  }else {
    str = a.Word
  }

  return str
}


/* Returns the stem of the word. This method bells out when the word is Changed
==============================*/
func (a *Octopus) ShallowStem() string {
  var str string = a.Word

  var slice []string

  var mytype *Octopus = a

  mystr := mytype.Step_1a()
  if mystr != mytype.Word {
    slice = append(slice, mystr)
  }

  var mytype2 *Octopus = a
  mystr2 := mytype2.Step_1b()
  if mystr2 != mytype2.Word {
    slice = append(slice, mystr2)
  }

  var mytype3 *Octopus = a
  mystr3 := mytype3.Step_1c()
  if mystr3 != mytype3.Word {
    slice = append(slice, mystr3)
  }

  var mytype4 *Octopus = a
  mystr4 := mytype4.Step_2()
  if mystr4 != mytype4.Word {
    slice = append(slice, mystr4)
  }

  var mytype5 *Octopus = a
  mystr5 := mytype5.Step_3()
  if mystr5 != mytype5.Word {
    slice = append(slice, mystr5)
  }

  var mytype6 *Octopus = a
  mystr6 := mytype6.Step_4()
  if mystr6 != mytype6.Word {
    slice = append(slice, mystr6)
  }

  var mytype7 *Octopus = a
  mystr7 := mytype7.Step_5a()
  if mystr7 != mytype7.Word {
    slice = append(slice, mystr7)
  }

  var mytype8 *Octopus = a
  mystr8 := mytype8.Step_5b()
  if mystr8 != mytype8.Word {
    slice = append(slice, mystr8)
  }


  if len(slice) > 0 {
    str = slice[0]
  }

  return str
}

/* Returns the Step that was used to stem the word
=========================================*/
func (a *Octopus) ShallowStemmed() string {
  var stepUsed string = "None word is a stem"

  var mytype *Octopus = a

  mystr := mytype.Step_1a()
  if mystr != mytype.Word {
    stepUsed = "Step_1a()"
    return stepUsed
  }

  mytype.Word = mystr
  mystr2 := mytype.Step_1b()
  if mystr2 != mytype.Word {
    stepUsed = "Step_1b()"
    return stepUsed
  }

  mytype.Word = mystr2
  mystr3 := mytype.Step_1c()
  if mystr3 != mytype.Word {
    stepUsed = "Step_1c()"
    return stepUsed
  }

  mytype.Word = mystr3
  mystr4 := mytype.Step_2()
  if mystr4 != mytype.Word {
    stepUsed = "Step_2()"
    return stepUsed
  }

  mytype.Word = mystr4
  mystr5 := mytype.Step_3()
  if mystr5 != mytype.Word {
    stepUsed = "Step_3()"
    return stepUsed
  }

  mytype.Word = mystr5
  mystr6 := mytype.Step_4()
  if mystr6 != mytype.Word {
    stepUsed = "Step_4()"
    return stepUsed
  }

  mytype.Word = mystr6
  mystr7 := mytype.Step_5a()
  if mystr7 != mytype.Word {
    stepUsed = "Step_5a()"
    return stepUsed
  }

  mytype.Word = mystr7
  mystr8 := mytype.Step_5b()
  if mystr8 != mytype.Word {
    stepUsed = "Step_5b()"
    return stepUsed
  }

  return stepUsed
}


/* Deep Stem makes sure that the word goes through every step in the algorithm.
   Any change made by one step is passed on to the next step
================================================================*/
func (a *Octopus) DeepStem() string {
  var str string = a.Word
  if len(a.Word) > 2 {
    var octopus *Octopus = a

    tentacle := octopus.Step_1a()

    octopus.Word = tentacle
    tentacle2 := octopus.Step_1b()

    octopus.Word = tentacle2
    tentacle3 := octopus.Step_1c()

    octopus.Word = tentacle3
    tentacle4 := octopus.Step_2()

    octopus.Word = tentacle4
    tentacle5 := octopus.Step_3()

    octopus.Word = tentacle5
    if len(octopus.Word) > 4 {
      tentacle6 := octopus.Step_4()
      octopus.Word = tentacle6
    }

    tentacle7 := octopus.Step_5a()

    octopus.Word = tentacle7
    tentacle8 := octopus.Step_5b()

    str = tentacle8
  }

  return str
}
