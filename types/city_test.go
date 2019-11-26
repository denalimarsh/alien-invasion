package types

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCity(t *testing.T) {
	// Make a new city
	city := NewCity(CityName1)

	// Confirm that the city has been instantiated
	assert.Equal(t, city.GetName(), CityName1)
	assert.NotNil(t, city.IncomingPaths)
	assert.NotNil(t, city.OutgoingPaths)
	assert.NotNil(t, city.Aliens)
}

func TestRegisterOutgoingPath(t *testing.T) {
	// Setup two new cities and a direction
	city1 := NewCity(CityName1)
	city2 := NewCity(CityName2)
	north, err := StringToDirection("north")
	assert.Nil(t, err)

	// Create a new path and register it on the first city
	outPath := NewPath(city2, north)
	city1.RegisterOutgoingPath(outPath)

	// Confirm that the new outgoing path matches
	returnedOutPath := city1.OutgoingPaths[0]
	assert.Equal(t, outPath, returnedOutPath)
}

func TestRegisterIncomingPath(t *testing.T) {
	// Setup two new cities and a direction
	city1 := NewCity(CityName1)
	city2 := NewCity(CityName2)
	north, err := StringToDirection("north")
	assert.Nil(t, err)

	// Create a new path and register it on the second city
	incPath := NewPath(city1, north)
	city2.RegisterIncomingPath(incPath)

	// Confirm that the new incoming path matches
	returnedIncPath := city2.IncomingPaths[0]
	assert.Equal(t, incPath, returnedIncPath)
}

func TestGetRandomOutgoingPath(t *testing.T) {
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

	// Random path should return the same path (only 1 available)
	randPath, err := city1.GetRandomOutgoingPath(seed)
	assert.Nil(t, err)
	assert.Equal(t, randPath, outPath)
}

func TestRemoveAllPaths(t *testing.T) {
	// Setup two new cities and a direction
	city1 := NewCity(CityName1)
	city2 := NewCity(CityName2)
	north, err := StringToDirection("north")
	assert.Nil(t, err)

	// Register an incoming and outgoing path on each city, respectively
	outPath := NewPath(city2, north)
	city1.RegisterOutgoingPath(outPath)
	incPath := NewPath(city1, north)
	city2.RegisterIncomingPath(incPath)

	// Attempt to remove all paths
	city1.RemoveAllPaths()

	// Confirm that the new outgoing path has been deleted
	lenOutPaths := len(city1.OutgoingPaths)
	assert.Equal(t, 0, lenOutPaths)

	// Confirm that the new incoming path has been deleted
	lenInPaths := len(city1.IncomingPaths)
	assert.Equal(t, 0, lenInPaths)
}

func TestAlienArrival(t *testing.T) {
	// Setup new city
	city1 := NewCity(CityName1)

	// Create new Alien and register it in the city
	alien1 := NewAlien(AlienID1, city1)
	city1.AlienArrival(alien1)

	// Confirm that the Alien is registered
	ids := city1.GetAlienIDs()
	assert.Equal(t, ids[0], AlienID1)
}

func TestAlienDeparture(t *testing.T) {
	// Setup new city
	city1 := NewCity(CityName1)

	// Create new Alien, register, then unregister
	alien1 := NewAlien(AlienID1, city1)
	city1.AlienArrival(alien1)
	city1.AlienDeparture(alien1)

	// Confirm that the Alien is no longer registered
	ids := city1.GetAlienIDs()
	assert.Equal(t, len(ids), 0)
}
