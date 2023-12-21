package simulator

import "evolver/models"

// mating season
func summer(ecosystem models.Ecosystem) {

}

// survival of the fittest
func winter(ecosystem models.Ecosystem) {

}

func Simulate(ecosystem models.Ecosystem, years int) models.Ecosystem {
	for ; ecosystem.Years < years; ecosystem.Years++ {
		summer(ecosystem)
		winter(ecosystem)
	}
	return ecosystem
}
