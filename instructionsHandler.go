package main

import (
	"io"
	"net/http"

	"github.com/jimareed/diagram-animator/diagram"
)

func getInstructionsHandler(w http.ResponseWriter, r *http.Request) {

	d := getInstructionsDiagram()

	io.WriteString(w, "<html><body>" + diagram.Diagram2Svg(d) + "</body></html>\n")
	
}
