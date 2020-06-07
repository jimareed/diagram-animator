package diagram

import (
	"testing"
)

func TestEmptyBlockDiagram(t *testing.T) {
	diagram := BlockDiagram{}
	diagram.width = 0
	diagram.height = 0
	diagram.blockWidth = 0
	diagram.blockHeight = 0

	if len(diagram.blocks) != 0 {
		t.Log("Failed to create a diagram with empty blocks")
		t.Fail()
	}

	s := BlockDiagram2String(diagram)

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
