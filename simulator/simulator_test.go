package simulator

import (
	"evolver/models"
	"testing"
)

var ecosystem = models.Ecosystem{
	Years: 0,
	Environment: models.Environment{
		Climate: "arid",
	},
	Species: []models.Species{
		{
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
		},
	},
}

func TestSimulatorYear(t *testing.T) {
	expectedYears := 10
	result := Simulate(ecosystem, expectedYears)
	if result.Years != expectedYears {
		t.Logf("Simulator year increments - exp: %v | rec: %v", expectedYears, result.Years)
		t.Fail()
	}
}
