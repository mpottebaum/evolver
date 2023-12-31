package eco

import (
	models "evolver/models"
)

var PhenotypeLib = map[string][]models.Gene{
	// <name>: <genes>
	"size": {
		models.Gene{
			Dominant:  "small",
			Recessive: "large",
		},
	},
}

//	inputs:
//		species genetic codes and genotype frequencies,
//			[][]string [['species name','species type','phenotype name',],['species name','phenotype name',]]
//		environment
//			[]string ['climate']
//	outputs:
//		ecosystem with species and environment

func CreateEcosystem(environmentConfig []string, speciesConfigs [][]string) models.Ecosystem {
	// create environment
	environment := models.Environment{
		Climate: environmentConfig[0],
	}
	// create species structs
	species := []models.Species{}
	for _, speciesConfig := range speciesConfigs {
		// grab name and type from config slice
		newSpecies := models.Species{
			Name: speciesConfig[0],
			Type: speciesConfig[1],
		}
		// get phenotypes from remaining items in config slice
		speciesPhenotypes := []models.Phenotype{}
		for i := 2; i < len(speciesConfig); i++ {
			phenotypeName := speciesConfig[i]
			if phenotypeGenes, exists := PhenotypeLib[phenotypeName]; exists {
				phenotype := models.Phenotype{
					Name:  phenotypeName,
					Genes: phenotypeGenes,
				}
				speciesPhenotypes = append(speciesPhenotypes, phenotype)
			}
		}
		// add phenotypes to newSpecies
		newSpecies.GeneticCode = speciesPhenotypes
		// add species to ecosystem's species slice
		species = append(species, newSpecies)
	}

	return models.Ecosystem{
		Year:        1,
		Environment: environment,
		Species:     species,
	}
}
