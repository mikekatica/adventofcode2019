package problem3

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	parsed := ParseArrayFromProblemInput("./input_test")
	correct := [][]string{{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, {"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}}
	for i, x := range correct {
		for j, y := range x {
			if y != parsed[i][j] {
				t.Errorf("Problem. %v did not parse into %v.", parsed, correct)
			}
		}
	}
}

func TestWireCompare(t *testing.T) {
	w1 := Wire{
		xPos:   1,
		yPos:   7,
		startX: 0,
		startY: 0,
	}
	w2 := Wire{
		xPos:   2,
		yPos:   3,
		startX: 0,
		startY: 0,
	}
	if w1.equals(w2) {
		t.Error("The wires are not actually equal")
	}
	if !w1.equals(w1) {
		t.Error("A wire is always equal to itself, and it is not reporting as such.")
	}
}

func TestWireRight(t *testing.T) {
	originalWire := NewWire(0, 0)
	newWire := Wire{
		xPos:   5,
		yPos:   0,
		startX: 0,
		startY: 0,
	}
	matrix := New(6, 2)
	desMatrix := [][]int{{0, 1, 1, 1, 1, 1}, {0, 0, 0, 0, 0, 0}}
	matrix.WireRight(5, originalWire)
	if !originalWire.equals(newWire) {
		t.Errorf("Wires do not match: %v : %v", originalWire, newWire)
	}
	for i, x := range desMatrix {
		for j, y := range x {
			if y != matrix.dim[i][j] {
				t.Errorf("The wire matrices do not match: %v, %v", matrix.dim, desMatrix)
			}
		}
	}
}

func TestWireUp(t *testing.T) {
	originalWire := NewWire(0, 0)
	newWire := Wire{
		xPos:   0,
		yPos:   5,
		startX: 0,
		startY: 0,
	}
	matrix := New(2, 6)
	desMatrix := [][]int{{0, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}}
	matrix.WireUp(5, originalWire)
	if !originalWire.equals(newWire) {
		t.Errorf("Wires do not match: %v : %v", originalWire, newWire)
	}
	for i, x := range desMatrix {
		for j, y := range x {
			if y != matrix.dim[i][j] {
				t.Errorf("The wire matrices do not match: %v, %v", matrix.dim, desMatrix)
			}
		}
	}
}

func TestWireLeft(t *testing.T) {
	originalWire := NewWire(5, 0)
	newWire := Wire{
		xPos:   0,
		yPos:   0,
		startX: 0,
		startY: 0,
	}
	matrix := New(6, 2)
	desMatrix := [][]int{{1, 1, 1, 1, 1, 0}, {0, 0, 0, 0, 0, 0}}
	matrix.WireLeft(5, originalWire)
	if !originalWire.equals(newWire) {
		t.Errorf("Wires do not match: %v : %v", originalWire, newWire)
	}
	for i, x := range desMatrix {
		for j, y := range x {
			if y != matrix.dim[i][j] {
				t.Errorf("The wire matrices do not match: %v, %v", matrix.dim, desMatrix)
			}
		}
	}
}

func TestWireDown(t *testing.T) {
	originalWire := NewWire(0, 5)
	newWire := Wire{
		xPos:   0,
		yPos:   0,
		startX: 0,
		startY: 0,
	}
	matrix := New(2, 6)
	desMatrix := [][]int{{1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {0, 0}}
	matrix.WireDown(5, originalWire)
	if !originalWire.equals(newWire) {
		t.Errorf("Wires do not match: %v : %v", originalWire, newWire)
	}
	for i, x := range desMatrix {
		for j, y := range x {
			if y != matrix.dim[i][j] {
				t.Errorf("The wire matrices do not match: %v, %v", matrix.dim, desMatrix)
			}
		}
	}
}

func TestCrossingWires(t *testing.T) {
	m1 := [][]int{{2, 1, 1}, {1, 0, 1}, {1, 1, 2}}
	matrix := FuelSystemMatrix{
		dim: m1,
	}
	results := matrix.FindCrossedWires().Front().Value.(Wire)
	rWire := Wire{2, 2, 0, 0}
	if !rWire.equals(results) {
		t.Errorf("Results of finding crossed wires did not return %v, got %v",
			rWire, results)
	}
}

func TestMaxDim(t *testing.T) {
	ops := []string{"R7", "U10", "L8", "D12", "R2", "U5"}
	maxX := 7
	maxY := 10
	minX := -1
	minY := -2
	tMaxX, tMaxY, tMinX, tMinY := CalcMaxDimensions(ops)
	if maxX != tMaxX || maxY != tMaxY || minX != tMinX || minY != tMinY {
		t.Errorf("Wanted max: %v,%v min: %v,%v : Got max: %v,%v min: %v,%v",
			maxX, maxY, minX, minY, tMaxX, tMaxY, tMinX, tMinY)
	}
}

func TestSimpleMatrix(t *testing.T) {
	o1 := []string{"R8", "U5", "L5", "D3"}
	o2 := []string{"U7", "R6", "D4", "L4"}
	m := New(10, 10)
	w1 := NewWire(0, 0)
	w2 := NewWire(0, 0)
	for _, x := range o1 {
		ParseOpAndApply(x, m, w1)
	}
	for _, x := range o2 {
		ParseOpAndApply(x, m, w2)
	}
	for i := len(m.dim) - 1; i >= 0; i-- {
		fmt.Println(m.dim[i])
	}
}

func TestAll(t *testing.T) {
	fmt.Println("================")
	parsed := ParseArrayFromProblemInput("./input_test")
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
	fmt.Println("================")
	for _, x := range p2 {
		ParseOpAndApply(x, fm, w2)
	}
	crossedWires := fm.FindCrossedWires()
	dist, closest := GetLowestDistance(crossedWires, w1)
	fmt.Println(dist)
	fmt.Println(closest)
	if dist != 135 {
		t.Errorf("Distances do not match. Calculated %v, need %v", dist, 135)
	}
}
func TestAll2(t *testing.T) {
	fmt.Println("================")
	parsed := ParseArrayFromProblemInput("./input_test2")
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
	fmt.Println("================")
	for _, x := range p2 {
		ParseOpAndApply(x, fm, w2)
	}
	crossedWires := fm.FindCrossedWires()
	dist, closest := GetLowestDistance(crossedWires, w1)
	fmt.Println(dist)
	fmt.Println(closest)
	if dist != 159 {
		t.Errorf("Distances do not match. Calculated %v, need %v", dist, 159)
	}
}
