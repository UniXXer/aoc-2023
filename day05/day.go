package day05

import (
	"strconv"
	"strings"
	"unixxer-aoc-2023/commons"
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
		default:
			{
				mappings := parseSeeds(line)
				if len(mappings) > 0 {
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
	}

	result := -1
	for _, v := range seeds {
		location := mapValue(mapValue(mapValue(mapValue(mapValue(mapValue(mapValue(v, seedToSoil), soilToFertilizer), fertilizerToWater), waterToLight), lightToTemperature), temperatureToHumidity), humidityToLocation)

		if result == -1 || result > location {
			result = location
		}
	}

	commons.PrintInfoFormat("The lowest location of seeds is %d", result)

	// PART 2
	result2 := -1
	for i := 0; i < len(seeds); i = i + 2 {
		start := seeds[i]
		length := seeds[i+1]

		for j := start; j < start+length; j++ {
			location := mapValue(mapValue(mapValue(mapValue(mapValue(mapValue(mapValue(j, seedToSoil), soilToFertilizer), fertilizerToWater), waterToLight), lightToTemperature), temperatureToHumidity), humidityToLocation)

			if result2 == -1 || result2 > location {
				result2 = location
			}
		}

	}

	commons.PrintInfoFormat("The lowest location of seed-range is %d", result2)

	return nil
}

func solveSecond(inputs []string) error {

	// not needed

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

func mapValue(value int, mapper [][]int) int {
	for i := 0; i < len(mapper); i++ {
		dest := mapper[i][0]
		source := mapper[i][1]
		length := mapper[i][2]

		if value >= source && value < source+length {
			return (value - source) + dest
		}
	}

	return value
}
