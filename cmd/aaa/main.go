package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	aaa "github.com/nii236/adjectiveadjectiveanimal"
)

func main() {
	// unique seed
	rand.Seed(time.Now().Local().UnixNano())

	var alen int
	var rUp bool
	var rNum bool

	flag.IntVar(&alen, "len", 2, "how many adjective")
	flag.BoolVar(&rUp, "up", false, "random upper case")
	flag.BoolVar(&rNum, "num", false, "random number (replace)")
	flag.Parse()

	if alen < 2 {
		fmt.Println("need at least 2 adjective")
		return
	}

	// generate words first
	result := strings.Join(aaa.Generate(alen, &aaa.Options{}), "-")

	// randomise upper case
	if rUp {
		dat := []rune(result)
		for i := 0; i < 3; i++ {
			r := rand.Intn(len(result))
			x := string(dat[r])
			dat[r] = []rune(strings.ToUpper(x))[0]
		}
		result = string(dat)
	}

	// randomise number
	if rNum {
		dat := []rune(result)
		for i := 0; i < 2; i++ {
			r := rand.Intn(len(result))
			if dat[r] == rune('-') {
				continue
			}
			n := rand.Intn(10)
			// fmt.Println(i, n)
			dat[r] = []rune(strconv.Itoa(n))[0]
		}
		result = string(dat)
	}

	fmt.Println(result)
}
