package eco

import (
	"evolver/models"
	"testing"
)

func TestCreateEcosystem(t *testing.T) {
	t.Logf("CreateEcosystem")

	expectedYear := 0
	expectedSpecies := models.Species{
		Name: "dopeness",
		Type: "animal",
		GeneticCode: []models.Phenotype{
			{
				Name: "size",
				Genes: []models.Gene{
					{
						Dominant:  "small",
						Recessive: "large",
					},
				},
			},
		},
	}
	expectedEnv := models.Environment{
		Climate: "arid",
	}

	envConfig := []string{
		expectedEnv.Climate,
	}
	speciesConfigs := [][]string{
		{
			expectedSpecies.Name,
			expectedSpecies.Type,
			expectedSpecies.GeneticCode[0].Name,
		},
	}
	result := CreateEcosystem(envConfig, speciesConfigs)
	if result.Years != expectedYear {
		t.Logf("Year - exp: %v | rec: %v", expectedYear, result.Years)
		t.Fail()
	}
	species := result.Species[0]
	if species.Name != expectedSpecies.Name || species.Type != expectedSpecies.Type || species.GeneticCode[0].Name != expectedSpecies.GeneticCode[0].Name {
		t.Logf("Species - exp: %v | rec: %v", expectedSpecies, species)
		t.Fail()
	}
	if environment := result.Environment; environment.Climate != expectedEnv.Climate {
		t.Logf("Environment - exp: %v | rec: %v", expectedEnv, environment)
		t.Fail()
	}
}
