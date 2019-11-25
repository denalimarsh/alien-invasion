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

	// TODO: parse cities, paths from file

	log.Printf("World: %v", world)

	// TODO: add cities, paths to world

	// TODO: generate aliens and place them in cities

	return nil
}

// TODO: implement Play
func Play() error {
	return nil
}
