package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Result struct {
	ImagePath string
	Count     int
}

var rootCmd = &cobra.Command{
	Use:   "ocr-index",
	Short: "Index and search the text in images",
	Long:  `A naive exercise in OCR with Tesseract, and indexing in a simple key-value store`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
