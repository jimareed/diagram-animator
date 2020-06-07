package main

import (
	"io"
	"net/http"

	"github.com/jimareed/diagram-animator/diagram"
)

func getAnimationHandler(w http.ResponseWriter, r *http.Request) {

	d := getMainDiagram()

	io.WriteString(w, "<html><body>" + diagram.Diagram2Svg(d) + "</body></html>\n")
	
}

