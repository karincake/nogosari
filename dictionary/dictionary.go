package nogosari

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Create a dictionary from file that contains list of word and its word-
// functions.
// Note that any error occured during the conversion of word-function will
// be ignored, meaning that the word will have no word-function
func CreateFromFile(fn string, binary bool) map[string]uint16 {
	// stopwords
	f, err := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic("open file error:" + err.Error())
	}
	defer f.Close()

	words := make(map[string]uint16)
	sc := bufio.NewScanner(f)

	if binary {
		for sc.Scan() {
			s := sc.Text()
			ss := strings.Fields(s)
			val := uint16(0)
			if len(ss) > 1 {
				i, err := strconv.ParseInt(ss[1], 2, 64)
				if err != nil {
					continue
				}
				val = uint16(i)
			}
			words[ss[0]] = val
		}
	} else {
		for sc.Scan() {
			s := sc.Text()
			ss := strings.Fields(s)
			val := uint16(0)
			if len(ss) > 1 {
				xVal, err := strconv.Atoi(ss[1])
				if err != nil {
					continue
				}
				val = uint16(xVal)
			}
			words[ss[0]] = val
		}
	}

	if err := sc.Err(); err != nil {
		panic("scan file error:" + err.Error())
	}

	return words
}

// Similar to CreateFromFile, but without word-function information
func CreatePlainFromFile(fn string) map[string]any {
	// stopwords
	f, err := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic("open file error:" + err.Error())
	}
	defer f.Close()

	words := make(map[string]any)
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		s := sc.Text()
		words[s] = struct{}{}
	}

	if err := sc.Err(); err != nil {
		panic("scan file error:" + err.Error())
	}

	return words
}
