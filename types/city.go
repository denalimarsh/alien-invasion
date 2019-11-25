package types

// City : struct containing the city's name, aliens, and paths
type City struct {
	Name      string
	Aliens    map[int]*Alien
	Paths     map[int]*Path
	Destroyed bool
}

// NewCity :
func NewCity(name string) *City {
	return &City{
		Name:      name,
		Aliens:    make(map[int]*Alien),
		Paths:     make(map[int]*Path), // TODO: change key type from int to Direction?
		Destroyed: false,
	}
}

// AddAlien :
func (c *City) AddAlien(alien *Alien) {
	c.Aliens[alien.ID] = alien
}

// AddPath :
func (c *City) AddPath(path *Path) {
	c.Paths[path.Direction.Integer()] = path
}

// GetPath :
func (c *City) GetPath(i int) *Path {
	return c.Paths[i]
}

// GetRandomPath :
// TODO: Need source of randomness for this f(x)
// func (c City) GetRandomPath() *Path {
// 	return c.Paths[rand.Intn(c.NumPaths())]
// }

// NumAliens :
func (c *City) NumAliens() int {
	return len(c.Aliens)
}

// PathsToString :
func (c *City) PathsToString() string {
	var buffer string
	for i := 0; i < 4; i++ {
		path := c.Paths[i]
		if path != nil {
			if path.Traversable {
				pathStr := path.String()
				buffer += pathStr + " "
			}
		}
	}
	return buffer
}
