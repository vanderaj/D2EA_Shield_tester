package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// ID,ShipName,HullMass,baseshipStrength

type shipT struct {
	ID                 int
	shipName           string
	hullMass           int
	baseShieldStrength int
}

func getShipStat(ships []shipT) shipT {
	var result shipT

	for _, ship := range ships {
		if ship.shipName == config.shipName {
			result = ship
			break
		}
	}

	return result
}

func loadShips() []shipT {

	var ships []shipT
	var record []string
	var err error

	csvfile, err := os.Open(config.shipFile)
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
		var ship shipT

		record, err = r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		//ID,ShipName,HullMass,baseShieldStrength
		ship.ID, err = strconv.Atoi(record[0])
		ship.shipName = record[1]
		ship.hullMass, err = strconv.Atoi(record[2])
		ship.baseShieldStrength, err = strconv.Atoi(record[3])

		ships = append(ships, ship)
	}

	return ships
}
