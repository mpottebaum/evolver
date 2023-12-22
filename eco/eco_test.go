package eco

import (
	"evolver/models"
	"testing"
)

func TestCreateEcosystem(t *testing.T) {
	t.Logf("CreateEcosystem")

	expectedYear := 0
	expectedEnv := models.Environment{
		Climate: "arid",
	}
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
	expectedPop := models.Population{
		Species: expectedSpecies,
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
	if environment := result.Environment; environment.Climate != expectedEnv.Climate {
		t.Logf("Environment - exp: %v | rec: %v", expectedEnv, environment)
		t.Fail()
	}
	species := result.Species[0]
	if species.Name != expectedSpecies.Name || species.Type != expectedSpecies.Type || species.GeneticCode[0].Name != expectedSpecies.GeneticCode[0].Name {
		t.Logf("Species - exp: %v | rec: %v", expectedSpecies, species)
		t.Fail()
	}
	if pops := result.Populations[0]; pops.Species.Name != expectedSpecies.Name {
		t.Logf("Populations - exp: %v | rec: %v", expectedPop, pops)
		t.Fail()
	}
}
