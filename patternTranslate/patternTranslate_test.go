package patternTranslate_test

import (
	"testing"

	gp "github.com/draganov89/gopher-translate/gopherishPattern"
	pt "github.com/draganov89/gopher-translate/patternTranslate"
)

func TestTranslator_GetSortedHistory(t *testing.T) {
	gDict := gp.GetGopherishDictionary()
	tr := pt.CreateTranslator(gDict)

	_ = tr.TranslateWord("cosmos")
	_ = tr.TranslateWord("xray")
	_ = tr.TranslateWord("apple")
	_ = tr.TranslateWord("predator")
	_ = tr.TranslateWord("world")
	_ = tr.TranslateWord("old")

	expect := []pt.Record{
		{"apple", "gapple"},
		{"cosmos", "osmoscogo"},
		{"old", "gold"},
		{"predator", "edatorprogo"},
		{"world", "gworld"},
		{"xray", "gexray"},
	}

	got := tr.GetSortedHistory()

	if len(got) != len(expect) {
		t.Error("Lenght of expected slice is different then the lenght of got slice!")
	}

	for ind := 0; ind < len(got); ind++ {
		if expect[ind] != got[ind] {
			t.Errorf("Element at position %d does not match! Expected %q, got %q", ind, expect[ind], got[ind])
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
