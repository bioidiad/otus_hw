package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type sortedMap struct {
	set            map[string]int
	descSortedKeys []string
}

func (sm *sortedMap) valSortedSet(words []string) {
	sm.set = make(map[string]int)
	for _, word := range words {
		sm.set[word]++
	}

	Keys := make([]string, 0, len(sm.set))

	for key := range sm.set {
		Keys = append(Keys, key)
	}

	sort.Slice(Keys, func(i, j int) bool {
		return Keys[i] < Keys[j]
	})

	sort.SliceStable(Keys, func(i, j int) bool {
		return sm.set[Keys[i]] > sm.set[Keys[j]]
	})

	sm.descSortedKeys = Keys
}

var (
	splitter = regexp.MustCompile(`\s+`)
	trimmer  = regexp.MustCompile(`^[^а-яА-Яa-zA-Z0-9 ]+|[^а-яА-Яa-zA-Z0-9 ]+$`)
)

func splitWords(text string) []string {
	var result []string
	for _, word := range splitter.Split(text, -1) {
		tmp := trimmer.ReplaceAllString(word, "")
		if tmp != "" {
			result = append(result, strings.ToLower(tmp))
		}
	}

	return result
}

func Top10(text string) []string {
	if text == "" {
		return []string{}
	}

	var sm sortedMap
	words := splitWords(text)
	sm.valSortedSet(words)
	if len(sm.descSortedKeys) > 10 {
		return sm.descSortedKeys[:10]
	} else {
		return sm.descSortedKeys
	}
}
