package game

import (
	"log"
	"math/rand"
	"time"

	"github.com/denalimarsh/invasion/types"
	"github.com/denalimarsh/invasion/utils"
)

var world *types.World

// Init : initializes the world with a new RNG seed
func Init() {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	world = types.NewWorld(seed)
}

// Setup : generates a new World by processing the input file, then
//		   randomly populates the World's Cities with Alien invaders
func Setup(file string, numAliens int) error {
	err := utils.LoadFileToWorld(world, file)
	if err != nil {
		log.Fatal(err)
	}

	err = world.PopulateAliens(numAliens)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Play : begins the invasion sequence which will execute for
//		    10,000 turns or until there is >= 1 alien remaining
func Play() error {
	logGameStart()
	turn := 1
	for turn <= 10000 && world.NumAliens() > 1 {
		world.MoveAliens(turn)
		world.DestroyCities(turn)
		turn++
	}
	logGameEnd(turn)
	return nil
}

// --------------------------------------------------
//			Unexported helper methods
// --------------------------------------------------
func logGameStart() {
	log.Println("The game has started!")
	log.Println("")
}

func logGameEnd(turn int) {
	log.Println("")
	log.Printf("Game completed on turn %v.", turn)
	world.Print()
}
