package main

import (
	"fmt"
	"os"
	"strings"

	// "github.com/otiai10/gosseract"
	"github.com/rapidloop/skv"
)

type Result struct {
	ImagePath string
	Count     int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("I need an image to index!")
		os.Exit(1)
	}

	// imgPath := os.Args[1]

	// ocrClient := gosseract.NewClient()
	// defer ocrClient.Close()

	// ocrClient.SetImage(imgPath)
	// text, err := ocrClient.Text()

	// var counts map[string]int

	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	os.Exit(1)
	// }

	// counts = wordCounts(text)

	// fmt.Println("Word counts:", counts)

	// err = save(imgPath, counts)
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	os.Exit(1)
	// }

	term := os.Args[1]

	err, results := search(term)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	for _, r := range results {
		fmt.Printf("%d occurrences in %s\n", r.Count, r.ImagePath)
	}
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
