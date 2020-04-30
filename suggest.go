package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/xyproto/vt100"
)

func corpus(searchword, glob string) []string {
	wordCount := make(map[string]int)

	filenames, err := filepath.Glob(glob)
	if err != nil {
		return []string{}
	}

	re := regexp.MustCompile(`[[:^alpha:]]` + searchword + `\.([[:alpha:]]*)`)

	var data []byte
	var highestCount int
	for _, filename := range filenames {
		data, err = ioutil.ReadFile(filename)
		if err != nil {
			continue
		}
		submatches := re.FindAllStringSubmatch(string(data), -1)
		for _, submatch := range submatches {
			word := submatch[1]
			if _, ok := wordCount[word]; ok {
				wordCount[word]++
				if wordCount[word] > highestCount {
					highestCount = wordCount[word]
				}
			} else {
				wordCount[word] = 1
				if wordCount[word] > highestCount {
					highestCount = wordCount[word]
				}
			}
		}
	}

	// Copy the words from the map to a string slice, such
	// that the most frequent words appear first.
	sl := make([]string, len(wordCount), len(wordCount))
	slIndex := 0
	for i := highestCount; i >= 0; i-- {
		for word, count := range wordCount {
			if count == i && len(word) > 0 {
				sl[slIndex] = word
				slIndex++
			}
		}
	}

	return sl
}

// SuggestMode lets the user tab through the suggested words
func (e *Editor) SuggestMode(c *vt100.Canvas, status *StatusBar, tty *vt100.TTY, suggestions []string) string {
	suggestIndex := 0
	if len(suggestions) == 0 {
		return ""
	}
	status.ClearAll(c)
	s := suggestions[suggestIndex]
	suggestIndex++
	status.SetMessage("Suggest: " + s)
	status.ShowNoTimeout(c, e)
	doneCollectingLetters := false
	for !doneCollectingLetters {
		key := tty.String()
		switch key {
		case "c:9": // tab
			// Cycle suggested words
			if suggestIndex == len(suggestions) {
				suggestIndex = 0
			}
			s = suggestions[suggestIndex]
			suggestIndex++
			status.ClearAll(c)
			status.SetMessage("Suggest: " + s)
			status.ShowNoTimeout(c, e)
		case "c:27", "c:17": // esc or ctrl-q
			s = ""
			fallthrough
		case "c:13", "c:32": // return or space
			doneCollectingLetters = true
		}
	}
	status.ClearAll(c)
	// The chosen word
	return s
}