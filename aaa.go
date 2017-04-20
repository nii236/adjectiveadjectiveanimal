package aaa

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Options are the title available for the AAA generator
type Options struct {
	// Title will return the words in title form
	Title bool
}

// Generate will generate a friendly token of a numAdj
// adjectives followed by an animal separated by sep
func Generate(numAdj int, options *Options) []string {
	result := []string{}
	for i := 0; i < numAdj; i++ {
		result = append(result, randomAdjective())
	}
	result = append(result, randomAnimal())
	if options.Title {
		titledResult := []string{}
		for _, word := range result {
			titledResult = append(titledResult, strings.Title(word))
		}
		return titledResult
	}
	return result
}

func randomAdjective() string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(adjectives)-1)))
	if err != nil {
		return ""
	}
	return strings.ToLower(adjectives[int(n.Int64())])
}

func randomAnimal() string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(animals)-1)))
	if err != nil {
		return ""

	}
	return strings.ToLower(animals[int(n.Int64())])
}
