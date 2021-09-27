package service_handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	if transl == nil {
		log.Fatalln("Translator argument should not be nil!")
	}
	mux := http.NewServeMux()
	handler := &ServiceHandler{transl, mux}

	handler.serveMux.HandleFunc("/word", handler.TranslateWord)
	handler.serveMux.HandleFunc("/sentence", handler.TranslateSentence)
	handler.serveMux.HandleFunc("/history", handler.GetTranslationHistory)
	log.Println("Handlers registered!")

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
		log.Printf("Error with GetTranslationHistory handler: %v method not supported! Expected %v\n", r.Method, "GET")
		return
	}

	translated := s.translator.GetSortedHistory()

	bytesRes, err := json.Marshal(translated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went unexpectedly wrong! (It works on my machine...)"))
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesRes)
}

// TranslateSentence is a handler responsible for the "/sentence" endpoint.
// Returning the translated sentence in gopher language.
func (s *ServiceHandler) TranslateSentence(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Error with TranslateSentence handler: %v method not supported! Expected %v\n", r.Method, "POST")
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	var reqArgMap map[string]string

	err = json.Unmarshal(body, &reqArgMap)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	sentence, ok := reqArgMap["english_sentence"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON request!"))
		log.Println("Invalid json property!")
		return
	}

	translated, err := s.translator.TranslateSentence(sentence)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		log.Println(err.Error())
		return
	}

	respMap := map[string]string{
		"gopher_sentence": translated,
	}

	bytesRes, err := json.Marshal(respMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went unexpectedly wrong! (It works on my machine...)"))
		log.Println(err.Error())
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
		log.Printf("Error with TranslateWord handler: %v method not supported! Expected %v\n", r.Method, "POST")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	var reqArg map[string]string

	err = json.Unmarshal(body, &reqArg)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	word, ok := reqArg["english_word"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON request!"))
		log.Println("Invalid JSON property!")
		return
	}

	translated := s.translator.TranslateWord(word)

	respMap := map[string]string{
		"gopher_word": translated,
	}

	bytesRes, err := json.Marshal(respMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went unexpectedly wrong! (It works on my machine...)"))
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesRes)
}
