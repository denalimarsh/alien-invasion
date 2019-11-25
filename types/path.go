package types

import "log"

// Path :
type Path struct {
	Origin      *City
	Direction   Direction
	Destination *City
	Traversable bool
}

// NewPath :
func NewPath(origin *City, direction Direction, destination *City) *Path {
	return &Path{
		Origin:      origin,
		Destination: destination,
		Direction:   direction,
		Traversable: true,
	}
}

// String :
func (p *Path) String() string {
	direction, err := p.Direction.String()
	if err != nil {
		log.Fatal(err)
	}
	destination := p.Destination.Name
	return direction + "=" + destination
}
