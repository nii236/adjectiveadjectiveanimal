package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	aaa "github.com/nii236/adjectiveadjectiveanimal"
)

func main() {
	args := os.Args[1:]
	i := 2
	if len(args) > 0 {
		var err error
		i, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Non number supplied for adjective amount. Exiting.")
			os.Exit(-1)
		}
	}
	result := strings.Join(aaa.Generate(i, &aaa.Options{}), "-")
	fmt.Println(result)
}
