package types

import (
	"errors"
	"math/rand"
)

// City : a node which is connected to other Cities via Paths
//		  and can host a mapping of Aliens
type City struct {
	Name          string
	OutgoingPaths []*Path
	IncomingPaths []*Path
	Aliens        map[int]*Alien
}

// NewCity : returns a new City
func NewCity(name string) *City {
	return &City{
		Name:          name,
		OutgoingPaths: make([]*Path, 0),
		IncomingPaths: make([]*Path, 0),
		Aliens:        make(map[int]*Alien),
	}
}

// RegisterOutgoingPath : registers an outgoing path e.g. (This_City -> Other_City)
func (c *City) RegisterOutgoingPath(path *Path) {
	c.OutgoingPaths = append(c.OutgoingPaths, path)
}

// RegisterIncomingPath : registers an incoming path e.g. (Other_City -> This_City)
func (c *City) RegisterIncomingPath(path *Path) {
	c.IncomingPaths = append(c.IncomingPaths, path)
}

// GetRandomOutgoingPath : gets a random outgoing path
func (c City) GetRandomOutgoingPath(r *rand.Rand) (*Path, error) {
	if c.NumOutgoingPaths() == 0 {
		return nil, errors.New("this city has no remaining outgoing paths")
	}
	return c.OutgoingPaths[r.Intn(c.NumOutgoingPaths())], nil
}

// RemoveAllPaths : removes all paths worldwide which reference this city
func (c *City) RemoveAllPaths() {
	// Remove all outgoing paths
	c.OutgoingPaths = c.OutgoingPaths[:0]
	for i := 0; i < len(c.IncomingPaths); i++ {
		incomingPath := c.IncomingPaths[i]
		if incomingPath != nil {
			cityWithIncomingPath := incomingPath.GetCity()
			cityWithIncomingPath.removeOutgoingPath(c.GetName())
		}
	}
}

// AlienArrival : records the arrival of an Alien visitor
func (c *City) AlienArrival(alien *Alien) {
	c.Aliens[alien.GetID()] = alien
}

// AlienDeparture : records the departure of an Alien guest
func (c *City) AlienDeparture(alien *Alien) {
	delete(c.Aliens, alien.GetID())
}

// GetAlienIDs : get the unique IDs
func (c *City) GetAlienIDs() []int {
	ids := make([]int, 0, len(c.Aliens))
	for id := range c.Aliens {
		ids = append(ids, id)
	}
	return ids
}

// GetName : returns the City's name
func (c *City) GetName() string {
	return c.Name
}

// NumAliens : returns the current number of Aliens hosted by the city
func (c *City) NumAliens() int {
	return len(c.Aliens)
}

// NumOutgoingPaths : returns the current number of outgoing Paths
func (c *City) NumOutgoingPaths() int {
	return len(c.OutgoingPaths)
}

// OutgoingPathsToString : formats the City's outgoing Paths in string
//						   representation for printing
func (c *City) OutgoingPathsToString() string {
	var buffer string
	for i := 0; i < c.NumOutgoingPaths(); i++ {
		path := c.OutgoingPaths[i]
		if path != nil {
			pathStr := path.String()
			buffer += pathStr + " "
		}
	}
	return buffer
}

// --------------------------------------------------
//			Unexported helper methods
// --------------------------------------------------
// removeOutgoingPath : instructs another city to remove any paths which reference this citiy
func (c *City) removeOutgoingPath(toCity string) {
	for i := len(c.OutgoingPaths) - 1; i >= 0; i-- {
		if c.OutgoingPaths[i].GetCity().GetName() == toCity {
			c.OutgoingPaths = append(c.OutgoingPaths[:i], c.OutgoingPaths[i+1:]...)
		}
	}
}
