package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/jimareed/diagram-animator/diagram"

)

var mainDiagram = diagram.Diagram{}

func getMainDiagram() diagram.Diagram {
	return mainDiagram
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getAnimationHandler).Methods("GET")
	r.HandleFunc("/block-diagram", getBlockDiagramHandler).Methods("GET")
	r.HandleFunc("/diagram", getDiagramHandler).Methods("GET")

	mainDiagram = diagram.DefaultDiagram()
	mainDiagram = diagram.AddBlock(mainDiagram, 16, 170)
	mainDiagram = diagram.AddBlock(mainDiagram, 233, 121)
	mainDiagram = diagram.AddBlock(mainDiagram, 510, 248)
	mainDiagram = diagram.AddText(mainDiagram, 190, 95, "Diagram", 24)
	mainDiagram = diagram.AddText(mainDiagram, 280, 95, "Animator", 24)
	mainDiagram = diagram.AddText(mainDiagram, 220, 240, "Add animation", 20)
	mainDiagram = diagram.AddText(mainDiagram, 48, 280, "to more effectively communicate your message", 20)
	mainDiagram = diagram.AddText(mainDiagram, 122, 390, "Instructions", 20)
	mainDiagram = diagram.AddText(mainDiagram, 348, 390, "Examples", 20)

	mainDiagram = diagram.AddTransition(mainDiagram, 2)
	mainDiagram = diagram.AddTransition(mainDiagram, 5)
	mainDiagram = diagram.AddTransition(mainDiagram, 10)
	mainDiagram = diagram.AddTransition(mainDiagram, 12)
	mainDiagram = diagram.AddTransition(mainDiagram, 14)
	mainDiagram = diagram.AddTransition(mainDiagram, 18)
	mainDiagram = diagram.AddTransition(mainDiagram, 22)
	mainDiagram = diagram.AddTransition(mainDiagram, 28)
	mainDiagram = diagram.AddTransition(mainDiagram, 32)

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
