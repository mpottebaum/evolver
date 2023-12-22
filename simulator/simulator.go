package simulator

import "evolver/models"

func Simulate(ecosystem models.Ecosystem, years int) models.Ecosystem {
	for ; ecosystem.Years < years; ecosystem.Years++ {
		ecosystem.Summer().Winter()
	}
	return ecosystem
}
