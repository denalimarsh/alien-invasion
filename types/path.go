package types

// Path :
type Path struct {
	Direction   string
	Destination *City
	Traversable bool
}

// NewPath :
func NewPath(direction string, destination *City) *Path {
	return &Path{
		Destination: destination,
		Direction:   direction,
		Traversable: true,
	}
}
