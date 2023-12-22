package models

type Species struct {
	Name        string
	Type        string
	GeneticCode []Phenotype
}

type Population struct {
	Species   Species
	Organisms []Organism
	Zone      Zone
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
	Dominant  string
	Recessive string
}

type Genotype struct {
	Father string
	Mother string
}
