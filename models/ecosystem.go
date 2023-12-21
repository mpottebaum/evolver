package models

type Ecosystem struct {
	Year        int
	Species     []Species
	Environment Environment
}

type Environment struct {
	Climate string
}
