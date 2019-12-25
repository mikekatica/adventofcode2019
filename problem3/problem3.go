package problem3

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type FuelSystemMatrix struct {
	dim [][]int
}

type Wire struct {
	xPos   int
	yPos   int
	startX int
	startY int
}

func NewWire(initXpos int, initYpos int) *Wire {
	w := Wire{
		xPos:   initXpos,
		yPos:   initYpos,
		startX: initXpos,
		startY: initYpos,
	}
	return &w
}

func (w *Wire) equals(w2 Wire) bool {
	if w.xPos == w2.xPos && w.yPos == w2.yPos {
		return true
	}
	return false
}

func (w *Wire) getDistance(w2 Wire) float64 {
	xDist := math.Abs(float64(w.startX - w2.xPos))
	yDist := math.Abs(float64(w.startY - w2.yPos))
	return xDist + yDist
}

func New(MaxRightSize int, MaxUpSize int) *FuelSystemMatrix {
	dimArr := make([][]int, MaxUpSize)
	for i := 0; i < MaxUpSize; i++ {
		yArr := make([]int, MaxRightSize)
		dimArr[i] = yArr
	}
	f := FuelSystemMatrix{
		dim: dimArr,
	}
	return &f
}

func (m *FuelSystemMatrix) WireRight(dist int, w *Wire) {
	newPos := w.xPos + dist
	fmt.Printf("Moving wire right %v to position %v, %v\n", dist, newPos, w.yPos)
	xArr := m.dim[w.yPos]
	for i := w.xPos + 1; i < newPos+1; i++ {
		xArr[i]++
	}
	w.xPos = newPos
}

func (m *FuelSystemMatrix) WireLeft(dist int, w *Wire) {
	newPos := w.xPos - dist
	fmt.Printf("Moving wire left %v to position %v, %v\n", dist, newPos, w.yPos)
	xArr := m.dim[w.yPos]
	for i := w.xPos - 1; i > newPos-1; i-- {
		xArr[i]++
	}
	w.xPos = newPos
}

func (m *FuelSystemMatrix) WireUp(dist int, w *Wire) {
	newPos := w.yPos + dist
	fmt.Printf("Moving wire up %v to position %v, %v\n", dist, w.xPos, newPos)
	for i := w.yPos + 1; i < newPos+1; i++ {
		xArr := m.dim[i]
		xArr[w.xPos]++
	}
	w.yPos = newPos
}

func (m *FuelSystemMatrix) WireDown(dist int, w *Wire) {
	newPos := w.yPos - dist
	fmt.Printf("Moving wire down %v to position %v, %v\n", dist, w.xPos, newPos)
	for i := w.yPos - 1; i > newPos-1; i-- {
		xArr := m.dim[i]
		xArr[w.xPos]++
	}
	w.yPos = newPos
}

func (m *FuelSystemMatrix) FindCrossedWires() *list.List {
	results := list.New()
	for i := range m.dim {
		for j := range m.dim[i] {
			if i != 0 && j != 0 && m.dim[i][j] > 1 {
				results.PushBack(Wire{
					xPos: j,
					yPos: i,
				})
			}
		}
	}
	return results
}

func ParseOp(op string) (string, int) {
	op1 := string(op[0])
	dist, _ := strconv.Atoi(op[1:])
	return op1, dist
}

func CalcMaxDimensions(wRep []string) (int, int, int, int) {
	maxX := 0
	minX := 0
	currentX := 0
	maxY := 0
	minY := 0
	currentY := 0
	for _, x := range wRep {
		op, dist := ParseOp(x)
		switch op {
		case "R":
			currentX += dist
			if currentX > maxX {
				maxX = currentX
			}
		case "L":
			currentX -= dist
			if currentX < minX {
				minX = currentX
			}
		case "U":
			currentY += dist
			if currentY > maxY {
				maxY = currentY
			}
		case "D":
			currentY -= dist
			if currentY < minY {
				minY = currentY
			}
		}
	}
	absMinX := int(math.Abs(float64(minX)))
	absMinY := int(math.Abs(float64(minY)))
	if absMinX > maxX {
		maxX = absMinX
	}
	if absMinY > maxY {
		maxY = absMinY
	}
	return maxX, maxY, minX, minY
}

func ParseOpAndApply(op string, m *FuelSystemMatrix, w *Wire) {
	op1, dist := ParseOp(op)
	switch op1 {
	case "R":
		m.WireRight(dist, w)
	case "L":
		m.WireLeft(dist, w)
	case "U":
		m.WireUp(dist, w)
	case "D":
		m.WireDown(dist, w)
	}
}

func ParseArrayFromProblemInput(path string) [][]string {
	dat, _ := ioutil.ReadFile(path)
	sdata := string(dat)
	trimData := strings.Split(strings.TrimSuffix(sdata, "\n"), "\n")
	var dataSplit [][]string
	for i := range trimData {
		dataSplit = append(dataSplit, strings.Split(trimData[i], ","))
	}
	return dataSplit
}

func GetLowestDistance(l *list.List, w1 *Wire) (float64, *Wire) {
	var w Wire
	lowestDistance := float64(0)
	current := l.Front()
	for i := 0; i < l.Len(); i++ {
		dist := w1.getDistance(current.Value.(Wire))
		if lowestDistance == 0 || (dist < lowestDistance && dist != 0) {
			lowestDistance = dist
			w = current.Value.(Wire)
		}
		current = current.Next()
	}
	return lowestDistance, &w
}

func DoProblem3() float64 {
	parsed := ParseArrayFromProblemInput("./problem3/input")
	p1 := parsed[0]
	p2 := parsed[1]

	var maxX, maxY, minX, minY int
	maxXp1, maxYp1, minXp1, minYp1 := CalcMaxDimensions(p1)
	maxXp2, maxYp2, minXp2, minYp2 := CalcMaxDimensions(p2)
	if maxXp1 > maxXp2 {
		maxX = maxXp1
	} else {
		maxX = maxXp2
	}
	if minXp1 < minXp2 {
		minX = minXp1
	} else {
		minX = minXp2
	}
	if maxYp1 > maxYp2 {
		maxY = maxYp1
	} else {
		maxY = maxYp2
	}
	if minYp1 < minYp2 {
		minY = minYp1
	} else {
		minY = minYp2
	}
	if minX > 0 {
		minX = 0
	}
	if minY > 0 {
		minY = 0
	}
	fmt.Printf("X Ranges: %v - %v, Y Ranges: %v - %v\n", minX, maxX, minY, maxY)
	fm := New((maxX-minX)+2, (maxY-minY)+2)
	startingPointX := 1 - minX
	startingPointY := 1 - minY
	w1 := NewWire(startingPointX, startingPointY)
	w2 := NewWire(startingPointX, startingPointY)
	fmt.Printf("Wire 1: %v\n", w1)
	fmt.Printf("Wire 2: %v\n", w2)
	for _, x := range p1 {
		ParseOpAndApply(x, fm, w1)
	}

	for _, x := range p2 {
		ParseOpAndApply(x, fm, w2)
	}
	crossedWires := fm.FindCrossedWires()
	dist, closest := GetLowestDistance(crossedWires, w1)
	fmt.Println(closest)
	return dist
}
