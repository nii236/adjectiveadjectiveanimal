package main

import (
	"fmt"
	"strings"

	aaa "github.com/nii236/adjectiveadjectiveanimal"
)

func main() {
	result := strings.Join(aaa.Generate(2, &aaa.Options{}), "-")
	fmt.Println(result)
}
