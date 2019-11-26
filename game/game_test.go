package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	// Load our generic file
	file := LoadFilePath()

	// Initalize a new World
	Init(false)

	err := Setup(file, TestNumAliens)
	assert.Nil(t, err)

}

func TestPlay(t *testing.T) {
	// Load our generic file
	file := LoadFilePath()

	// Initalize a new World
	Init(false)

	err := Setup(file, TestNumAliens)
	assert.Nil(t, err)

	// Capture Invade's output for testing
	output := CaptureOutput(func() {
		err = Play()
		assert.Nil(t, err)
	})

	// Confirm that the output contain's the expected values
	assert.Contains(t, output, "The game has started!")
	assert.Contains(t, output, "Game completed on turn")
}
