package models

type Ecosystem struct {
	Years       int
	Environment Environment
	Species     []Species
	Populations []Population
}

// mating season
func (e Ecosystem) Summer() Ecosystem {
	return e
}

// survival of the fittest
func (e Ecosystem) Winter() Ecosystem {
	return e
}
