package main

import (
	"fmt"
	"os"

	gp "github.com/draganov89/gopher-translate/gopherishPattern"
	pt "github.com/draganov89/gopher-translate/patternTranslate"
	ws "github.com/draganov89/gopher-translate/webService"
)

func main() {

	// No need for named parameters
	// such as '--port or -p'
	port := fmt.Sprintf(":%s", os.Args[1])

	// ancient gopherish dictionary
	gDict := gp.GetGopherishDictionary()
	gTranslator := pt.CreateTranslator(gDict)
	service := ws.CreateService(gTranslator)

	service.ListenAndServe(port)
}
