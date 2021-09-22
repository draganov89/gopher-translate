package webservice

import (
	"net/http"

	pt "github.com/draganov89/gopher-translate/patternTranslate"
	sh "github.com/draganov89/gopher-translate/serviceHandler"
)

type service struct {
	handler *sh.ServiceHandler
}

// CreateService is a constructo function that initializes
// a new service object
func CreateService(t *pt.Translator) *service {

	newMux := http.NewServeMux()
	handler := sh.CreateServiceHandler(t, newMux)

	service := &service{handler}

	return service
}

// ListenAndServe starts a new service listener
func (s *service) ListenAndServe(port string) {
	http.ListenAndServe(port, s.handler.GetServiceMux())
}
