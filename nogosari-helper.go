package nogosari

import (
	"html"
	"strings"
)

// shorten normalize code
func normalize(s string) string {
	s = strings.ToLower(s)
	s = html.UnescapeString(s)
	s = rxURL.ReplaceAllString(s, "")
	s = rxEmail.ReplaceAllString(s, "")
	s = rxTwitter.ReplaceAllString(s, "")
	s = rxEscapeStr.ReplaceAllString(s, "")
	s = rxPeriod.ReplaceAllString(s, " ") // replaced with space instead of empty string
	s = strings.TrimSpace(s)

	return s
}

// shorten indexing code
func index(s string) map[string]uint16 {
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
	if len(sb)-lastI > 0 {
		words[string(sb[lastI:])] = 1
	}

	return words
}
