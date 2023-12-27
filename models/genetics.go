package models

type Species struct {
	Name        string
	Type        string
	GeneticCode []Phenotype
	Populations []Population
}

type Population struct {
	Organisms []Organism
	Zone      Zone
	Quadrant  string
}

type Organism struct {
	Id          string
	Name        string
	Age         int
	Coordinates Coordinates
	Genotypes   []Genotype
}

type Phenotype struct {
	Name  string
	Genes []Gene
}

type Gene struct {
	Dominant, Recessive string
}

type Genotype struct {
	Father, Mother string
}
