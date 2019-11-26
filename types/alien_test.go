package types

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAlien(t *testing.T) {
	// Make new city for testing purposes
	city1 := NewCity(CityName1)

	// Create new alien
	alien1 := NewAlien(AlienID1, city1)

	// Confirm that the alien has been instantiated
	assert.Equal(t, alien1.GetID(), AlienID1)
	assert.Equal(t, alien1.GetLocation(), city1)
}

func TestMove(t *testing.T) {
	// Generate new source of randomness
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Setup two new cities and a direction
	city1 := NewCity(CityName1)
	city2 := NewCity(CityName2)
	north, err := StringToDirection("north")
	assert.Nil(t, err)

	// Create a new path and register it on the first city
	outPath := NewPath(city2, north)
	city1.RegisterOutgoingPath(outPath)

	// Create new alien and place in first city
	alien1 := NewAlien(AlienID1, city1)

	// With one path, random move will move alien to second city
	alien1.Move(seed)

	// Confirm that the alien has moved to the second city
	newCity := alien1.GetLocation()
	assert.Equal(t, newCity, city2)
}

func TestIsTrapped(t *testing.T) {
	// Generate new source of randomness
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Setup two new cities and a direction
	city1 := NewCity(CityName1)
	city2 := NewCity(CityName2)
	north, err := StringToDirection("north")
	assert.Nil(t, err)

	// Create a new path and register it on the first city
	outPath := NewPath(city2, north)
	city1.RegisterOutgoingPath(outPath)

	// Create new alien and place in first city
	alien1 := NewAlien(AlienID1, city1)

	// Alien isn't trapped, city1 has a path
	assert.False(t, alien1.IsTrapped())

	// With one path, random move will move alien to second city
	alien1.Move(seed)

	// Confirm that the alien has moved to the second city
	newCity := alien1.GetLocation()
	assert.Equal(t, newCity, city2)

	// Alien is now trapped, city2 does not have a path
	assert.True(t, alien1.IsTrapped())
}
