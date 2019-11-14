package main

import (
	"flag"
	"fmt"
)

type configT struct {
	shipName                                          string
	shieldGeneratorSize                               int
	shieldBoosterCount                                int
	prismatics                                        bool
	explosiveDPS, kineticDPS, thermalDPS, absoluteDPS float64
	damageEffectiveness                               float64
	scbHitPoint, guardianShieldHitPoint               float64
	boosterFile, generatorFile, shipFile, shieldFile  string
	ship                                              shipT
}

const defaultShipName = "Alliance Chieftain"
const defaultShieldGeneratorSize = 6
const libPath string = "../lib/"

var config configT

func processFlags() {
	flag.StringVar(&config.shipName, "shipname", config.shipName, "Type of ship, will load defaults for that ship type")
	flag.IntVar(&config.shieldGeneratorSize, "gensize", config.shieldGeneratorSize, "Define the size of the shield generator")
	flag.IntVar(&config.shieldBoosterCount, "boosters", config.shieldBoosterCount, "Number of Shield Boosters")
	flag.Float64Var(&config.explosiveDPS, "edps", config.explosiveDPS, "Explosive DPS percentage (use 0 for Thargoids)")
	flag.Float64Var(&config.kineticDPS, "kdps", config.kineticDPS, "Kinetic DPS percentage (use 0 for Thargoids)")
	flag.Float64Var(&config.thermalDPS, "tdps", config.thermalDPS, "Thermal DPS percentage (use 0 for Thargoids)")
	flag.Float64Var(&config.absoluteDPS, "adps", config.absoluteDPS, "Absolute DPS percentage (use 100 for Thargoids)")
	flag.Float64Var(&config.damageEffectiveness, "dmg", config.damageEffectiveness, "Damage effectiveness (25 for fixed weapons, 65 for hazrez PvP, 100 constant attack)")
	flag.Float64Var(&config.scbHitPoint, "scb", config.scbHitPoint, "SCB HitPoints (default 0)")
	flag.Float64Var(&config.guardianShieldHitPoint, "gshp", config.guardianShieldHitPoint, "Guardian HitPoints (default 0)")

	prismatics := flag.Bool("noprismatics", false, "Disable Prismatic shields")
	pve := flag.Bool("pve", false, "PvE gimballed suit high hazrez defaults")
	pvp := flag.Bool("pvp", false, "PvP fixed PA/railgun meta defaults")
	thargoid := flag.Bool("thargoid", false, "Useful Thargoid defaults")
	shortboost := flag.Bool("shortboost", false, "Load the short booster list")

	flag.Parse()
	config.damageEffectiveness = config.damageEffectiveness / 100 // convert from integer to percentage

	if config.shieldBoosterCount < 0 {
		fmt.Println("Can't have negative shield boosters, setting to 0")
		config.shieldBoosterCount = 0
	}

	if config.shieldBoosterCount > 8 {
		fmt.Println("No current ship has more than 8 shield boosters, setting to 8")
		config.shieldBoosterCount = 8
	}

	if *prismatics {
		fmt.Println("Disabling Prismatics")
		config.prismatics = false
	}

	if *shortboost {
		fmt.Println("Loading short booster list")
		config.boosterFile = "../ShieldBoosterVariants_short.csv"
	}

	if *pve {
		if *pvp {
			fmt.Println("Loading PvP defaults")
			config.explosiveDPS = 10 // rare to see missles
			config.kineticDPS = 60   // PvE NPCs often have MCs, etc
			config.thermalDPS = 50   // PvE NPCs seem to love their lasers
			config.absoluteDPS = 0   // rare to come across NPCs with rails or PAs, YMMV
			config.damageEffectiveness = 0.65
		}
	}

	if *pvp {
		fmt.Println("Loading PvP defaults")
		config.explosiveDPS = 10 // missles etc
		config.kineticDPS = 83   // multi-cannons, cannons, frag cannons, etc
		config.thermalDPS = 47   // lasers
		config.absoluteDPS = 30  // Plasma accelerators are meta, as are rail guns with super penetrator
		config.damageEffectiveness = 0.50
	}

	if *thargoid {
		fmt.Println("Loading Thargoid defaults")
		config.explosiveDPS = 0
		config.kineticDPS = 0
		config.thermalDPS = 0
		config.absoluteDPS = 200          // this is a higher level Thargoid, use 65 for scouts, 100 for Threat 5, 200 for 8+
		config.damageEffectiveness = 0.30 // many Thargoid attacks are spaced a bit apart
	}
}

func loadConfig() configT {

	config = configT{
		shipName:               defaultShipName,
		shieldGeneratorSize:    defaultShieldGeneratorSize,
		shieldBoosterCount:     3,    // 1-4 = small ships (2 typical), 2-6 = medium ships (4 typical), 4-8 = large ships (6-7 typical)
		prismatics:             true, // do you have prismatics unlocked?
		explosiveDPS:           33,   // missles
		kineticDPS:             33,   // cannons and missles
		thermalDPS:             33,   // laser weapons
		absoluteDPS:            0,    // 0 for most NPCs except Thargoids, 100 for a Cyclops, 150 for a Basilisk, 200 for a Hydra
		damageEffectiveness:    65,   // 25 = fixed weapons or single enemy, 65 = hazrez, 100% always being attacked
		scbHitPoint:            0,    // Using Corolis, enter the total in the Defense > Shield Sources > Cells green pie slice. e.g. 2 x 7A SCBs is 3621 MJ *2 = 7242 MJ
		guardianShieldHitPoint: 0,    // Using Corolis, enter the total in the Defense > Shield Sources > Shield Additions yellow pie slice. e.g. 1 x 5D GSRB = 215 MJ
		boosterFile:            libPath + "ShieldBoosterVariants.csv",
		generatorFile:          libPath + "ShieldGeneratorVariants.csv",
		shipFile:               libPath + "ShipStats.csv",
		shieldFile:             libPath + "ShieldStats.csv",
	}

	return config
}
