package game

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/denalimarsh/invasion/types"
)

// ProcessFileToWorld : processes a text file containing a list of
//						cities and paths into a World.
func ProcessFileToWorld(filePath string, world *types.World) error {
	// Open file and create line scanner
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Scan each line
	for scanner.Scan() {
		// Split the line on space
		text := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(text) > 0 {
			processLine(text, world)
		}
	}
	// While loop breaks on error or EOF, check for error
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// processLine : processes a line of text to create a city along with its paths
func processLine(words []string, world *types.World) {
	// Attempt to load the source city
	srcCityName := words[0]
	srcCity, err := world.GetCityByName(srcCityName)
	// If source city doesn't exist, create and process it
	if err != nil {
		srcCity = types.NewCity(srcCityName)
		world.ProcessNewCity(srcCity)
	}

	for i := 1; i < len(words); i++ {
		pathText := strings.Split(strings.TrimSpace(words[i]), "=")

		// Get cardinal direction, must be valid or world generation will fail
		direction, err := types.StringToDirection(strings.ToLower(pathText[0]))
		if err != nil {
			log.Fatal(err)
		}

		// Attempt to load the named destination city
		destCityName := strings.TrimSpace(pathText[1])
		destCity, err := world.GetCityByName(destCityName)
		// If destination city doesn't exist, create and process it
		if err != nil {
			destCity = types.NewCity(destCityName)
			world.ProcessNewCity(destCity)
		}

		// Create a new outgoing path and add it to the src city
		outgoingPath := types.NewPath(destCity, direction)
		srcCity.RegisterOutgoingPath(outgoingPath)

		// Create a new incoming path and add it to the dest city
		incomingPath := types.NewPath(srcCity, direction)
		destCity.RegisterIncomingPath(incomingPath)
	}
}
