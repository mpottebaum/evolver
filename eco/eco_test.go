package eco

import (
	"evolver/grid"
	"evolver/models"
	"testing"
)

func TestCreateEcosystem(t *testing.T) {
	startQuad := "Q1"

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
		Populations: []models.Population{
			{
				Quadrant: startQuad,
			},
		},
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
	grid := grid.CreateGrid(startQuad)
	result := CreateEcosystem(envConfig, speciesConfigs, grid)
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
	if pops := species.Populations[0]; pops.Quadrant != startQuad || pops.Zone[0] < grid.Quadrants[pops.Quadrant][0] || pops.Zone[1] < grid.Quadrants[pops.Quadrant][1] || pops.Zone[2] > grid.Quadrants[pops.Quadrant][2] || pops.Zone[3] > grid.Quadrants[pops.Quadrant][3] {
		t.Logf("Populations - exp: %v | rec: %v", expectedSpecies.Populations[0], pops)
		t.Logf("Start Quadrant: %v", grid.Quadrants[pops.Quadrant])
		t.Fail()
	}
}
