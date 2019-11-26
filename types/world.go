package types

import (
	"errors"
	"log"
	"math/rand"
)

// World : contains all cities, aliens, and cities indexed by number
type World struct {
	Cities   map[string]*City
	Aliens   map[int]*Alien
	SiteIDs  map[int]string
	RandSeed *rand.Rand
}

// NewWorld : initializes world and returns reference
func NewWorld(seed *rand.Rand) *World {
	return &World{
		Cities:   make(map[string]*City),
		Aliens:   make(map[int]*Alien),
		SiteIDs:  make(map[int]string),
		RandSeed: seed,
	}
}

// ProcessNewCity : creates a new city and adds it to the world
func (w *World) ProcessNewCity(city *City) {
	// Add the city to Cities
	w.Cities[city.GetName()] = city
	// Add the city to SiteIDs
	w.SiteIDs[w.numSites()] = city.GetName()
}

// ProcessNewAlien : creates a new alien and adds it to the world
func (w *World) ProcessNewAlien(alien *Alien) {
	w.Aliens[alien.GetID()] = alien
}

// PopulateAliens : randomly populates the world's cities with aliens
func (w *World) PopulateAliens(numAliens int) error {
	if w.NumCities() <= 0 {
		log.Fatal("world has no cities to populate")
	}

	for i := 1; i <= numAliens; i++ {
		// Generate random landing site ID for this alien
		siteID := rand.Intn(w.NumCities())

		targetCity, err := w.GetCityByID(siteID)
		if err != nil {
			log.Fatal(err)
		}

		// Make a new Alien and place it in the target city
		alien := NewAlien(i, targetCity)
		w.ProcessNewAlien(alien)
		// Update the city so it's aware of its new guest
		targetCity.AlienArrival(alien)
	}
	return nil
}

// DestroyCity : destroys a city, removing all Aliens and Paths
//				 associated with the city
func (w *World) DestroyCity(city *City) {
	// Remove all paths which reference the city
	city.RemoveAllPaths()

	// Remove any Aliens in the city
	for _, id := range city.GetAlienIDs() {
		delete(w.Aliens, id)
	}

	// Remove the city from the world
	delete(w.Cities, city.GetName())
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

// Print :
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
}

// --------------------------------------------------
//			Unexported helper methods
// --------------------------------------------------
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
	log.Printf("-------------------------------------------------")
	log.Printf("\t\t\tThe world:")
	log.Printf("-------------------------------------------------")
}
