package types

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// World :
type World struct {
	Cities       map[string]*City
	Aliens       map[int]*Alien
	LandingSites map[int]string
}

// NewWorld : initializes world and returns reference
func NewWorld() *World {
	return &World{
		Cities:       make(map[string]*City),
		Aliens:       make(map[int]*Alien),
		LandingSites: make(map[int]string),
	}
}

// PopulateAliens :
func (w *World) PopulateAliens(numAliens int) error {
	if w.NumCities() < 1 {
		log.Fatal("this world has no cities to populate")
	}
	// TODO: Move random out of this function
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= numAliens; i++ {
		// Generate random landing site ID for this alien
		siteID := rand.Intn(len(w.Cities))

		targetCity, err := w.GetCity(w.LandingSites[siteID])
		if err != nil {
			log.Fatal(err)
		}

		// Make a new alien
		alien := NewAlien(i, targetCity)
		// Add alien to the world's invasion force
		w.AddAlien(alien)
		// Add the alien to the city's invasion force
		log.Printf("alien %v added to %v", alien.ID, targetCity.Name) // TODO: Remove print
		targetCity.AddAlien(alien)
	}
	return nil
}

// AddCity :
func (w *World) AddCity(city *City) {
	w.Cities[city.Name] = city
}

// AddLandingSite :
func (w *World) AddLandingSite(site string) {
	cityCount := len(w.LandingSites)
	w.LandingSites[cityCount] = site
}

// AddAlien :
func (w *World) AddAlien(alien *Alien) {
	w.Aliens[alien.ID] = alien
}

// ContainsCity :
func (w *World) ContainsCity(city string) bool {
	_, ok := w.Cities[city]
	return ok
}

// ContainsAlien :
func (w *World) ContainsAlien(alien int) bool {
	_, ok := w.Aliens[alien]
	return ok
}

// GetCity :
func (w *World) GetCity(city string) (*City, error) {
	if w.ContainsCity(city) {
		return w.Cities[city], nil
	}
	return nil, errors.New("city not found")
}

// GetAlien :
func (w *World) GetAlien(alien int) (*Alien, error) {
	if w.ContainsAlien(alien) {
		return w.Aliens[alien], nil
	}
	return nil, errors.New("alien not found")
}

// NumCities :
func (w *World) NumCities() int {
	return len(w.Cities)
}

// NumAliens :
func (w *World) NumAliens() int {
	return len(w.Aliens)
}

// Print :
func (w *World) Print() {
	for i := 0; i < w.NumCities(); i++ {
		cityNumber := strconv.Itoa(i)
		cityName := w.LandingSites[i]
		cityLine := cityNumber + ". " + cityName + ": "
		city, err := w.GetCity(cityName)
		if err != nil {
			log.Fatal(err)
		}
		// TODO: Print city's aliens
		if !city.Destroyed {
			cityLine += city.PathsToString()
			fmt.Println(cityLine)
		}
	}
}
