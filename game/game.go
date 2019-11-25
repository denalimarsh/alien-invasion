package game

import (
	"log"

	"github.com/denalimarsh/invasion/types"
)

var world types.World

// Setup :
func Setup(file string, numAliens int) error {
	// Create an empty new world
	world := types.NewWorld()

	// Add cities, paths to world from given file
	err := ProcessFile(file, world)
	if err != nil {
		log.Fatal(err)
	}

	// Print starting world
	world.Print()

	// Randomly place aliens in cities
	err = world.PopulateAliens(numAliens)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Play :  TODO: implement Play
func Play() error {
	return nil
}
