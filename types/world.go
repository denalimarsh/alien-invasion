package types

import (
	"errors"
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

// AddCity :
func (w *World) AddCity(city *City) {
	w.Cities[city.Name] = city
	w.LandingSites[w.NumCities()] = city.Name
}

// AddAlien :
func (w *World) AddAlien(alien *Alien) {
	w.Aliens[alien.ID] = alien
}

// ContainsCity :
func (w *World) ContainsCity(city string) bool {
	_, ok := w.Cities[city]
	if !ok {
		return false
	}
	return true
}

// ContainsAlien :
func (w *World) ContainsAlien(alien int) bool {
	_, ok := w.Aliens[alien]
	if !ok {
		return false
	}
	return true
}

// GetCity :
func (w *World) GetCity(city string) (*City, error) {
	if w.ContainsCity(city) {
		return w.Cities[city], nil
	}
	return nil, errors.New("city '%v' not found")
}

// GetAlien :
func (w *World) GetAlien(alien int) (*Alien, error) {
	if w.ContainsAlien(alien) {
		return w.Aliens[alien], nil
	}
	return nil, errors.New("alien '%v' not found")
}

// NumCities :
func (w *World) NumCities() int {
	return len(w.Cities)
}

// NumAliens :
func (w *World) NumAliens() int {
	return len(w.Aliens)
}
