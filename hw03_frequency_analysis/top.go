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

func (sm *sortedMap) GetTopNWords(num int) []string {
	if len(sm.descSortedKeys) > num {
		return sm.descSortedKeys[:num]
	}
	return sm.descSortedKeys
}

func (sm *sortedMap) valSortedSet(words []string) {
	sm.set = make(map[string]int)
	for _, word := range words {
		sm.set[word]++
	}

	keys := make([]string, 0, len(sm.set))

	for key := range sm.set {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if sm.set[keys[i]] == sm.set[keys[j]] {
			return keys[i] < keys[j]
		}
		return sm.set[keys[i]] > sm.set[keys[j]]
	})
	sm.descSortedKeys = keys
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

	return sm.GetTopNWords(10)
}
