package aaa

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"math/big"
	mrand "math/rand"
	"strings"
)

// Options are the title available for the AAA generator
type Options struct {
	UseSeed bool
	Seed    []byte
	// Title will return the words in title form
	Title bool
}

// Generate will generate a friendly token of a numAdj
// adjectives followed by an animal separated by sep
func Generate(numAdj int, options *Options) []string {
	if options.UseSeed {
		return withSeed(numAdj, options)
	}

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

func withSeed(numAdj int, options *Options) []string {
	result := []string{}

	h := md5.New()
	h.Write(options.Seed)
	seed := int64(binary.BigEndian.Uint64(h.Sum(nil)))

	mrand.NewSource(seed)
	for i := 0; i < numAdj; i++ {
		n := mrand.Intn(len(adjectives) - 1)
		rndAdj := strings.ToLower(adjectives[n])
		result = append(result, rndAdj)
	}

	n := mrand.Intn(len(animals) - 1)
	rndAni := strings.ToLower(animals[n])
	result = append(result, rndAni)
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
