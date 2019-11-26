package types

import (
	"errors"
)

// Direction : enum containing the four cardinal directions
type Direction string

const (
	// North : cardinal direction #1
	North = "north"
	// East : cardinal direction #2
	East = "east"
	// South : cardinal direction #3
	South = "south"
	// West : cardinal direction #4
	West = "west"
	// Default : empty string indicates invalid value
	Default = ""
)

// StringToDirection : if valid, casts string to a Direction
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
		return Default, errors.New("invalid direction")
	}
}

// String : returns the Direction as a string
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

// Integer : returns the Direction's integer representation
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
