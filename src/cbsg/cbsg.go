package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var words struct {
	Adverbs    []string
	Verbs      []string
	Adjectives []string
	Nouns      []string
}

func main() {
	var wordsFileName = flag.String("words", "", "JSON file for words")
	var bullshitCount = flag.Int("n", 1, "number of BS sentences to generate")

	flag.Parse()

	wordsFile, err := os.Open(*wordsFileName)
	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(wordsFile)
	if err = jsonParser.Decode(&words); err != nil {
		panic(err)
	}

	for i := 0; i < *bullshitCount; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Printf("%s %s %s %s\n", words.Adverbs[rand.Intn(len(words.Adverbs)-1)], words.Verbs[rand.Intn(len(words.Verbs)-1)], words.Adjectives[rand.Intn(len(words.Adjectives)-1)], words.Nouns[rand.Intn(len(words.Nouns)-1)])
	}

	return
}
