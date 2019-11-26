package game

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/denalimarsh/invasion/types"
)

// Game : game contains world and a random seed for game RNG
type Game struct {
	World    *types.World
	RandSeed *rand.Rand
}

// NewGame : initializes a new game instance
func NewGame() *Game {
	// Generate seed for worldwide RNG
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	world := types.NewWorld()

	return &Game{
		World:    world,
		RandSeed: seed,
	}
}

// Setup : generates a new World by processing the input file, then
//		   randomly populates the World's Cities with Alien invaders
func (g *Game) Setup(file string, numAliens int) error {
	// Add cities, paths to world from given file
	err := g.LoadFileToWorld(file)
	if err != nil {
		log.Fatal(err)
	}

	g.getWorld().Print()

	// Randomly place aliens in cities
	err = g.getWorld().PopulateAliens(numAliens, g.getRandSeed())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Play : begins the invasion sequence which will execute for
//		    10,000 turns or until there is >= 1 alien remaining
func (g *Game) Play() error {
	g.logGameStart()
	turn := 1
	for turn <= 10000 && g.getWorld().NumAliens() > 1 {
		g.moveAliens()
		g.cleanWorld(turn)
		turn++
	}
	g.logGameEnd(turn)
	return nil
}

// --------------------------------------------------
//			Unexported helper methods
// --------------------------------------------------
// moveAliens : randomaly moves all Aliens from their current location
//				to a new location if it is available
func (g *Game) moveAliens() {
	for i := 1; i <= g.getWorld().NumAliens(); i++ {
		alien, _ := g.getWorld().GetAlienByID(i)
		if alien != nil {
			alien.Move(g.getRandSeed())
		}
	}
}

// cleanWorld : removes any City which contains 2 or more Aliens, additionally
//				removing all local Aliens and all Paths which reference the
//				destroyed City
func (g *Game) cleanWorld(turn int) {
	for i := 0; i < g.getWorld().NumCities(); i++ {
		city, _ := g.getWorld().GetCityByID(i)
		if city != nil {
			if city.NumAliens() > 1 {
				g.getWorld().DestroyCity(city)
				g.logCityDestroyed(turn, city.GetName(), city.GetAlienIDs())
			}
		}
	}
}

func (g *Game) logGameStart() {
	log.Println("The game has started!")
	log.Println("")
}

func (g *Game) logGameEnd(turn int) {
	log.Println("")
	log.Printf("Game completed on turn %v.", turn)
	g.getWorld().Print()
}

func (g *Game) logCityDestroyed(turn int, city string, ids []int) {
	alienIDs := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ", "), "[]")
	indexLastComma := strings.LastIndex(alienIDs, ",")
	aliens := alienIDs[:indexLastComma] + " and" + alienIDs[indexLastComma+1:]
	log.Printf("Turn %d: %s has been destroyed by aliens %s ", turn, city, aliens)
}

func (g *Game) getWorld() *types.World {
	return g.World
}

func (g *Game) getRandSeed() *rand.Rand {
	return g.RandSeed
}
