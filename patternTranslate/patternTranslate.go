package patternTranslate

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Record struct {
	EngWord    string
	GopherWord string
}

type Translator struct {
	patterns map[*regexp.Regexp]string
	history  []Record
}

func CreateTranslator(dictionary map[*regexp.Regexp]string) *Translator {
	t := &Translator{dictionary, make([]Record, 0)}
	return t
}

func (t *Translator) TranslateWord(englishWord string) string {

	englishWord = strings.ToLower(englishWord)
	translated := englishWord
	for k, v := range t.patterns {
		match := k.MatchString(englishWord)
		if match {
			translated = k.ReplaceAllString(englishWord, v)
			break
		}
	}
	t.addToHistory(englishWord, translated)
	return translated
}

func (t *Translator) TranslateSentence(englishSentence string) string {
	words := strings.Split(englishSentence, " ")

	var strBuilder strings.Builder
	// first word first letter should be capital
	strBuilder.WriteString(strings.Title(t.TranslateWord(words[0])))

	for i := 1; i < len(words); i++ {
		fmt.Fprintf(&strBuilder, " %s", t.TranslateWord(words[i]))
	}

	return strBuilder.String()
}

func (t *Translator) addToHistory(eng, goph string) {
	t.history = append(t.history, Record{eng, goph})
}

func (t *Translator) GetSortedHistory() []Record {
	sort.Slice(t.history, func(p, q int) bool {
		return t.history[p].EngWord < t.history[q].EngWord
	})

	result := make([]Record, len(t.history))
	copy(result, t.history)
	return result
}
