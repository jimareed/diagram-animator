package main

import (
	"io"
	"net/http"

	"github.com/jimareed/diagram-animator/diagram"
)

func getExamplesHandler(w http.ResponseWriter, r *http.Request) {

	d := getExamplesDiagram()

	io.WriteString(w, "<html><body>" + diagram.Diagram2Svg(d) + "</body></html>\n")
	
}
