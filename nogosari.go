package nogosari

import (
	"strings"
)

// Tokenize remove symbols and URLs from s, then split it into words
func Tokenize(s string) []string {
	// Normalize s and remove all symbol
	s = normalize(s)

	return strings.Fields(s)
}

// Full version of tokenize which includes symbol removal
func FullTokenize(s string) []string {
	// Normalize s and remove all symbol
	s = normalize(s)
	s = rxSymbol.ReplaceAllString(s, " ")

	return strings.Fields(s)
}

// Similar to tokenize, except it excludes symbols and returns map[string]uint
// for indexing purpose.
// The `uint16` value indicates the word count
func Index(s string) map[string]uint16 {
	// Normalize s and remove all symbol
	s = normalize(s)

	return index(s)
}

// Full version of index which includes symbol removal
func FullIndex(s string) map[string]uint16 {
	// Normalize s and remove all symbol
	s = normalize(s)
	s = rxSymbol.ReplaceAllString(s, " ")

	return index(s)
}

// Convert a word back to its base form based on the given dictionary
func Stem(word string, words map[string]uint16) string {
	word = strings.ToLower(word)

	var (
		rootFound    = false
		originalWord = word
		particle     string
		possesive    string
		suffix       string
	)

	if len(word) < 3 {
		return word
	}

	if _, ok := words[word]; ok {
		return word
	}

	// Check if prefix must be removed first
	if rxPrefixFirst.MatchString(word) {
		// Remove prefix
		rootFound, word = removePrefixes(word, words)
		if rootFound {
			return word
		}

		// Remove particle
		particle, word = removeParticle(word)
		if _, ok := words[word]; ok {
			return word
		}

		// Remove possesive
		possesive, word = removePossesive(word)
		if _, ok := words[word]; ok {
			return word
		}

		// Remove suffix
		suffix, word = removeSuffix(word)
		if _, ok := words[word]; ok {
			return word
		}
	} else {
		// Remove particle
		particle, word = removeParticle(word)
		if _, ok := words[word]; ok {
			return word
		}

		// Remove possesive
		possesive, word = removePossesive(word)
		if _, ok := words[word]; ok {
			return word
		}

		// Remove suffix
		suffix, word = removeSuffix(word)
		if _, ok := words[word]; ok {
			return word
		}

		// Remove prefix
		rootFound, word = removePrefixes(word, words)
		if rootFound {
			return word
		}
	}

	// If no root found, do loopPengembalianAkhiran
	removedSuffixes := []string{"", suffix, possesive, particle}
	if suffix == "kan" {
		removedSuffixes = []string{"", "k", "an", possesive, particle}
	}

	rootFound, word = loopPengembalianAkhiran(originalWord, removedSuffixes, words)
	if rootFound {
		return word
	}

	// When EVERYTHING failed, return original word
	return originalWord
}
