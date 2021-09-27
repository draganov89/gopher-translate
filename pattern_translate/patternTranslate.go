package pattern_translate

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type ITranslator interface {
	TranslateWord(string) string
	TranslateSentence(string) (string, error)
	GetSortedHistory() *History
}

// Translator encapsulates the
// translation functionality
type Translator struct {
	patterns map[*regexp.Regexp]string
	history  *HistoryHandler
}

type History struct {
	History []map[string]string `json:"history"`
}

// CreateTranslator is a constructor func that initializes
// a new Translator object
func CreateTranslator(dictionary map[*regexp.Regexp]string) *Translator {
	if dictionary == nil {
		log.Fatalln("Dictionary argument of CreateTranslator can not be nil!")
	}
	t := &Translator{
		dictionary,
		&HistoryHandler{
			translations: map[string]string{},
			keys:         []string{},
		}}
	return t
}

// TranslateWord gets an english word and returns the
// word translated in gopher
func (t *Translator) TranslateWord(englishWord string) string {
	translated := t.translateWord(englishWord)
	t.history.addToHistory(englishWord, translated)
	return translated
}

// TranslateSentence gets an english sentence and
// returns the sentence translated in gopher
func (t *Translator) TranslateSentence(englishSentence string) (string, error) {
	words := strings.Split(englishSentence, " ")

	var strBuilder strings.Builder

	// first letter of first word should be capital
	_, err := strBuilder.WriteString(strings.Title(t.translateWord(words[0])))

	if err != nil {
		return "", err
	}

	for i := 1; i < len(words); i++ {
		fmt.Fprintf(&strBuilder, " %s", t.translateWord(words[i]))
	}

	translated := strBuilder.String()
	t.history.addToHistory(englishSentence, translated)
	return translated, nil
}

// GetSortedHistory returns a history object that
// represents the sorted history of all translations
func (t *Translator) GetSortedHistory() *History {

	hist := &History{make([]map[string]string, 0, len(t.history.keys))}
	t.history.sortKeys()

	for _, key := range t.history.keys {
		newMap := map[string]string{
			key: t.history.translations[key],
		}
		hist.History = append(hist.History, newMap)
	}

	return hist
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
