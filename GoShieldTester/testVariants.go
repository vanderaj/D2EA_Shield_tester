package main

import (
	"sync"
)

func testCase(ch chan resultT, wg *sync.WaitGroup, shieldGenerator generatorT, shields []shieldT, boosterVariants []boosterT, shieldBoosterLoadoutList [][]int) {
	bestTestCase := resultT{
		survivalTime: 0.0,
	}

	var result resultT
	var shieldRating string

	if shieldGenerator.genType == "Bi-Weave" {
		shieldRating = "C"
	} else {
		shieldRating = "A"
	}

	generatorStat := getGeneratorStat(shieldRating, shieldGenerator, shields)
	shieldGeneratorBaseHitPoint := getShieldHP(generatorStat)

	for _, shieldBoosterLoadout := range shieldBoosterLoadoutList {
		var loadoutStats = getLoadoutStats(shieldGenerator, shieldGeneratorBaseHitPoint, generatorStat.regen, shieldBoosterLoadout, boosterVariants)

		var actualDPS float64 = config.damageEffectiveness*
			(config.explosiveDPS*(1-loadoutStats.explosiveResistance)+
				config.kineticDPS*(1-loadoutStats.kineticResistance)+
				config.thermalDPS*(1-loadoutStats.thermalResistance)+
				config.absoluteDPS) - loadoutStats.regenRate*(1-config.damageEffectiveness)

		var survivalTime float64 = loadoutStats.hitPoints / actualDPS

		result = resultT{
			shieldGenerator:      shieldGenerator,
			shieldBoosterLoadout: shieldBoosterLoadout,
			loadOutStats:         loadoutStats,
			survivalTime:         survivalTime,
		}

		if actualDPS > 0 && bestTestCase.survivalTime >= 0 {
			if result.survivalTime > bestTestCase.survivalTime {
				bestTestCase = result
			}
		} else if actualDPS < 0 {
			if result.survivalTime < bestTestCase.survivalTime {
				bestTestCase = result
			}
		}
	}

	ch <- bestTestCase
	wg.Done()
}

func testGenerators(generators []generatorT, shields []shieldT, boosterVariants []boosterT, boosterList [][]int) resultT {
	bestResult := resultT{survivalTime: 0.0}

	ch := make(chan resultT, len(generators))
	wg := sync.WaitGroup{}

	for _, generator := range generators {
		wg.Add(1)
		go testCase(ch, &wg, generator, shields, boosterVariants, boosterList)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		if bestResult.survivalTime < 0 {
			if result.survivalTime < bestResult.survivalTime {
				bestResult = result
			}
		} else {
			if result.survivalTime < 0 {
				bestResult = result
			} else if result.survivalTime > bestResult.survivalTime {
				bestResult = result
			}
		}
	}

	return bestResult
}
