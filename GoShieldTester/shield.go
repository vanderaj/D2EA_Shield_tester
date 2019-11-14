package main

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// ID,class,rating,type,maxmass,optmass,minmass,maxmul,optmul,minmul,regen,brokenregen

type shieldT struct {
	ID           int
	shieldClass  int
	shieldRating string
	shieldType   string
	maxmass      int
	optmass      int
	minmass      int
	maxmul       float64
	optmul       float64
	minmul       float64
	regen        float64
	brokenregen  float64
}

// return the number of HP in megajoules for a particular shield given a base ship shield
func getShieldHP(shieldGeneratorStat shieldT) float64 {

	// Terms re-used in this calculation
	term1 := shieldGeneratorStat.optmul - shieldGeneratorStat.minmul
	term2 := shieldGeneratorStat.maxmul - shieldGeneratorStat.minmul
	term3 := shieldGeneratorStat.maxmass - shieldGeneratorStat.minmass
	term4 := shieldGeneratorStat.maxmass - shieldGeneratorStat.optmass

	//  calculate the normalized mass
	//   [Double]$MassNorm = [math]::min(1, [Double]$(($maxmass - $ShipMass) / ($maxmass - $minmass)))
	massnorm := math.Min(1, float64((shieldGeneratorStat.maxmass-config.ship.hullMass)/term3))

	// Calculate power function exponent
	// [Double]$Exponent = [math]::Log10(($optmul - $minmul) / ($maxmul - $minmul)) /
	// [math]::Log10([math]::min(1,[Double]$(($maxmass - $optmass) / ($maxmass - $minmass))))

	massexp := math.Log10(term1/term2) / math.Log10(math.Min(1, float64(term4/term3)))

	// Calcualte final multiplier
	// [Double]$Multiplier = $minmul + [Double][math]::pow($MassNorm, $Exponent) * ($maxmul - $minmul)
	multiplier := shieldGeneratorStat.minmul + math.Pow(massnorm, massexp)*term2

	shieldHitPoints := float64(config.ship.baseShieldStrength) * multiplier

	return shieldHitPoints
}

func loadShields() []shieldT {

	var shields []shieldT
	var record []string
	var err error

	csvfile, err := os.Open(config.shieldFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvfile)

	if err != nil {
		log.Fatal(err)
	}

	// Consume and discard the header row
	record, err = r.Read()

	for {
		var shield shieldT

		record, err = r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// ID,class,rating,type,maxmass,optmass,minmass,maxmul,optmul,minmul,regen,brokenregen
		shield.ID, err = strconv.Atoi(record[0])
		shield.shieldClass, err = strconv.Atoi(record[1])
		shield.shieldRating = record[2]
		shield.shieldType = record[3]

		shield.maxmass, err = strconv.Atoi(record[4])
		shield.optmass, err = strconv.Atoi(record[5])
		shield.minmass, err = strconv.Atoi(record[6])

		shield.maxmul, err = strconv.ParseFloat(record[7], 64)
		shield.optmul, err = strconv.ParseFloat(record[8], 64)
		shield.minmul, err = strconv.ParseFloat(record[9], 64)

		shield.regen, err = strconv.ParseFloat(record[10], 64)
		shield.brokenregen, err = strconv.ParseFloat(record[11], 64)

		shields = append(shields, shield)
	}

	return shields
}
