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

	if len(diagram.elements) != 0 {
		t.Log("Failed to create a diagram with empty blocks")
		t.Fail()
	}

	s := Diagram2String(diagram)

	if s != "{"+
		"\"width\": 0,"+
		"\"height\": 0,"+
		"\"blockWidth\": 0,"+
		"\"blockHeight\": 0,"+
		"\"elements\": [],"+
		"\"connectors\": []"+
		"}\n" {
		t.Log("Failed to convert an empty diagram to string")
		t.Fail()
	}
}


func TestAddElement(t *testing.T) {
	d := DefaultDiagram()

	if len(d.elements) != 0 {
		t.Log("Failed to create a diagram with empty blocks")
		t.Fail()
	}

	d = AddBlock(d, 20, 20)

	if len(d.elements) != 1 {
		t.Log("Failed to create a diagram with empty blocks")
		t.Fail()
	}
}



