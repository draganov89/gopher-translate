package webservice

import (
	"net/http"

	pt "github.com/draganov89/gopher-translate/patternTranslate"
)

type service struct {
	handler *ServiceHandler
}

func CreateService(t *pt.Translator) *service {

	newMux := http.NewServeMux()

	handler := &ServiceHandler{t, newMux}

	handler.serveMux.HandleFunc("/word", handler.translateWord)
	handler.serveMux.HandleFunc("/sentence", handler.translateSentence)

	service := &service{handler}

	return service
}

func (s *service) ListenAndServe(port string) {
	http.ListenAndServe(port, s.handler.serveMux)
}
