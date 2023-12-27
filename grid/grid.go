package grid

import "evolver/models"

//(0,0)   (+x,0)
//	Q1 | Q2
//	---|---
//	Q4 | Q3
//(0,+y)  (+x,+y)

const Width int = 1000
const Height int = 1000

func CreateGrid(startQuadrant string) models.Grid {
	quadrants := models.Quadrants{
		// [x min, y min, x max, y max]
		"Q1": {0, 0, Width / 2, Height / 2},
		"Q2": {Width / 2, 0, Width, Height / 2},
		"Q3": {Width / 2, Height / 2, Width, Height},
		"Q4": {0, Height / 2, Width / 2, Height},
	}
	grid := models.Grid{
		Width:         Width,
		Height:        Height,
		Overlapper:    models.Overlapper{},
		Quadrants:     quadrants,
		StartQuadrant: startQuadrant,
	}
	return grid
}
