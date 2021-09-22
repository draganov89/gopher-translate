package patternTranslate

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Translator struct {
	patterns map[*regexp.Regexp]string
	searches map[string]string
	histKeys []string
}

type history struct {
	History []map[string]string `json:"history"`
}

func CreateTranslator(dictionary map[*regexp.Regexp]string) *Translator {
	t := &Translator{dictionary, map[string]string{}, []string{}}
	return t
}

func (t *Translator) TranslateWord(englishWord string) string {
	translated := t.translateWord(englishWord)
	t.addToHistory(englishWord, translated)
	return translated
}

func (t *Translator) translateWord(englishWord string) string {
	englishWord = strings.ToLower(englishWord)

	translated := englishWord
	for k, v := range t.patterns {
		match := k.MatchString(englishWord)
		if match {
			translated = k.ReplaceAllString(englishWord, v)
			break
		}
	}
	return translated
}

func (t *Translator) TranslateSentence(englishSentence string) string {
	words := strings.Split(englishSentence, " ")

	var strBuilder strings.Builder

	// first letter of first word should be capital
	strBuilder.WriteString(strings.Title(t.translateWord(words[0])))

	for i := 1; i < len(words); i++ {
		fmt.Fprintf(&strBuilder, " %s", t.translateWord(words[i]))
	}

	translated := strBuilder.String()
	t.addToHistory(englishSentence, translated)
	return translated
}

func (t *Translator) addToHistory(eng, goph string) {
	t.histKeys = append(t.histKeys, eng)
	t.searches[eng] = goph
}

func (t *Translator) GetSortedHistory() *history {

	sort.Slice(t.histKeys, func(p, q int) bool {
		return t.histKeys[p] < t.histKeys[q]
	})

	hist := &history{make([]map[string]string, 0, len(t.histKeys))}

	for _, key := range t.histKeys {
		newMap := map[string]string{
			key: t.searches[key],
		}
		hist.History = append(hist.History, newMap)
	}

	return hist
}
