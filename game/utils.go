package game

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/denalimarsh/invasion/types"
)

// ProcessFile :
func ProcessFile(filePath string, world *types.World) error {
	// Open file and create line scanner
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Scan each line,
	for scanner.Scan() {
		// Split the line on space
		text := strings.Split(scanner.Text(), " ")
		processLine(text, world)
	}
	// While loop breaks on error or EOF, check for error
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// processLine : processes a line into a City
func processLine(words []string, world *types.World) {
	// Attempt to load city
	cityName := words[0]
	city, err := world.GetCity(cityName)

	// If no result, create a new city
	if err != nil {
		// TODO: Either add this check below as well, or improve it
		if len(words) == 1 {
			log.Printf("'%v' has no paths and is inaccessible. Ignoring...", cityName)
		} else {
			city = types.NewCity(cityName)
			world.AddCity(city)
			world.AddLandingSite(cityName)
		}
	}

	// TODO: input validation on words
	for i := 1; i < len(words); i++ {
		pathText := strings.Split(strings.TrimSpace(words[i]), "=")

		// Get cardinal direction
		direction, err := types.StringToDirection(pathText[0])
		if err != nil {
			log.Fatal(err)
		}

		// Attempt to load destination city
		destCityName := strings.TrimSpace(pathText[1])
		destCity, err := world.GetCity(destCityName)

		// Create a new city for the destination city
		if err != nil {
			destCity = types.NewCity(destCityName)
			world.AddCity(destCity)
			world.AddLandingSite(destCityName)
		}

		// Create a new path and add it to the city
		path := types.NewPath(city, direction, destCity)
		city.AddPath(path)
	}
}
