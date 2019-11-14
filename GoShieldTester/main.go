package main

/*
 * Go version of Down To Earth Astronomy's Shield Tester
 * Original version here:  https://github.com/DownToEarthAstronomy/D2EA_Shield_tester
 *
 * Go port by Andrew van der Stock, vanderaj@gmail.com
 *
 * Why this version? It's fast - about 15,500 times faster than the PowerShell version, so multi-threading is not necessary even with 8 utility slots and the full list
 */

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Down to Earth Astronomy's ShieldTester (https://github.com/DownToEarthAstronomy/D2EA_Shield_tester)")
	fmt.Println("Go port v0.05 by Andrew van der Stock, vanderaj@gmail.com")
	fmt.Println()

	config = loadConfig()

	processFlags()

	var ships = loadShips()
	config.ship = getShipStat(ships)

	if config.ship.ID == 0 {
		fmt.Println("The ship [", config.shipName, "] is not in the list of supported ships. Using the default.")
		config.shipName = defaultShipName
		config.ship = getShipStat(ships)
	}

	if config.shieldGeneratorSize < 1 || config.shieldGeneratorSize > 8 {
		fmt.Println("Generator sizes can only be between 1 and 8. Resetting to the default")
		config.shieldGeneratorSize = defaultShieldGeneratorSize
	}

	var shields = loadShields()

	var generators = loadGenerators()

	var boosterVariants = loadboosterVariants()

	fmt.Printf("Loaded %d shields and %d boosters\n", len(generators), len(boosterVariants))

	var shieldBoosterLoadoutList = getBoosterLoadoutList(len(boosterVariants))

	startTime := time.Now()
	var result = testGenerators(generators, shields, boosterVariants, shieldBoosterLoadoutList)
	endTime := time.Now()
	dur := endTime.Sub(startTime)

	fmt.Println("Tested", len(shieldBoosterLoadoutList)*len(generators), "loadouts in", dur)

	showResults(result, boosterVariants, dur)
}
