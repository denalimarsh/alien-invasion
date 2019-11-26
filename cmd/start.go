package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/denalimarsh/invasion/game"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const flagInFile = "file"
const flagNumberAliens = "numAliens"
const flagAdvancedTech = "advancedTech"

// startCmd : represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "starts the extraterrestrial invasion of a world",
	Example: "invasion --file=\"./assets/small_world.txt\" --numAliens=",
	Run: func(cmd *cobra.Command, args []string) {
		// Parse flags
		filePath := viper.GetString(flagInFile)
		if strings.TrimSpace(filePath) == "" {
			log.Fatal("invalid --file flag")
		}

		numAliens := viper.GetInt(flagNumberAliens)
		if numAliens <= 0 {
			log.Fatal("invalid --numAliens flag")
		}

		advancedTechStr := viper.GetString(flagAdvancedTech)
		advancedTech, err := strconv.ParseBool(strings.TrimSpace(advancedTechStr))
		if err != nil {
			log.Fatal("invalid --advancedTech flag")
		}

		// Initializes a new world
		game.Init(advancedTech)

		// Generates World's Cities, Paths, and Aliens
		err = game.Setup(filePath, numAliens)
		if err != nil {
			log.Fatal(err)
		}

		// Executes invasion as a 2-phase turn based game
		game.Play()
	},
}

// init : prepares required flags and adds them to the start cmd
func init() {
	rootCmd.AddCommand(startCmd)

	// Add flags and mark as required
	startCmd.PersistentFlags().String(flagInFile, "./assets/small_world.txt", "Path to start file")
	startCmd.PersistentFlags().Int(flagNumberAliens, 10, "Number of aliens participating in the invasion")
	startCmd.PersistentFlags().Bool(flagAdvancedTech, false, "Determines the alien army's technical capabilities")
	startCmd.MarkFlagRequired(flagInFile)
	startCmd.MarkFlagRequired(flagNumberAliens)
	startCmd.MarkFlagRequired(flagAdvancedTech)

	// Bind flags
	viper.BindPFlag(flagInFile, startCmd.Flag(flagInFile))
	viper.BindPFlag(flagNumberAliens, startCmd.Flag(flagNumberAliens))
	viper.BindPFlag(flagAdvancedTech, startCmd.Flag(flagAdvancedTech))
}
