package types

import "log"

// Path : a direct connection between two cities
type Path struct {
	City      *City
	Direction Direction
}

// NewPath : returns a new Path
func NewPath(city *City, direction Direction) *Path {
	return &Path{
		City:      city,
		Direction: direction,
	}
}

// GetDirection : returns the Path's Direction
func (p *Path) GetDirection() Direction {
	return p.Direction
}

// GetCity : returns the Path's City
func (p *Path) GetCity() *City {
	return p.City
}

// String : returns the Path's string representation
func (p *Path) String() string {
	direction, err := p.GetDirection().String()
	if err != nil {
		log.Fatal(err)
	}
	destination := p.GetCity().GetName()
	return direction + "=" + destination
}
