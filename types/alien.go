package types

import (
	"math/rand"
)

// Alien : an individual Alien which contains its current location
//		   and a unique ID.
type Alien struct {
	ID       int
	Location *City
}

// NewAlien : returns a new Alien
func NewAlien(id int, location *City) *Alien {
	return &Alien{
		ID:       id,
		Location: location,
	}
}

// Move : randomly moves the Alien from their current city to a new city
func (a *Alien) Move(r *rand.Rand) {
	currCity := a.GetLocation()
	if !a.IsTrapped() {
		outgoingPath, _ := currCity.GetRandomOutgoingPath(r)
		destCity := outgoingPath.GetCity()
		currCity.AlienDeparture(a)
		destCity.AlienArrival(a)
		a.Location = destCity
	}
}

// IsTrapped : returns true is the Alien is trapped in its current location
func (a *Alien) IsTrapped() bool {
	return a.GetLocation().NumOutgoingPaths() == 0
}

// Teleport : teleports the alien to a new city, only allowed if the alien
//			  is trapped
func (a *Alien) Teleport(portCity *City) {
	currCity := a.GetLocation()
	currCity.AlienDeparture(a)
	portCity.AlienArrival(a)
	a.Location = portCity
}

// GetID : returns the Alien's unique identifier
func (a *Alien) GetID() int {
	return a.ID
}

// GetLocation : returns the Alien's current location
func (a *Alien) GetLocation() *City {
	return a.Location
}
