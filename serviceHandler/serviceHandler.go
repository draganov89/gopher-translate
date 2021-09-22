package serviceHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pt "github.com/draganov89/gopher-translate/patternTranslate"
)

type ServiceHandler struct {
	translator *pt.Translator
	serveMux   *http.ServeMux
}

func CreateServiceHandler(transl *pt.Translator, mux *http.ServeMux) *ServiceHandler {
	handler := &ServiceHandler{transl, mux}

	handler.serveMux.HandleFunc("/word", handler.TranslateWord)
	handler.serveMux.HandleFunc("/sentence", handler.TranslateSentence)
	handler.serveMux.HandleFunc("/history", handler.GetTranslationHistory)

	return handler
}

func (s *ServiceHandler) GetServiceMux() *http.ServeMux {
	return s.serveMux
}

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
