package serviceHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pt "github.com/draganov89/gopher-translate/patternTranslate"
)

// ServiceHandler object encapsulates
// the handlers and routings functionality
type ServiceHandler struct {
	translator *pt.Translator
	serveMux   *http.ServeMux
}

// CreateServiceHandler is a constructor function that
// initializes a new ServiceHandler object
func CreateServiceHandler(transl *pt.Translator, mux *http.ServeMux) *ServiceHandler {
	handler := &ServiceHandler{transl, mux}

	handler.serveMux.HandleFunc("/word", handler.TranslateWord)
	handler.serveMux.HandleFunc("/sentence", handler.TranslateSentence)
	handler.serveMux.HandleFunc("/history", handler.GetTranslationHistory)

	return handler
}

// GetServiceMux is a getter function for the serveMux
// unexported field
func (s *ServiceHandler) GetServiceMux() *http.ServeMux {
	return s.serveMux
}

// GetTranslationHistory is a handler responsible for the "/history" endpoint.
// Returning the history of all translations in json format.
func (s *ServiceHandler) GetTranslationHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	translated := s.translator.GetSortedHistory()

	bytesRes, err := json.Marshal(translated)
	if err != nil {
		fmt.Println("Error marshaling the response object!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesRes)
}

// TranslateSentence is a handler responsible for the "/sentence" endpoint.
// Returning the translated sentence in gopher language.
func (s *ServiceHandler) TranslateSentence(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		fmt.Println("Error occured while reading request body!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var reqArg map[string]string

	err = json.Unmarshal(body, &reqArg)

	if err != nil {
		fmt.Println("Error parsing the request body!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sentence, ok := reqArg["english_sentence"]
	if !ok {
		fmt.Println("Error - request body not in the correct format!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	translated := s.translator.TranslateSentence(sentence)

	resultObj := struct {
		Gopher_sentence string `json:"gopher_sentence"`
	}{
		translated,
	}

	bytesRes, err := json.Marshal(resultObj)
	if err != nil {
		fmt.Println("Error marshaling the response object!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesRes)
}

// TranslateWord is a handler responsible for the "/word" endpoint.
// Returning the translated word in gopher language.
func (s *ServiceHandler) TranslateWord(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		fmt.Println("Error occured while reading request body!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var reqArg map[string]string

	err = json.Unmarshal(body, &reqArg)

	if err != nil {
		fmt.Println("Error parsing the request body!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	word, ok := reqArg["english_word"]
	if !ok {
		fmt.Println("Error - request body not in the correct format!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	translated := s.translator.TranslateWord(word)

	resultObj := struct {
		Gopher_word string `json:"gopher_word"`
	}{
		translated,
	}

	bytesRes, err := json.Marshal(resultObj)
	if err != nil {
		fmt.Println("Error marshaling the response object!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesRes)
}
