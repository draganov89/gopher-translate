package serviceHandler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	pt "github.com/draganov89/gopher-translate/patternTranslate"
	sh "github.com/draganov89/gopher-translate/serviceHandler"
)

type TranslatorMock struct{}

func (tr *TranslatorMock) TranslateSentence(sentence string) string {
	if sentence == "Apples are green!" {
		return "Gapples gare eengrogo!"
	}
	return sentence
}

func (tr *TranslatorMock) TranslateWord(word string) string {
	if word == "apple" {
		return "gapple"
	}
	return word
}

func (tr *TranslatorMock) GetSortedHistory() *pt.History {
	return &pt.History{
		History: []map[string]string{
			{
				"Apples are green!": "Gapples gare eengrogo!",
			},
			{
				"apple": "gapple",
			},
		},
	}
}

func TestServiceHandler_TranslateWord(t *testing.T) {
	translator := &TranslatorMock{}
	handler := sh.CreateServiceHandler(translator)

	// ============== TEST INVALID HTTP METHOD =================
	req := httptest.NewRequest("GET", "/word", nil)
	rw := httptest.NewRecorder()

	handler.TranslateWord(rw, req)
	expected := http.StatusNotFound
	got := rw.Code
	if got != expected {
		t.Errorf("Unexpected status code for invalid http method! Expected %v, got %v!", expected, got)
	}

	// ============== TEST NIL BODY =================
	req = httptest.NewRequest("POST", "/word", nil)
	rw = httptest.NewRecorder()

	handler.TranslateWord(rw, req)
	expected = http.StatusBadRequest
	got = rw.Code
	if got != expected {
		t.Errorf("Unexpected status code for nil req body! Expected %v, got %v!", expected, got)
	}

	// ============== TEST UNMARSHAL FAIL =================
	req = httptest.NewRequest("POST", "/word", strings.NewReader("invalid string"))
	rw = httptest.NewRecorder()

	handler.TranslateWord(rw, req)
	expected = http.StatusBadRequest
	got = rw.Code
	if got != expected {
		t.Errorf("Unexpected status code for invalid req body! Expected %v, got %v!", expected, got)
	}

	// ============== TEST INVALID REQUEST BODY =================
	req = httptest.NewRequest("POST", "/word", strings.NewReader(`{"english_breakfast": "apple"}`))
	rw = httptest.NewRecorder()

	handler.TranslateWord(rw, req)
	expected = http.StatusBadRequest
	got = rw.Code
	if got != expected {
		t.Errorf("Unexpected status code for invalid req body! Expected %v, got %v!", expected, got)
	}

}