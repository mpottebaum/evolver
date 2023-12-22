package models

// [x min, y min, x max, y max]
type Zone = [4]int

// [x, y]
type Coordinates = [2]int

type Overlapper = map[Coordinates][]Organism

type Quadrants = []Zone

type Grid struct {
	Width      int
	Height     int
	Overlapper Overlapper
	Quadrants  Quadrants
}
