package cmd

import (
	"fmt"
	"strings"

	"github.com/otiai10/gosseract"
	"github.com/rapidloop/skv"
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index <image_file>",
	Args:  cobra.ExactArgs(1),
	Short: "Add an image to the search index",
	Long: `Takes an image file path, runs Optical Character
Recognition (OCR) on it, and indexes the words it finds.

Assumes that the file won't move from the path. When we look
up words in the index, the result will be the file path given
to this command.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("index called")

		imgPath := args[0]

		ocrClient := gosseract.NewClient()
		defer ocrClient.Close()

		ocrClient.SetImage(imgPath)
		text, err := ocrClient.Text()

		if err != nil {
			return
		}

		counts := wordCounts(text)

		fmt.Println("Word counts:", counts)

		err = save(imgPath, counts)
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// indexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// indexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func save(imgPath string, counts map[string]int) error {
	store, err := skv.Open("index.db")

	if err != nil {
		return err
	}
	defer store.Close()

	for word, count := range counts {
		// err := store.Get(word, &results)

		// if results == nil {
		results := []Result{Result{imgPath, count}}

		err = store.Put(word, results)

		if err != nil {
			return err
		}
		// }

	}

	return nil
}

func wordCounts(text string) (counts map[string]int) {
	words := strings.Fields(text)
	counts = make(map[string]int)

	for i := 0; i < len(words); i++ {
		counts[words[i]]++
	}

	return
}
