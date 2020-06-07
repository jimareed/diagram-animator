package diagram

import (
	"fmt"
)

// Block contains the x,y coordinates of the start of the block
type Block struct {
	x, y int
}

// Diagram contains a list of blocks (structure used by block-diagram-editor)
type BlockDiagram struct {
	width  int
	height int
	blockWidth int
	blockHeight int
	blocks []Block
	connectors []Connector
}

func BlockDiagram2String(blockDiagram BlockDiagram) string {

	b := ""

	for _, s := range blockDiagram.blocks {
		if len(b) > 0 {
			b += ","
		}
		b += fmt.Sprintf("{\"x\": %d, \"y\": %d}", s.x, s.y)
	}

	s := fmt.Sprintf(
		"{"+
			"\"width\": %d,"+
			"\"height\": %d,"+
			"\"blockWidth\": %d,"+
			"\"blockHeight\": %d,"+
			"\"blocks\": [%s],"+
			"\"connectors\": []"+
			"}\n",
		blockDiagram.width, blockDiagram.height, blockDiagram.blockWidth, blockDiagram.blockHeight, b)

	return s
}

