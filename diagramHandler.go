package main

import (
	"io"
	"net/http"

	"github.com/jimareed/diagram-animator/diagram"

)

func getDiagramHandler(w http.ResponseWriter, r *http.Request) {

	d := getMainDiagram()

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, diagram.Diagram2String(d))
	
}

