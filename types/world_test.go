package types

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	// Create a new world
	world := NewWorld(rand.New(rand.NewSource(time.Now().UnixNano())))

	// Confirm that the world has been instantiated
	assert.NotNil(t, world.Cities)
	assert.NotNil(t, world.Aliens)
	assert.NotNil(t, world.SiteIDs)
}

func TestProcessNewCity(t *testing.T) {
	world := NewWorld(rand.New(rand.NewSource(time.Now().UnixNano())))

	// Process the city
	city := NewCity(CityName1)
	world.ProcessNewCity(city)

	// Confirm that the city has been added can be accessed by both name and ID
	returnedCityByName, err := world.GetCityByName(CityName1)
	assert.Nil(t, err)
	assert.Equal(t, returnedCityByName, city)

	returnedCityByID, err := world.GetCityByID(0)
	assert.Nil(t, err)
	assert.Equal(t, returnedCityByID, city)

	// Confirm that the city count has increased
	numCities := world.NumCities()
	assert.Equal(t, numCities, 1)
}

func TestProcessNewAlien(t *testing.T) {
	world := NewWorld(rand.New(rand.NewSource(time.Now().UnixNano())))
	city := NewCity(CityName1)
	alien := NewAlien(AlienID1, city)

	// Process the alien
	world.ProcessNewAlien(alien)

	// Confirm that the alien has been added can be accessed by ID
	returnedAlienByID, err := world.GetAlienByID(AlienID1)
	assert.Nil(t, err)
	assert.Equal(t, returnedAlienByID, alien)

	// Confirm that the alien count has increased
	numAliens := world.NumAliens()
	assert.Equal(t, numAliens, 1)
}

func TestPopulateAliens(t *testing.T) {
	world := NewWorld(rand.New(rand.NewSource(time.Now().UnixNano())))

	// Setup some new cities
	city1 := NewCity(CityName1)
	city2 := NewCity(CityName2)
	world.ProcessNewCity(city1)
	world.ProcessNewCity(city2)

	// Populate the world with aliens
	world.PopulateAliens(TestNumAliens)

	// Confirm that the alien count has increased
	numAliens := world.NumAliens()
	assert.Equal(t, numAliens, TestNumAliens)
}

func TestDestroyCity(t *testing.T) {
	world := NewWorld(rand.New(rand.NewSource(time.Now().UnixNano())))

	// Setup a new city
	city1 := NewCity(CityName1)
	world.ProcessNewCity(city1)

	// Confirm that the city is available
	returnedCity1, err := world.GetCityByName(CityName1)
	assert.Nil(t, err)
	assert.Equal(t, returnedCity1, city1)

	// Confirm that the total count has increased
	numCities1 := world.NumCities()
	assert.Equal(t, numCities1, 1)

	// Destroy the city
	world.DestroyCity(city1)

	// Confirm that the city is NOT available
	_, err = world.GetCityByName(CityName1)
	assert.Error(t, err)

	// Confirm that the total count has decreased
	numCities2 := world.NumCities()
	assert.Equal(t, numCities2, 0)
}
