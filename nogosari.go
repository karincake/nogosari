package nogosari

import (
	"html"
	"strings"
)

// Tokenize remove symbols and URLs from s, then split it into words
func Tokenize(s string) []string {
	// Normalize s and remove all symbol
	s = strings.ToLower(s)
	s = html.UnescapeString(s)
	s = rxURL.ReplaceAllString(s, "")
	s = rxEmail.ReplaceAllString(s, "")
	s = rxTwitter.ReplaceAllString(s, "")
	s = rxEscapeStr.ReplaceAllString(s, "")
	s = rxSymbol.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)

	return strings.Fields(s)
}

// Similar to tokenize, except it returns map[string]uint for indexing purpose.
// The `uint16` value indicates the word count
func Index(s string) map[string]uint16 {
	// Normalize s and remove all symbol
	s = strings.ToLower(s)
	s = html.UnescapeString(s)
	s = rxURL.ReplaceAllString(s, "")
	s = rxEmail.ReplaceAllString(s, "")
	s = rxTwitter.ReplaceAllString(s, "")
	s = rxEscapeStr.ReplaceAllString(s, "")
	s = rxSymbol.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)

	words := make(map[string]uint16)
	found := false
	lastI := 0
	sb := []byte(s)
	for i := range sb {
		if sb[i] == ' ' {
			if !found {
				continue
			}
			word := string(sb[lastI:i])
			lastI = i + 1
			if _, ok := words[word]; !ok {
				words[word] = 1
			} else {
				words[word] = 1 + words[word]
			}
			found = false
		} else {
			if !found {
				lastI = i
			}
			found = true
		}
	}

	return words
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