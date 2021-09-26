package pattern_translate_test

import (
	"testing"

	gp "github.com/draganov89/gopher-translate/gopherish_pattern"
	pt "github.com/draganov89/gopher-translate/pattern_translate"
)

func TestTranslator_GetSortedHistory(t *testing.T) {
	gDict := gp.GetGopherishDictionary()
	tr := pt.CreateTranslator(gDict)

	words := []string{"cosmos", "xray", "apple", "predator", "world", "old"}

	for _, word := range words {
		_ = tr.TranslateWord(word)
	}

	expect := []struct {
		eng    string
		gopher string
	}{
		{"apple", "gapple"},
		{"cosmos", "osmoscogo"},
		{"old", "gold"},
		{"predator", "edatorprogo"},
		{"world", "gworld"},
		{"xray", "gexray"},
	}

	hist := tr.GetSortedHistory()

	if len(hist.History) != len(expect) {
		t.Errorf("Lenght of expected slice is different then the lenght of got slice! Expected %v, got %v", len(expect), len(hist.History))
		t.FailNow()
	}

	for ind, word := range expect {
		for k, v := range hist.History[ind] {
			if word.eng != k || word.gopher != v {
				t.Errorf("Element at position %d does not match! Expected {%q: %q}, got {%q: %q}", ind, word.eng, word.gopher, k, v)
			}
		}
	}
}

func TestTranslator_TranslateWord(t *testing.T) {

	gDict := gp.GetGopherishDictionary()
	tr := pt.CreateTranslator(gDict)

	tests := []struct {
		name string
		word string
		want string
	}{
		{"Start with vowel", "apple", "gapple"},
		{"Start with 'xr'", "xray", "gexray"},
		{"Start with consonants", "slim", "imslogo"},
		{"Single vowel", "a", "ga"},
		{"Start with vowel, end with sp. char", "ant.", "gant."},
		{"Start with cons., end with sp. char", "spear:", "earspogo:"},
		{"Start with xr., end with sp. char", "xray!", "gexray!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tr.TranslateWord(tt.word)
			if got != tt.want {
				t.Errorf("Translator.TranslateWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslator_TranslateSentence(t *testing.T) {

	gDict := gp.GetGopherishDictionary()
	tr := pt.CreateTranslator(gDict)

	tests := []struct {
		name     string
		sentence string
		want     string
	}{
		{"Sentence 1", "Hello my friend...", "Ellohogo ymogo iendfrogo..."},
		{"Sentence 2", "Be careful gopher!", "Ebogo arefulcogo ophergogo!"},
		{"Sentence 3", "Xray is xray, you just stand still!", "Gexray gis gexray, gyou ustjogo andstogo illstogo!"},
		{"One word Sentence", "Hello!", "Ellohogo!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tr.TranslateSentence(tt.sentence)
			if got != tt.want {
				t.Errorf("Translator.TranslateSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
