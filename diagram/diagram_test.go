package diagram

import (
	"testing"
)

func TestEmptyDiagram(t *testing.T) {
	diagram := Diagram{}
	diagram.width = 0
	diagram.height = 0
	diagram.blockWidth = 0
	diagram.blockHeight = 0

	if len(diagram.blocks) != 0 {
		t.Log("Failed to create a diagram with empty blocks")
		t.Fail()
	}

	s := diagram2String(diagram)

	if s != "{"+
		"\"width\": 0,"+
		"\"height\": 0,"+
		"\"blockWidth\": 0,"+
		"\"blockHeight\": 0,"+
		"\"blocks\": [],"+
		"\"connectors\": []"+
		"}\n" {
		t.Log("Failed to convert an empty diagram to string")
		t.Fail()
	}
}


func TestEmptySvg(t *testing.T) {
	diagram := Diagram{}
	diagram.width = 600
	diagram.height = 400
	diagram.blockWidth = 90
	diagram.blockHeight = 60

	if len(diagram.blocks) != 0 {
		t.Log("Failed to create a diagram with empty blocks")
		t.Fail()
	}

	s := Diagram2Svg(diagram)

	if s != "<svg width=\"600\" height=\"400\" align=\"center\">" +
	" <rect x=\"0\" y=\"0\" id=\"editor-canvas\" width=\"600\" height=\"400\" stroke=\"white\" fill=\"transparent\" stroke-width=\"0\"></rect>"+
	" <text x=\"20\" y=\"35\">Hello World!</text>"+
	"</svg>\n" {
		t.Log("Failed to convert an empty diagram to string")
		t.Fail()
	}
}
