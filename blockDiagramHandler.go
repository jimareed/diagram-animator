package main

import (
	"io"
	"net/http"

	"github.com/jimareed/diagram-animator/diagram"

)

func getBlockDiagramHandler(w http.ResponseWriter, r *http.Request) {

	d := diagram.BlockDiagram{}

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, diagram.BlockDiagram2String(d))
	
}

