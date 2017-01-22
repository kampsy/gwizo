/*Package gwizo implements Porter Stemmer algorithm, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/
*/
package gwizo

import (
	"fmt"
	"strings"
)

// list of vowels and y.
const (
	letterY = "y"
	letterA = "a"
	letterE = "e"
	letterI = "i"
	letterO = "o"
	letterU = "u"
)

/*Token collects Tokenmation about the word to be stemmed.
  A consonant is defined in the paper as follows:
           A consonant in a word is a letter other than A, E, I, O or
           U, and other than Y preceded by a consonant. (The fact that
           the term `consonant' is defined to some extent in terms of
           itself does not make it ambiguous.) So in TOY the consonants
           are T and Y, and in SYZYGY they are S, Z and G. If a letter
           is not a consonant it is a vowel.


  From the paper:
           A consonant will be denoted by c, a vowel by v. A list
           ccc... of length greater than 0 will be denoted by C, and a
           list vvv... of length greater than 0 will be denoted by V.
           Any word, or part of a word, therefore has one of the four
           forms:
               CVCV ... C
               CVCV ... V
               VCVC ... C
               VCVC ... V

           These may all be represented by the single form

               [C]VCVC ... [V]

           where the square brackets denote arbitrary presence of their
           contents. Using (VC){m} to denote VC repeated m times, this
           may again be written as
               [C](VC){m}[V].
           m will be called the \measure\ of any word or word part when
           represented in this form. The case m = 0 covers the null
           word. Here are some examples:
               m=0    TR,  EE,  TREE,  Y,  BY.
               m=1    TROUBLE,  OATS,  TREES,  IVY.
               m=2    TROUBLES,  PRIVATE,  OATEN,  ORRERY.
*/
type Token struct {
	VowCon  string // example vcvcvc. Where v = vowel and c = consonant.
	Measure int    // Number of times the rune pair vc appears in the word.
}

/*Parse assigns the Word, VowCon and Measure filds for the word w.
 */
func Parse(w string) Token {
	// Collection of vowels and consonants
	var collection []string
	// Change the word to lowercase letters.
	wordLower := strings.ToLower(w)
	for num := 0; num < len(wordLower); num++ {
		// Check if y is the first letter of the word, if true then y is a vowel.
		// Check if the vowels a, e, i, o, u are at index[0] of the word
		if num == 0 {
			if string(wordLower[num]) == letterY || string(wordLower[num]) == letterA ||
				string(wordLower[num]) == letterE || string(wordLower[num]) == letterI ||
				string(wordLower[num]) == letterO || string(wordLower[num]) == letterU {
				collection = append(collection, "v")
			} else {
				collection = append(collection, "c")
			}
			continue
		}
		// If Y is preceded by a vowel Y becomes a consonant and if Y is preceded
		// by a consonant Y becomes a vowel.
		if collection[num-1] == "v" && string(wordLower[num]) == letterY {
			collection = append(collection, "c")
			continue
		} else if collection[num-1] == "c" && string(wordLower[num]) == letterY {
			collection = append(collection, "v")
			continue
		}

		if string(wordLower[num]) == letterA || string(wordLower[num]) == letterE ||
			string(wordLower[num]) == letterI || string(wordLower[num]) == letterO ||
			string(wordLower[num]) == letterU {
			collection = append(collection, "v")
		} else {
			collection = append(collection, "c")
		}
	}
	// make a pair of vowels and consonants eg vcvcvc
	pair := strings.Join(collection, "")
	var token Token
	token.VowCon = pair
	token.Measure = strings.Count(pair, "vc")
	return token
}

// implementation of String Method and so Stringer interface
func (token *Token) String() string {
	return fmt.Sprintf("%s %d", token.VowCon, token.Measure)
}

// HasVowel returns bool of (*v*)
func HasVowel(word string) bool {
	token := Parse(word)
	return strings.Contains(token.VowCon, "v")
}

// HasConsonant returns bool of (*c*)
func HasConsonant(word string) bool {
	token := Parse(word)
	return strings.Contains(token.VowCon, "c")
}

// MeasureNum return the measure int
func MeasureNum(word string) int {
	token := Parse(word)
	return token.Measure
}

// MeasureGreaterThan0 checks if measure value is grater than 0
func MeasureGreaterThan0(word string) bool {
	token := Parse(word)
	if token.Measure > 0 {
		return true
	}
	return false
}

// MeasureEqualTo1 checks if measure value == 1
func MeasureEqualTo1(word string) bool {
	token := Parse(word)
	if token.Measure == 1 {
		return true
	}
	return false
}

// MeasureGreaterThan1 checks if measure value is grater than 1
func MeasureGreaterThan1(word string) bool {
	token := Parse(word)
	if token.Measure > 1 {
		return true
	}
	return false
}

// HasEndst checks if word has suffix S or T
func HasEndst(word string) bool {
	s := strings.HasSuffix(word, "s")
	t := strings.HasSuffix(word, "t")

	if s == true || t == true {
		return true
	}
	return false
}

// HasEndl checks if word has suffix L
func HasEndl(word string) bool {
	l := strings.HasSuffix(word, "l")
	if l == true {
		return true
	}
	return false
}

// HascvcEndLastNotwxy checks if VowCon pattern ends with cvc, where second
// c is not W, X, Y
func HascvcEndLastNotwxy(word string) bool {
	token := Parse(word)
	cvc := strings.HasSuffix(token.VowCon, "cvc")
	wordLen := len(word)
	lastLetter := string(word[(wordLen - 1)])
	w := strings.Contains(lastLetter, "w")
	x := strings.Contains(lastLetter, "x")
	y := strings.Contains(lastLetter, letterY)

	if cvc == true && w == false && x == false && y == false {
		return true
	}
	return false
}

// HasSameDoubleConsonant checks if the word's suffix has a
// double consonant "cc"
func HasSameDoubleConsonant(word string) bool {
	token := Parse(word)
	cc := strings.HasSuffix(token.VowCon, "cc")
	if cc == true {
		wordLen := (len(word) - 1)
		letter := string(word[wordLen])
		letter2 := string(word[(wordLen - 1)])
		if letter == letter2 {
			return true
		}
	}
	return false
}
