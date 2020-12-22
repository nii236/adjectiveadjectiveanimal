[![GoDoc](https://godoc.org/github.com/nii236/adjectiveadjectiveanimal?status.svg)](https://godoc.org/github.com/nii236/adjectiveadjectiveanimal)

# Adjective Adjective Animal

This is a port of [this node.js project](https://github.com/a-type/adjective-adjective-animal/).

It will generate a specified number of adjectives, then an animal at the end. The words are separated by a character of your choice.

# Installation

```bash
$ go get -u github.com/nii236/adjectiveadjectiveanimal/cmd/aaa
```

# Example Usage

```
package main

import (
	"fmt"

	aaa "github.com/nii236/adjectiveadjectiveanimal"
)

func main() {
	result := aaa.Generate(3, "-")
	fmt.Println(result)
}
```

```
$ aaa
unprotected-ultramicroscopic-foolproof-millipede
```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
