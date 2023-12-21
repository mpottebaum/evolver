package models

type Ecosystem struct {
	Years       int
	Species     []Species
	Environment Environment
}

type Environment struct {
	Climate string
}
