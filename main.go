package main

import (
	"fmt"
	"os"

	gp "github.com/draganov89/gopher-translate/gopherishPattern"
	pt "github.com/draganov89/gopher-translate/patternTranslate"
	ws "github.com/draganov89/gopher-translate/webservice"
)

func main() {

	defPort := "8899"
	var port string
	if len(os.Args) < 3 {
		fmt.Printf("Missing argument '--port'! Default port will be used :%v\n", defPort)
		port = defPort
	} else {
		port = os.Args[2]
	}
	port = fmt.Sprintf(":%s", port)

	gDict := gp.GetGopherishDictionary()
	gTranslator := pt.CreateTranslator(gDict)
	service := ws.CreateService(gTranslator)

	service.ListenAndServe(port)
}
