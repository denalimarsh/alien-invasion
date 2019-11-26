package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	// Confirm that the game has been instantiated
	assert.NotNil(t, game.getWorld())
	assert.NotNil(t, game.getRandSeed())
}

func TestSetup(t *testing.T) {
	// Load our generic file
	file := LoadFilePath()

	game := NewGame()

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

	game := NewGame()

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
