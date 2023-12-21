package models

type Species struct {
	Name        string
	Type        string
	GeneticCode []Phenotype
}

type Organism struct {
	Id        string
	Name      string
	Genotypes []Genotype
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
