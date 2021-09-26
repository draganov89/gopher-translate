package service_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pt "github.com/draganov89/gopher-translate/pattern_translate"
)

// ServiceHandler object encapsulates
// the handlers and routings functionality
type ServiceHandler struct {
	translator pt.ITranslator
	serveMux   *http.ServeMux
}

// CreateServiceHandler is a constructor function that
// initializes a new ServiceHandler object
func CreateServiceHandler(transl pt.ITranslator) *ServiceHandler {
	mux := http.NewServeMux()
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
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went unexpectedly wrong! (It works on my machine...)"))
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
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var reqArgMap map[string]string

	err = json.Unmarshal(body, &reqArgMap)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	sentence, ok := reqArgMap["english_sentence"]
	if !ok {
		fmt.Printf("Invalid json property!")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON request!"))
		return
	}

	translated := s.translator.TranslateSentence(sentence)

	respMap := map[string]string{
		"gopher_sentence": translated,
	}

	bytesRes, err := json.Marshal(respMap)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went unexpectedly wrong! (It works on my machine...)"))
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
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var reqArg map[string]string

	err = json.Unmarshal(body, &reqArg)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	word, ok := reqArg["english_word"]
	if !ok {
		fmt.Printf("Invalid json property!")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON request!"))
		return
	}

	translated := s.translator.TranslateWord(word)

	respMap := map[string]string{
		"gopher_word": translated,
	}

	bytesRes, err := json.Marshal(respMap)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went unexpectedly wrong! (It works on my machine...)"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesRes)
}
