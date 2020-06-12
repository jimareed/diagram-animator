package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/jimareed/diagram-animator/diagram"

)

var mainDiagram = diagram.Diagram{}

func getMainDiagram() diagram.Diagram {

	mainDiagram = diagram.DefaultDiagram()
	mainDiagram = diagram.AddBlock(mainDiagram, 16, 170)
	mainDiagram = diagram.AddBlock(mainDiagram, 233, 121)
	mainDiagram = diagram.AddBlock(mainDiagram, 510, 248)
	mainDiagram = diagram.AddText(mainDiagram, 190, 95, "Diagram", 24)
	mainDiagram = diagram.AddText(mainDiagram, 280, 95, "Animator", 24)
	mainDiagram = diagram.AddText(mainDiagram, 220, 240, "Add animation", 20)
	mainDiagram = diagram.AddText(mainDiagram, 130, 280, "to help communicate your message", 20)
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

	mainDiagram = diagram.AddConnector(mainDiagram, 0, 1)
	mainDiagram = diagram.AddConnector(mainDiagram, 1, 2)

	return mainDiagram
}

func getInstructionsDiagram() diagram.Diagram {
	
	mainDiagram = diagram.DefaultDiagram()
	mainDiagram = diagram.AddText(mainDiagram, 110, 95, "Instructions:", 24)
	mainDiagram = diagram.AddText(mainDiagram, 110, 180, "1. Get block diagram from Block Diagram Editor", 20)
	mainDiagram = diagram.AddText(mainDiagram, 110, 220, "2. Use PostMan to add the diagram", 20)
	mainDiagram = diagram.AddText(mainDiagram, 110, 260, "3. Use PostMan to add transitions", 20)
	mainDiagram = diagram.AddText(mainDiagram, 130, 300, "(refer to readme for more details)", 20)

	mainDiagram = diagram.AddTransition(mainDiagram, 2)
	mainDiagram = diagram.AddTransition(mainDiagram, 5)
	mainDiagram = diagram.AddTransition(mainDiagram, 10)
	mainDiagram = diagram.AddTransition(mainDiagram, 15)
	mainDiagram = diagram.AddTransition(mainDiagram, 20)

	return mainDiagram
}

func getExamplesDiagram() diagram.Diagram {
	
	mainDiagram = diagram.DefaultDiagram()
	mainDiagram = diagram.AddBlock(mainDiagram, 16, 170)
	mainDiagram = diagram.AddBlock(mainDiagram, 233, 121)
	mainDiagram = diagram.AddBlock(mainDiagram, 510, 248)
	mainDiagram = diagram.AddText(mainDiagram, 190, 95, "Examples", 24)
	mainDiagram = diagram.AddText(mainDiagram, 220, 240, "Add animation", 20)
	mainDiagram = diagram.AddText(mainDiagram, 130, 280, "to help communicate your message", 20)
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

	mainDiagram = diagram.AddConnector(mainDiagram, 0, 1)
	mainDiagram = diagram.AddConnector(mainDiagram, 1, 2)

	return mainDiagram
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getAnimationHandler).Methods("GET")
	r.HandleFunc("/instructions", getInstructionsHandler).Methods("GET")
	r.HandleFunc("/examples", getExamplesHandler).Methods("GET")
	r.HandleFunc("/block-diagram", getBlockDiagramHandler).Methods("GET")
	r.HandleFunc("/diagram", getDiagramHandler).Methods("GET")

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
