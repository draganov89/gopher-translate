package webservice

import (
	"log"
	"net/http"

	pt "github.com/draganov89/gopher-translate/pattern_translate"
	sh "github.com/draganov89/gopher-translate/service_handler"
)

type service struct {
	handler *sh.ServiceHandler
}

// CreateService is a constructor function that initializes
// a new service object
func CreateService(t pt.ITranslator) *service {
	if t == nil {
		log.Fatalln("Translator argument should not be nil!")
	}
	handler := sh.CreateServiceHandler(t)
	service := &service{handler}
	return service
}

// ListenAndServe starts a new service listener
func (s *service) ListenAndServe(port string) error {
	return http.ListenAndServe(port, s.handler.GetServiceMux())
}
