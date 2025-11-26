package levels_final

import (
	"fmt"
	"sort"
	"strings"
)

func AnalyzeText(text string) {
	sep := func(r rune) bool {
		return r == ' ' || r == ',' || r == '.' || r == '!' || r == '?'
	}

	words := strings.FieldsFunc(text, sep)

	countWords := len(words)

	set := make(map[string]int)
	for i := range words {
		set[strings.ToLower(words[i])]++
	}

	countUniquewords := len(set)

	mostPopularWord := ""
	countMostPopularWord := 0
	for key, val := range set {
		if val > countMostPopularWord {
			countMostPopularWord = val
			mostPopularWord = key
		}
	}

	lenTop5 := make([]int, 5)
	top5Words := getTopWords(set, 5)
	for i := 0; i < len(top5Words); i++ {
		lenTop5[i] = set[top5Words[i]]
	}

	fmt.Printf("Количество слов: %d\n", countWords)
	fmt.Printf("Количество уникальных слов: %d\n", countUniquewords)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", mostPopularWord, countMostPopularWord)
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for i := range 5 {
		fmt.Printf("\"%s\": %d раз\n", top5Words[i], lenTop5[i])
	}
}

func getTopWords(wordMap map[string]int, n int) []string {
	ans := make([]string, n)
	counts := make([]int, len(wordMap))
	j := 0
	for _, v := range wordMap {
		counts[j] = v
		j++
	}
	sort.Ints(counts)
	countsTop := counts[len(wordMap)-n:]
	for i := 0; i < len(countsTop); i++ {
		for key, value := range wordMap {
			if value != countsTop[i] {
				continue
			} else {
				ans[n-1] = key
				n--
				break
			}
		}
	}
	return ans
}
