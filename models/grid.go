package models

import (
	"math/rand"
)

// [x min, y min, x max, y max]
type Zone [4]int

// [x, y]
type Coordinates [2]int

type Overlapper map[Coordinates][]Organism

type Quadrants map[string]Zone

type Grid struct {
	Width         int
	Height        int
	Overlapper    Overlapper
	Quadrants     Quadrants
	StartQuadrant string
}

func (g Grid) CreatePopulationZone(quadrantName string, zoneSize int) Zone {
	quadrant := g.Quadrants[quadrantName]
	quadWidth := quadrant[2] - quadrant[0]
	quadHeight := quadrant[3] - quadrant[1]
	var xMin, yMin, xMax, yMax int
	xMin = rand.Intn(quadWidth/2) + quadrant[0]
	xMax = xMin + zoneSize
	yMin = rand.Intn(quadHeight/2) + quadrant[1]
	yMax = yMin + zoneSize
	zone := Zone{
		xMin,
		yMin,
		xMax,
		yMax,
	}
	return zone
}
