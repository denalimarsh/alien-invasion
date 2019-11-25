package types

// Alien : an individual alien
type Alien struct {
	ID       int
	Location *City
	Alive    bool
}

// NewAlien :
func NewAlien(id int, location *City) *Alien {
	return &Alien{
		ID:       id,
		Location: location,
		Alive:    true,
	}
}
