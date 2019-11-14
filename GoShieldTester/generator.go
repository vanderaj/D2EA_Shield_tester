package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// ID,Type,Engineering,Experimental,RegenRateBobus,ExpRes,KinRes,ThermRes,OptimalMultiplierBonus

type generatorT struct {
	ID                       int
	genType                  string
	engineering              string
	experimental             string
	shieldStrength           float64
	regenRateBonus           float64
	expRes, kinRes, thermRes float64
	optimalMultiplierBonus   float64
}

func getGeneratorStat(rating string, generator generatorT, shields []shieldT) shieldT {
	var result shieldT

	for _, shield := range shields {
		if shield.shieldType == generator.genType && shield.shieldClass == config.shieldGeneratorSize && shield.shieldRating == rating {
			result = shield
		}
	}

	return result
}

func loadGenerators() []generatorT {

	var generators []generatorT
	var record []string
	var err error

	csvfile, err := os.Open(config.generatorFile)
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
		var generator generatorT

		record, err = r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// if prismatics are disabled, skip those entries
		if !config.prismatics && strings.Contains(record[1], "Prismatic") {
			continue
		}

		// ID,Type,Engineering,Experimental,RegenRateBobus,ExpRes,KinRes,ThermRes,OptimalMultiplierBonus

		generator.ID, err = strconv.Atoi(record[0])
		generator.genType = record[1]
		generator.engineering = record[2]
		generator.experimental = record[3]

		generator.regenRateBonus, err = strconv.ParseFloat(record[4], 64)
		generator.expRes, err = strconv.ParseFloat(record[5], 64)
		generator.expRes = 1.0 - generator.expRes
		generator.kinRes, err = strconv.ParseFloat(record[6], 64)
		generator.kinRes = 1.0 - generator.kinRes
		generator.thermRes, err = strconv.ParseFloat(record[7], 64)
		generator.thermRes = 1.0 - generator.thermRes
		generator.optimalMultiplierBonus, err = strconv.ParseFloat(record[8], 64)

		generators = append(generators, generator)
	}

	return generators
}
