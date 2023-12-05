package day05

import (
	"fmt"
	"strconv"
	"strings"
)

func solveFirst(inputs []string) error {

	seedToSoil := [][]int{}
	soilToFertilizer := [][]int{}
	fertilizerToWater := [][]int{}
	waterToLight := [][]int{}
	lightToTemperature := [][]int{}
	temperatureToHumidity := [][]int{}
	humidityToLocation := [][]int{}

	seeds := parseSeeds(inputs[0])

	state := ""

	for i := 2; i < len(inputs); i++ {
		line := strings.TrimSpace(inputs[i])
		switch line {
			case "seed-to-soil map:":
				state = line
			case "soil-to-fertilizer map:":
				state = line
			case "fertilizer-to-water map:":
				state = line
			case "water-to-light map:":
				state = line
			case "light-to-temperature map:":
				state = line
			case "temperature-to-humidity map:":
				state = line
			case "humidity-to-location map:":
				state = line
			default: {
				mappings := parseSeeds(line)

				switch state {
					case "seed-to-soil map:":
						seedToSoil = append(seedToSoil, mappings)
					case "soil-to-fertilizer map:":
						soilToFertilizer = append(soilToFertilizer, mappings)
					case "fertilizer-to-water map:":
						fertilizerToWater = append(fertilizerToWater, mappings)
					case "water-to-light map:":
						waterToLight = append(waterToLight, mappings)
					case "light-to-temperature map:":
						lightToTemperature = append(lightToTemperature, mappings)
					case "temperature-to-humidity map:":
						temperatureToHumidity = append(temperatureToHumidity, mappings)
					case "humidity-to-location map:":
						humidityToLocation = append(humidityToLocation, mappings)					
				}
			}	
		}
	}

	fmt.Println(seeds)
	fmt.Println(seedToSoil)
	fmt.Println(waterToLight)
	fmt.Println(humidityToLocation)

	fmt.Printf("solveFirst is not implemented yet!")
	fmt.Println()

	return nil

}

func solveSecond(inputs []string) error {

	fmt.Printf("solveSecond is not implemented yet!")
	fmt.Println()

	return nil

}

func parseSeeds(line string) []int {
	splits := strings.Split(line, " ")
	seeds := []int{}

	for _, seed := range splits {
		s, err := strconv.Atoi(seed)
		if err == nil {
			seeds = append(seeds, s)
		}
	}

	return seeds
}

func mapValue(value int, mapper[][]int) int {
	
}