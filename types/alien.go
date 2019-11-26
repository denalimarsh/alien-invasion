package types

import (
	"log"
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
	if currCity.NumOutgoingPaths() > 0 {
		outoingPath, err := currCity.GetRandomOutgoingPath(r)
		if err != nil {
			// TODO: trapped function
			log.Printf("Alien %v is trapped!", a.GetID())
			return
		}
		destCity := outoingPath.GetCity()
		a.Location = destCity

		currCity.AlienDeparture(a)
		destCity.AlienArrival(a)
	}
}

// IsTrapped : returns true is the Alien is trapped in its current location
func (a *Alien) IsTrapped() bool {
	return a.GetLocation().NumOutgoingPaths() == 0
}

// GetID : returns the Alien's unique identifier
func (a *Alien) GetID() int {
	return a.ID
}

// GetLocation : returns the Alien's current location
func (a *Alien) GetLocation() *City {
	return a.Location
}
