package namecheck_test

import (
	"github.com/GGYaX/namecheck/namecheck"
	"github.com/GGYaX/namecheck/twitter"
	"github.com/stretchr/testify/assert"
	"testing"
)

var validater namecheck.Validater = &twitter.Twitter{}
var isAvailable namecheck.IsAvailabler = &twitter.Twitter{}

func TestValidate(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{in: "qqqq", out: []string{}},
		{in: "qqqqtwitterqqqqqqqqqqqqqqqqqq", out: []string{"Is not short enough", "ContainsNoIllegalPattern"}},
		{in: "qwerqewrwcv12332_#", out: []string{"Is not short enough", "Not OnlyContainsLegalRunes"}},
		{in: "qwerqewrwsadfasdfasdfatwittercv12332_#", out: []string{"Is not short enough", "ContainsNoIllegalPattern", "Not OnlyContainsLegalRunes"}},
		{in: "", out: []string{"Is not long enough", "Not OnlyContainsLegalRunes"}},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			validater.Validate(test.in)
			assert.Equal(t, test.out, validater.Validate(test.in))
		})
	}
}

func TestSocialNetworks(t *testing.T) {
	//assert.Equal(t, &twitter.Twitter{}, SocialNetworks())
}

func TestCheckTwitterName(t *testing.T) {
	existed, err := isAvailable.IsAvailable("abc")
	assert.Equal(t, true, existed)
	assert.Zero(t, err)
	existed, err = isAvailable.IsAvailable("kcvaekjhrkq")
	assert.Equal(t, false, existed)
	assert.Zero(t, err)
}
