package game

import (
	"testing"

	"github.com/invasion/game"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	// Load our generic file
	file := LoadFilePath()

	Init()

	// Capture Setup's output for testing
	output := CaptureOutput(func() {
		err := game.Setup(file, TestNumAliens)
		assert.Nil(t, err)
	})

	// Confirm that the output contain's the expected value
	assert.Contains(t, output, "The World:")
}

func TestPlay(t *testing.T) {
	// Load our generic file
	file := LoadFilePath()

	Init()

	err := game.Setup(file, TestNumAliens)
	assert.Nil(t, err)

	// Capture Invade's output for testing
	output := CaptureOutput(func() {
		err = game.Play()
		assert.Nil(t, err)
	})

	// Confirm that the output contain's the expected values
	assert.Contains(t, output, "The game has started!")
	assert.Contains(t, output, "Game completed on turn")
}
