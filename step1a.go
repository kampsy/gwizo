/*Package gwizo implement Porter, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/
*/
package gwizo

import "strings"

/*Step1a from "An algorithm for suffix stripping".
  Deals with plurals and past participles. The subsequent steps
  are much more straightforward.

  From the paper:

      SSES -> SS                         caresses  ->  caress
      IES  -> I                          ponies    ->  poni
                                         ties      ->  ti
      SS   -> SS                         caress    ->  caress
      S    ->                            cats      ->  cat
*/
func Step1a(word string) string {
	// For SSES suffix. SSES -> SS
	sses := strings.HasSuffix(word, "sses")
	if sses {
		pre := strings.TrimSuffix(word, "sses")
		word = pre + "ss"
		return word
	}

	// For IES suffix. IES -> I
	ies := strings.HasSuffix(word, "ies")
	if ies {
		pre := strings.TrimSuffix(word, "ies")
		word = pre + letterI
		return word
	}

	// For SS suffix. SS -> SS
	ss := strings.HasSuffix(word, "ss")
	if ss {
		pre := strings.TrimSuffix(word, "ss")
		word = pre + "ss"
		return word
	}

	// For S suffix. S ->
	s := strings.HasSuffix(word, "s")
	if s {
		pre := strings.TrimSuffix(word, "s")
		word = pre
		return word
	}
	return word
}
