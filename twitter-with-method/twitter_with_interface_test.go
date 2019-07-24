package twitter_test

import (
	"github.com/GGYaX/namecheck"
	twitter2 "github.com/GGYaX/namecheck/twitter-with-method"
	"github.com/stretchr/testify/assert"
	_ "os"
	"testing"
)

var validator namecheck.Validater = &twitter2.Twitter{}

func TestValidate2(t *testing.T) {
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
			assert.Equal(t, test.out, validator.Validate(test.in))
		})
	}
}
