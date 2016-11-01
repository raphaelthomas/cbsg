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

	var adverbStatic string
	var verbStatic string
	var nounStatic string
	var adjectiveStatic string

	flag.StringVar(&adverbStatic, "adverb", "", "use this adverb")
	flag.StringVar(&verbStatic, "verb", "", "use this verb")
	flag.StringVar(&adjectiveStatic, "adjective", "", "use this adjective")
	flag.StringVar(&nounStatic, "noun", "", "use this noun")

	flag.Parse()

	wordsFile, err := os.Open(*wordsFileName)
	if err != nil {
		fmt.Printf("Error reading words file %s", *wordsFileName)
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(wordsFile)
	if err = jsonParser.Decode(&words); err != nil {
		panic(err)
	}

	for i := 0; i < *bullshitCount; i++ {
		rand.Seed(time.Now().UnixNano())

		var adverb string = adverbStatic
		var verb string = verbStatic
		var adjective string = adjectiveStatic
		var noun string = nounStatic

		if adverb == "" {
			adverb = words.Adverbs[rand.Intn(len(words.Adverbs)-1)]
		}
		if verb == "" {
			verb = words.Verbs[rand.Intn(len(words.Verbs)-1)]
		}
		if adjective == "" {
			adjective = words.Adjectives[rand.Intn(len(words.Adjectives)-1)]
		}
		if noun == "" {
			noun = words.Nouns[rand.Intn(len(words.Nouns)-1)]
		}

		fmt.Printf("%s %s %s %s\n", adverb, verb, adjective, noun)
	}

	return
}
