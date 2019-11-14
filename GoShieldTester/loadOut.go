package main

type loadOutStatT struct {
	hitPoints           float64
	regenRate           float64
	explosiveResistance float64
	kineticResistance   float64
	thermalResistance   float64
}

func getLoadoutStats(shieldGeneratorVariant generatorT, shieldGeneratorBaseHitPoint float64, shieldGenertorBaseRecharge float64, shieldBoosterLoadout []int, boosterVariants []boosterT) loadOutStatT {

	var expModifier float64 = 1.0
	var kinModifier float64 = 1.0
	var thermModifier float64 = 1.0
	var boosterHitPointBonus float64 = 0.0

	var expRes, kinRes, thermRes, hitPoints float64

	// Compute non diminishing returns modifiers
	for _, booster := range shieldBoosterLoadout {

		var boosterVariantStats = boosterVariants[booster-1]

		expModifier = expModifier * boosterVariantStats.expResBonus
		kinModifier = kinModifier * boosterVariantStats.kinResBonus
		thermModifier = thermModifier * boosterVariantStats.thermResBonus
		boosterHitPointBonus = boosterHitPointBonus + boosterVariantStats.shieldStrengthBonus
	}

	// Compensate for diminishing returns
	if expModifier < 0.7 {
		expModifier = 0.7 - (0.7-expModifier)/2
	}

	if kinModifier < 0.7 {
		kinModifier = 0.7 - (0.7-kinModifier)/2
	}

	if thermModifier < 0.7 {
		thermModifier = 0.7 - (0.7-thermModifier)/2
	}

	// Compute final Resistance
	expRes = 1 - ((1 - shieldGeneratorVariant.expRes) * expModifier)
	kinRes = 1 - ((1 - shieldGeneratorVariant.kinRes) * kinModifier)
	thermRes = 1 - ((1 - shieldGeneratorVariant.thermRes) * thermModifier)

	// Compute final Hitpoints
	// $HitPoints = $ShieldGenertorBaseHitPoint * (1 + $ShieldGenratorVariant.OptimalMultiplierBonus) * (1 + $BoosterHitPointBonus)
	hitPoints = float64(shieldGeneratorBaseHitPoint) * (1 + shieldGeneratorVariant.optimalMultiplierBonus) * (1 + boosterHitPointBonus)

	return loadOutStatT{
		hitPoints:           hitPoints + config.scbHitPoint + config.guardianShieldHitPoint,
		regenRate:           shieldGenertorBaseRecharge * (1.0 + shieldGeneratorVariant.regenRateBonus),
		explosiveResistance: expRes,
		kineticResistance:   kinRes,
		thermalResistance:   thermRes,
	}
}
