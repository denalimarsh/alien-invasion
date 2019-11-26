package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	// Initialize the world
	InitWorld()

	// Setup the world using the generic file
	file := LoadFilePath()

	// Capture Setup's output for testing
	output := CaptureOutput(func() {
		err := Setup(file, TestNumAliens)
		assert.Nil(t, err)
	})

	// Confirm that the output contain's the expected value
	assert.Contains(t, output, "The world:")
}

func TestInvade(t *testing.T) {
	// Initialize the world
	InitWorld()

	// Setup the world using the generic file
	file := LoadFilePath()
	err := Setup(file, TestNumAliens)
	assert.Nil(t, err)

	// Capture Setup's output for testing
	output := CaptureOutput(func() {
		err = Invade()
		assert.Nil(t, err)
	})

	// Confirm that the output contain's the expected values
	assert.Contains(t, output, "Aliens, begin the invasion!")
	assert.Contains(t, output, "Invasion completed on turn")
}
