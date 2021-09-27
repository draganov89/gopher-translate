package main

import (
	"fmt"
	"log"
	"os"

	gp "github.com/draganov89/gopher-translate/gopherish_pattern"
	pt "github.com/draganov89/gopher-translate/pattern_translate"
	ws "github.com/draganov89/gopher-translate/webservice"
)

func main() {

	defPort := "8899"
	var port string
	if len(os.Args) < 3 {
		log.Fatalf("Missing argument '--port'! Default port will be used :%v\n", defPort)
		port = defPort
	} else {
		port = os.Args[2]
	}
	port = fmt.Sprintf(":%s", port)

	gDict := gp.GetGopherishDictionary()
	gTranslator := pt.CreateTranslator(gDict)
	service := ws.CreateService(gTranslator)

	log.Println("Starting service ...")
	log.Fatal(service.ListenAndServe(port))
}
