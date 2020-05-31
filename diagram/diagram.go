package diagram

import (
	"fmt"
)

// Block contains the x,y coordinates of the start of the block
type Block struct {
	x, y int
}

type Connector struct {
	
}

// Diagram contains a list of blocks (structure used by block-diagram-editor)
type Diagram struct {
	width  int
	height int
	blockWidth int
	blockHeight int
	blocks []Block
	connectors []Connector
}

func diagram2String(diagram Diagram) string {

	b := ""

	for _, s := range diagram.blocks {
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
		diagram.width, diagram.height, diagram.blockWidth, diagram.blockHeight, b)

	return s
}

func Diagram2Svg(diagram Diagram) string {

	b := ""

	for _, s := range diagram.blocks {
		if len(b) > 0 {
			b += ","
		}
		b += fmt.Sprintf("{\"x\": %d, \"y\": %d}", s.x, s.y)
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
		" <text x=\"20\" y=\"35\">Hello World!</text>"+
		"</svg>\n",
		width, height, width, height)

		return s
}
