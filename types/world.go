package types

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

// World : contains all cities, aliens, and cities indexed by number
type World struct {
	Cities       map[string]*City
	Aliens       map[int]*Alien
	SiteIDs      map[int]string
	Seed         *rand.Rand
	AdvancedTech bool
}

// NewWorld : initializes world and returns reference
func NewWorld(seed *rand.Rand, advancedTech bool) *World {
	return &World{
		Cities:       make(map[string]*City),
		Aliens:       make(map[int]*Alien),
		SiteIDs:      make(map[int]string),
		Seed:         seed,
		AdvancedTech: advancedTech,
	}
}

// PopulateAliens : randomly populates the world's cities with aliens
func (w *World) PopulateAliens(numAliens int) error {
	if w.NumCities() <= 0 {
		log.Fatal("world has no cities to populate")
	}

	for i := 1; i <= numAliens; i++ {
		targetCity, err := w.GetCityByID(w.getSeed().Intn(w.NumCities()))
		if err != nil {
			log.Fatal(err)
		}

		alien := NewAlien(i, targetCity)
		w.ProcessNewAlien(alien)
		targetCity.AlienArrival(alien)
	}
	return nil
}

// ProcessNewCity : creates a new city and adds it to the world
func (w *World) ProcessNewCity(city *City) {
	w.Cities[city.GetName()] = city
	w.SiteIDs[w.numSites()] = city.GetName()
}

// ProcessNewAlien : creates a new alien and adds it to the world
func (w *World) ProcessNewAlien(alien *Alien) {
	w.Aliens[alien.GetID()] = alien
}

// MoveAliens : randomly moves all Aliens along available paths
func (w *World) MoveAliens(turn int) {
	for i := 1; i <= w.NumAliens(); i++ {
		alien, _ := w.GetAlienByID(i)
		if alien != nil {
			if !alien.IsTrapped() {
				alien.Move(w.getSeed())
			} else {
				// If advanced tech is enabled, aliens can teleport
				if w.hasAdvancedTech() {
					city := w.GetTeleportCity()
					printTeleport(turn, alien.GetID(), city.GetName())
					alien.Teleport(city)
				}
			}
		}
	}
}

// DestroyCities : removes cities with two or more Aliens
func (w *World) DestroyCities(turn int) {
	for i := 0; i < w.numSites(); i++ {
		city, _ := w.GetCityByID(i)
		if city != nil {
			if city.NumAliens() > 1 {
				w.DestroyCity(city)
				printCityDestroyed(turn, city.GetName(), city.GetAlienIDs())
			}
		}
	}
}

// DestroyCity : destroys a City, removing all Aliens and Paths
func (w *World) DestroyCity(city *City) {
	for _, id := range city.GetAlienIDs() {
		delete(w.Aliens, id)
	}
	city.RemoveAllPaths()
	delete(w.Cities, city.GetName())
}

// GetTeleportCity : returns a teleportable city
func (w *World) GetTeleportCity() *City {
	// Declare anon function so it's within scope
	var randTeleportableCity func() *City

	randTeleportableCity = func() *City {
		iterationIndex := w.getSeed().Intn(w.numSites())
		city, err := w.GetCityByID(iterationIndex)
		if err != nil {
			return randTeleportableCity()
		}
		if city.NumOutgoingPaths() > 0 {
			return city
		}
		return randTeleportableCity()
	}

	return randTeleportableCity()
}

// GetCityByName : returns a City by name
func (w *World) GetCityByName(city string) (*City, error) {
	if w.containsCity(city) {
		return w.Cities[city], nil
	}
	return nil, errors.New("city not found")
}

// GetCityByID : returns a City by ID
func (w *World) GetCityByID(city int) (*City, error) {
	return w.GetCityByName(w.SiteIDs[city])
}

// GetAlienByID : returns an Alien by ID
func (w *World) GetAlienByID(alien int) (*Alien, error) {
	if w.containsAlien(alien) {
		return w.Aliens[alien], nil
	}
	return nil, errors.New("alien not found")
}

// NumCities : returns the current number of active cities
func (w *World) NumCities() int {
	return len(w.Cities)
}

// NumAliens : returns the total number of aliens
func (w *World) NumAliens() int {
	return len(w.Aliens)
}

// Print : prints the world's cities and paths
func (w *World) Print() {
	printHeader()
	for i := 0; i < w.NumCities(); i++ {
		cityText := w.SiteIDs[i] + " "
		city, err := w.GetCityByID(i)
		if err != nil {
			continue
		}
		cityText += city.OutgoingPathsToString()
		log.Println(cityText)
	}
	printFooter()
}

// --------------------------------------------------
//			Unexported helper methods
// --------------------------------------------------
func (w *World) getSeed() *rand.Rand {
	return w.Seed
}

func (w *World) hasAdvancedTech() bool {
	return w.AdvancedTech
}

func (w *World) numSites() int {
	return len(w.SiteIDs)
}

func (w *World) containsCity(city string) bool {
	_, ok := w.Cities[city]
	return ok
}

func (w *World) containsAlien(alien int) bool {
	_, ok := w.Aliens[alien]
	return ok
}

func printHeader() {
	log.Println()
	log.Printf("\t\t\tThe World:")
	log.Printf("-------------------------------------------------")
}

func printFooter() {
	log.Printf("-------------------------------------------------")
	log.Println()
}

func printCityDestroyed(turn int, city string, ids []int) {
	alienIDs := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ", "), "[]")
	indexLastComma := strings.LastIndex(alienIDs, ",")
	aliens := alienIDs[:indexLastComma] + " and" + alienIDs[indexLastComma+1:]
	log.Printf("Turn %d: %s has been destroyed by aliens %s ", turn, city, aliens)
}

func printTeleport(turn int, id int, city string) {
	log.Printf("Turn %d: alien %d has teleported to %s ", turn, id, city)
}
