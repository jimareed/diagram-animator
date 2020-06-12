package diagram

import (
	"fmt"
	"math"
)

const arrowHeadLength = 21

// Block contains the x,y coordinates of the start of the block
type Element struct {
	x, y float64
	elementType string
	description string
	size int
	url string
}

type Point struct {
	x, y float64
}

type Connector struct {
	b1, b2 int
}

type Transition struct {
	duration int
}

// Diagram contains a list of blocks (structure used by block-diagram-editor)
type Diagram struct {
	width  float64
	height float64
	blockWidth float64
	blockHeight float64
	elements []Element
	connectors []Connector
	transitions []Transition
}

func DefaultDiagram() Diagram {
	d := Diagram{}

	d.width = 900.0
	d.height = 600.0
	d.blockWidth = 90.0
	d.blockHeight = 60.0

	return d
}

func AddBlock(d Diagram, x int, y int) Diagram {

	d.elements = append(d.elements, Element{float64(x),float64(y),"block","",0,""})
	return d
}

func AddConnector(d Diagram, b1 int, b2 int) Diagram {

	d.connectors = append(d.connectors, Connector{b1,b2})
	return d
}

func AddText(d Diagram, x int, y int, description string, size int) Diagram {

	return AddTextWithUrl(d, x, y, description, size, "")
}

func AddTextWithUrl(d Diagram, x int, y int, description string, size int, url string) Diagram {

	d.elements = append(d.elements, Element{float64(x),float64(y),"text",description,size,url})
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
		b += fmt.Sprintf("{\"x\": %f, \"y\": %f, \"elementType\": \"%s\"}", e.x, e.y, e.elementType)
	}

	s := fmt.Sprintf(
		"{"+
			"\"width\": %f,"+
			"\"height\": %f,"+
			"\"blockWidth\": %f,"+
			"\"blockHeight\": %f,"+
			"\"elements\": [%s],"+
			"\"connectors\": []"+
			"}\n",
		diagram.width, diagram.height, diagram.blockWidth, diagram.blockHeight, b)

	return s
}

func slope(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
    return (y2 - y1) / (x2 - x1)
}


func arrowHeadX(slope float64) float64 {
    return arrowHeadLength / math.Sqrt(slope * slope + 1)
}

func calcP1(d Diagram, c Connector) Point {
	p := Point{0.0, 0.0}
	p1 := Point{d.elements[c.b1].x, d.elements[c.b1].y}
	p2 := Point{d.elements[c.b2].x, d.elements[c.b2].y}

	s := float64(slope(p1.x, p1.y, p2.x, p2.y))

	/*
    if s == +Inf || s == -Inf {
		p.x = p1.x + d.blockWidth / 2;
		if (this.props.p1.y < this.props.p2.y) {
		  y = this.props.p1.y + this.props.blockHeight;
		} else {
		  y = this.props.p1.y;
		}
	} else {

	}
	*/

	if math.Abs(s) <= slope(0.0,0.0,d.blockWidth, d.blockHeight) {
		// right side
		if p1.x < p2.x {
			p.x = p1.x + d.blockWidth;
			p.y = p1.y + d.blockHeight / 2 + d.blockWidth / 2 * s
		} else {
		// left side
			p.x = p1.x;
			p.y = p1.y + d.blockHeight / 2 - d.blockWidth / 2 * s
		}
	} else {
		// top side
		if (p1.y > p2.y) {
			p.x = p1.x + d.blockWidth / 2 - (d.blockHeight / 2) / s
			p.y = p1.y
		// botton side
		} else {
			p.x = p1.x + d.blockWidth / 2 + (d.blockHeight / 2) / s
			p.y = p1.y + d.blockHeight
		}
	}

	return p
}

func calcP2(d Diagram, c Connector) Point {
	p := Point{213,165}

	p1 := Point{d.elements[c.b1].x, d.elements[c.b1].y}
	p2 := Point{d.elements[c.b2].x, d.elements[c.b2].y}

	s := float64(slope(p1.x, p1.y, p2.x, p2.y))

	arrowHeadX := arrowHeadX(s);
	arrowHeadY := arrowHeadX * s;

	if math.Abs(s) <= float64(slope(0,0,d.blockWidth, d.blockHeight)) {
		// right side
		if p1.x < p2.x {
			p.x = p2.x - arrowHeadX
			p.y = p2.y + d.blockHeight / 2 - d.blockWidth / 2 * s - arrowHeadY
		} else {
		// left side
			p.x = p2.x + d.blockWidth + arrowHeadX
			p.y = p2.y + d.blockHeight / 2 + d.blockWidth / 2 * s + arrowHeadY
		}
	} else {
		// top side
		if (p1.y > p2.y) {
			p.x = p1.x + d.blockWidth / 2 - (d.blockHeight / 2) / s
			p.y = p1.y
		// botton side
		} else {
			p.x = p1.x + d.blockWidth / 2 + (d.blockHeight / 2) / s
			p.y = p1.y + d.blockHeight
		}
	}
	
	return p
/*
    var x = 0;
    var y = 0;

    if (slope === Infinity|| slope === -Infinity) {
      x = this.props.p2.x + this.props.blockWidth / 2;
      if (this.props.p1.y < this.props.p2.y) {
        y = this.props.p2.y - arrowHeadLength;
      } else {
        y = this.props.p2.y + this.props.blockHeight + arrowHeadLength;
      }
    } else {
      var arrowHeadX = this.arrowHeadX(slope);
      var arrowHeadY = arrowHeadX * slope;
  
      if (Math.abs(slope) <= this.slope(0,0,this.props.blockWidth,this.props.blockHeight)) {
        // right side
        if (this.props.p1.x < this.props.p2.x) {
          x = this.props.p2.x;
          y = this.props.p2.y + this.props.blockHeight / 2 - this.props.blockWidth / 2 * slope;
  
          if (drawArrowHead) {
            x -= arrowHeadX;
            y -= arrowHeadY;
          }
          console.log("right: " + arrowHeadX , "," + arrowHeadY + "slope:" + slope)
        }
        // left side
        else {
          x = this.props.p2.x + this.props.blockWidth;
          y = this.props.p2.y + this.props.blockHeight / 2 + this.props.blockWidth / 2 * slope;
          if (drawArrowHead) {
            x += arrowHeadX;
            y += arrowHeadY;
          }
          console.log("left: " + arrowHeadX , "," + arrowHeadY + "slope:" + slope)
        }
      } else {
        // top side
        if (this.props.p1.y > this.props.p2.y) {
          x = this.props.p2.x + this.props.blockWidth / 2 + (this.props.blockHeight / 2) / slope;
          y = this.props.p2.y + this.props.blockHeight;
          if (drawArrowHead) {
            if (this.props.p1.x < this.props.p2.x) {
              arrowHeadX = arrowHeadX * -1;
            }
            x += arrowHeadX
            y += Math.abs(arrowHeadY);
          }
          console.log("top: " + arrowHeadX , "," + arrowHeadY + "slope:" + slope)
         }
        // botton side
        else {
          x = this.props.p2.x + this.props.blockWidth / 2 - (this.props.blockHeight / 2) / slope;
          y = this.props.p2.y;
          if (drawArrowHead) {
            if (this.props.p1.x < this.props.p2.x) {
              arrowHeadX = arrowHeadX * -1;
            }
            x += arrowHeadX;
            y -= Math.abs(arrowHeadY);
          }
          console.log("bottom: " + arrowHeadX , "," + arrowHeadY + "slope:" + slope)
         }
      }
    }

    return ({
      x: x,
      y: y
    })
*/

}

func Diagram2Svg(diagram Diagram) string {

	elements := ""
	i := 0
	for _, e := range diagram.elements {
		if e.elementType == "block" {
			elements += fmt.Sprintf("<rect class=\"transition%d\" x=\"%f\" y=\"%f\" width=\"90\" height=\"60\" id=\"1\" stroke=\"black\" fill=\"transparent\" stroke-width=\"4\"></rect>\n", i, e.x, e.y)
		} 
		if e.elementType == "text" {
			elements += fmt.Sprintf("<text class=\"transition%d\" x=\"%f\" y=\"%f\" fill=\"black\" font-size=\"%dpx\">%s</text>\n", i, e.x, e.y, e.size, e.description)
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

	connectors := ""
	i = 1
	for _, c := range diagram.connectors {
		p1 := calcP1(diagram, c)
		p2 := calcP2(diagram, c)
		connectors += fmt.Sprintf(
			"<line class=\"transition%d\" x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" stroke=\"black\" stroke-width=\"4\" marker-end=\"url(#arrowhead)\"></line>",
			i, p1.x, p1.y, p2.x, p2.y)
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
		"<svg width=\"%f\" height=\"%f\" align=\"center\">" +
        " <rect x=\"0\" y=\"0\" id=\"editor-canvas\" width=\"%f\" height=\"%f\" stroke=\"white\" fill=\"transparent\" stroke-width=\"0\"></rect>"+
		"%s\n"+
		"<defs>\n"+
        "<marker id=\"arrowhead\" markerWidth=\"5\" markerHeight=\"3.5\" refX=\"0\" refY=\"1.75\" orient=\"auto\">\n"+
        "    <polygon points=\"0 0, 5 1.75 0 3.5\"></polygon>\n"+
        "</marker>\n"+
    	"</defs>\n"+
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
		width, height, width, height, elements, connectors, transitions)

		return s
}
