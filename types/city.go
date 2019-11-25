package types

import (
	"math/rand"
)

// City : struct containing the city's name, aliens, and paths
type City struct {
	Name      string
	Aliens    []*Alien
	Paths     []*Path
	Destroyed bool
}

// NewCity :
func NewCity(name string) *City {
	return &City{
		Name:      name,
		Aliens:    make([]*Alien, 0),
		Paths:     make([]*Path, 0),
		Destroyed: false,
	}
}

// AddAlien :
func (c City) AddAlien(alien *Alien) {
	c.Aliens = append(c.Aliens, alien)
}

// AddPath :
func (c City) AddPath(path *Path) {
	c.Paths = append(c.Paths, path)
}

// GetRandomPath :
func (c City) GetRandomPath() *Path {
	return c.Paths[rand.Intn(c.NumPaths())]
}

// NumAliens :
func (c City) NumAliens() int {
	return len(c.Aliens)
}

// NumPaths :
func (c City) NumPaths() int {
	return len(c.Paths)
}
