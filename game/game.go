package game

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/denalimarsh/invasion/types"
)

var world types.World

// Setup : generates a new World by processing the input file, then
//		   randomly populates the World's Cities with Alien invaders
func Setup(file string, numAliens int) error {
	// Generate unique source of randomness
	rand.Seed(time.Now().UnixNano())

	// Create a new world
	world = types.NewWorld()

	// Add cities, paths to world from given file
	err := ProcessFileToWorld(file, &world)
	if err != nil {
		log.Fatal(err)
	}

	// Print initial world state
	world.Print()

	// Randomly place aliens in cities
	err = world.PopulateAliens(numAliens)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Play : begins the invasion sequence which will execute for
//		  10,000 turns or until there is >= 1 alien remaining
func Play() error {
	logInvasionStart()
	turn := 1
	for turn <= 10000 && world.NumAliens() > 1 {
		moveAliens()
		cleanWorld(turn)
		turn++
	}
	logInvasionEnd(turn)
	return nil
}

// --------------------------------------------------
//			Unexported helper methods
// --------------------------------------------------
// moveAliens : randomaly moves all Aliens from their current location
//				to a new location (if available)
func moveAliens() {
	// Generate a new source of randomness each turn
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 1; i <= world.NumAliens(); i++ {
		alien, _ := world.GetAlienByID(i)
		if alien != nil {
			alien.Move(r)
		}
	}
}

// cleanWorld : removes any City which contains 2 or more Aliens, additionally
//				removing all local Aliens and all Paths which reference the
//				destroyed City
func cleanWorld(turn int) {
	for i := 0; i < world.NumCities(); i++ {
		city, _ := world.GetCityByID(i)
		if city != nil {
			if city.NumAliens() > 1 {
				world.DestroyCity(city)
				logCityDestroyed(turn, city.GetName(), city.GetAlienIDs())
			}
		}
	}
}

func logInvasionStart() {
	log.Println("Aliens, begin the invasion!")
	log.Println("")
}

func logInvasionEnd(turn int) {
	log.Println("")
	log.Printf("Invasion completed on turn %v.", turn)
	world.Print()
}

func logCityDestroyed(turn int, city string, ids []int) {
	alienIDs := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ", "), "[]")
	indexLastComma := strings.LastIndex(alienIDs, ",")
	aliens := alienIDs[:indexLastComma] + " and" + alienIDs[indexLastComma+1:]
	log.Printf("Turn %d: %s has been destroyed by aliens %s ", turn, city, aliens)
}
