package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/denalimarsh/invasion/game"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const FlagInFile = "file"
const FlagNumberAliens = "N"

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "starts the extraterrestrial invasion of a world",
	Example: "invasion --file=\"./assets/world.txt\" --N=", // TODO: add an example file to bin/world.txt
	Run: func(cmd *cobra.Command, args []string) {
		// Parse flags
		filePath := viper.GetString(FlagInFile)
		if strings.TrimSpace(filePath) == "" {
			log.Fatal("invalid file path")
		}

		numAliens := viper.GetInt(FlagNumberAliens)
		if numAliens <= 0 {
			log.Fatal("invalid number of aliens")
		}

		fmt.Println("filePath: ", filePath)
		fmt.Println("numAliens: ", numAliens)

		log.Printf("Setting up the game...")
		err := game.Setup(filePath, numAliens)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Playing the game...")
		err = game.Play()
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Add flags and mark as required
	startCmd.PersistentFlags().String(FlagInFile, "", "Path to start file")
	startCmd.PersistentFlags().Int(FlagNumberAliens, 10, "Number of aliens participating in the invasion")
	startCmd.MarkFlagRequired(FlagInFile)
	startCmd.MarkFlagRequired(FlagNumberAliens)

	// Bind flags
	viper.BindPFlag(FlagInFile, startCmd.Flag(FlagInFile))
	viper.BindPFlag(FlagNumberAliens, startCmd.Flag(FlagNumberAliens))
}
