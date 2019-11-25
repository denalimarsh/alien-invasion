package types

import (
	"errors"
)

// Direction : enum containing the four cardinal directions
type Direction string

const (
	// North :
	North = "north"
	// East :
	East = "east"
	// South :
	South = "south"
	// West :
	West = "west"
	// Default :
	Default = ""
)

// StringToDirection : returns string as direction
func StringToDirection(text string) (Direction, error) {
	switch text {
	case "north":
		return North, nil
	case "east":
		return East, nil
	case "south":
		return South, nil
	case "west":
		return West, nil
	default:
		return Default, nil // TODO: should be error?
	}
}

// String :
func (d Direction) String() (string, error) {
	switch d {
	case North:
		return "north", nil
	case East:
		return "east", nil
	case South:
		return "south", nil
	case West:
		return "west", nil
	default:
		return "", errors.New("invalid direction")
	}
}

// Integer :
func (d Direction) Integer() int {
	switch d {
	case North:
		return 0
	case East:
		return 1
	case South:
		return 2
	case West:
		return 3
	default:
		return -1
	}
}
