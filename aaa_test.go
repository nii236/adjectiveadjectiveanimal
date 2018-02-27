package aaa_test

import (
	"testing"

	"github.com/nii236/adjectiveadjectiveanimal"
)

func TestGenerateOneToken(t *testing.T) {
	token := aaa.Generate(2, &aaa.Options{})
	t.Log(token)

}
