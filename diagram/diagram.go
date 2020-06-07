package diagram

import (
	"fmt"
)

// Block contains the x,y coordinates of the start of the block
type Element struct {
	x, y int
	elementType string
	description string
	size int
}

type Connector struct {
	
}

type Transition struct {
	duration int
}

// Diagram contains a list of blocks (structure used by block-diagram-editor)
type Diagram struct {
	width  int
	height int
	blockWidth int
	blockHeight int
	elements []Element
	connectors []Connector
	transitions []Transition
}

func DefaultDiagram() Diagram {
	d := Diagram{}

	d.width = 900
	d.height = 600
	d.blockWidth = 90
	d.blockHeight = 60

	return d
}

func AddBlock(d Diagram, x int, y int) Diagram {

	d.elements = append(d.elements, Element{x,y,"block","",0})
	return d
}

func AddText(d Diagram, x int, y int, description string, size int) Diagram {

	d.elements = append(d.elements, Element{x,y,"text",description,size})
	return d
}

func AddTransition(d Diagram, duration int) Diagram {

	d.transitions = append(d.transitions, Transition{duration})
	return d
}

func Diagram2String(diagram Diagram) string {

	b := ""

	for _, e := range diagram.elements {
		if len(b) > 0 {
			b += ","
		}
		b += fmt.Sprintf("{\"x\": %d, \"y\": %d, \"elementType\": \"%s\"}", e.x, e.y, e.elementType)
	}

	s := fmt.Sprintf(
		"{"+
			"\"width\": %d,"+
			"\"height\": %d,"+
			"\"blockWidth\": %d,"+
			"\"blockHeight\": %d,"+
			"\"elements\": [%s],"+
			"\"connectors\": []"+
			"}\n",
		diagram.width, diagram.height, diagram.blockWidth, diagram.blockHeight, b)

	return s
}

func slope(d Diagram, i1 int, i2 int) int {
	return d.elements[i2].x - d.elements[i1].x / d.elements[i2].y - d.elements[i1].y
}

func Diagram2Svg(diagram Diagram) string {

	elements := ""
	i := 0
	for _, e := range diagram.elements {
		if e.elementType == "block" {
			elements += fmt.Sprintf("<rect class=\"transition%d\" x=\"%d\" y=\"%d\" width=\"90\" height=\"60\" id=\"1\" stroke=\"black\" fill=\"transparent\" stroke-width=\"4\"></rect>\n", i, e.x, e.y)
		} 
		if e.elementType == "text" {
			elements += fmt.Sprintf("<text class=\"transition%d\" x=\"%d\" y=\"%d\" fill=\"black\" font-size=\"%dpx\">%s</text>\n", i, e.x, e.y, e.size, e.description)
		} 
		i++
	}

	transitions := ""
	i = 0
	for _, t := range diagram.transitions {
		transitions += fmt.Sprintf(
			".transition%d {"+
			"	animation-name: transitionOpacity;"+
			"	animation-duration: %ds;"+
			"	animation-iteration-count: 1;"+
			"}", i, t.duration)
		i++
	}

	width := diagram.width
	height := diagram.height

	if width <= 0 {
		width = 600
	}
	if height <= 0 {
		height = 400
	}

	s := fmt.Sprintf(
		"<svg width=\"%d\" height=\"%d\" align=\"center\">" +
        " <rect x=\"0\" y=\"0\" id=\"editor-canvas\" width=\"%d\" height=\"%d\" stroke=\"white\" fill=\"transparent\" stroke-width=\"0\"></rect>"+
		"%s\n"+
		"<style>\n"+
		"%s\n"+
		"@keyframes transitionOpacity {\n"+
        "    0%%   { opacity: 0; }\n"+
        "    50%%   { opacity: 0; }\n"+
        "    100%% { opacity: 1; }\n"+
        "}\n"+
	    "</style>\n"+
		"</svg>\n",
		width, height, width, height, elements, transitions)

		return s
}
