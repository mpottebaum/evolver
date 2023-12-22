package models

type Ecosystem struct {
	Years       int
	Environment Environment
	Species     []Species
	Populations []Population
}

type Environment struct {
	Climate string
}
