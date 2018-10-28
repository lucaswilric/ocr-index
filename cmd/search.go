package cmd

import (
	"fmt"

	"github.com/rapidloop/skv"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <word>",
	Args:  cobra.ExactArgs(1),
	Short: "Search for images that contain a word",
	Long: `This very naive index can only find single, complete, unmodified words.
And if the OCR generated garbage, you're outta luck.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")

		term := args[0]

		err, results := search(term)
		if err != nil {
			return
		}

		for _, r := range results {
			fmt.Printf("%d occurrences in %s\n", r.Count, r.ImagePath)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func search(word string) (error, []Result) {
	store, err := skv.Open("index.db")

	if err != nil {
		return err, nil
	}
	defer store.Close()

	var results []Result
	err = store.Get(word, &results)

	if err != nil {
		return err, nil
	}

	return nil, results
}
