package aaa_test

import (
	"strings"
	"testing"

	"github.com/nii236/gopher-resume/aaa"
)

func TestGenerate(t *testing.T) {
	for i := 1; i < 2; i++ {
		for j := 0; j < 1000; j++ {
			token := aaa.Generate(i, "-")
			expectedDashes := i
			gotDashes := strings.Count(token, "-")
			if expectedDashes != gotDashes {
				t.Errorf("Wrong number of dashes: Got %d, expected %d", gotDashes, expectedDashes)
			}
			if token == "" {
				t.Error("Empty token returned")
			}
		}
	}
}
