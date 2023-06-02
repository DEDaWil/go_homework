package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type KeyValue struct {
	key   string
	value int
}

func Top10(text string) []string {
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[1-9a-zа-яё]+(?:-?[а-яё]+)*`)
	words := re.FindAllString(text, len(text)+1)

	mapWords := map[string]int{}
	for _, word := range words {
		mapWords[word]++
	}

	sliceWords := make([]KeyValue, 0)
	for key, word := range mapWords {
		sliceWords = append(sliceWords, KeyValue{key: key, value: word})
	}

	sort.Slice(sliceWords, func(i, j int) bool {
		if sliceWords[i].value != sliceWords[j].value {
			return sliceWords[i].value > sliceWords[j].value
		}
		return sliceWords[i].key < sliceWords[j].key
	})

	sliceLen := len(sliceWords)
	if sliceLen > 10 {
		sliceLen = 10
	}

	result := make([]string, 0)
	for i := 0; i < sliceLen; i++ {
		result = append(result, sliceWords[i].key)
	}

	return result
}
