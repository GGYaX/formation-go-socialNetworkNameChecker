package twitter_test

import (
	twitter2 "github.com/GGYaX/namecheck/twitter-with-method"
	"github.com/stretchr/testify/assert"
	_ "os"
	"testing"
)

var twitter *twitter2.Twitter

func TestIsShortEnough(t *testing.T) {
	assert.Equal(t, true, twitter.IsShortEnough("qqqqqqqqqqqqqqq"))
	assert.Equal(t, false, twitter.IsShortEnough("qqqqqqqqqqqqqqqq"))
	assert.Equal(t, true, twitter.IsShortEnough("我我我我我我我我我我我我我😊我"))
	assert.Equal(t, false, twitter.IsShortEnough("我我我我我我我我我我我我我😊我我"))
}

func TestContainsNoIllegalPattern(t *testing.T) {
	assert.Equal(t, true, twitter.ContainsNoIllegalPattern("twitter"))
	assert.Equal(t, true, twitter.ContainsNoIllegalPattern("atwitterb"))
	assert.Equal(t, false, twitter.ContainsNoIllegalPattern("asdfadsf"))
}

func TestOnlyContainsLegalRunes(t *testing.T) {
	assert.Equal(t, true, twitter.OnlyContainsLegalRunes("qqqqqqqqqqqqqqq"))
	assert.Equal(t, true, twitter.OnlyContainsLegalRunes("qqqqqqqqqqqqqqqq"))
	assert.Equal(t, false, twitter.OnlyContainsLegalRunes("我我我我我我我我我我我我我😊我"))
}

func TestIsLongEnough(t *testing.T) {
	assert.Equal(t, true, twitter.IsLongEnough("q"))
	assert.Equal(t, false, twitter.IsLongEnough(""))
}

func BenchmarkIsLongEnough(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twitter.IsLongEnough("asdf")
	}
}

func TestTableIsShortEnough(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{in: "qqqqqqqqqqqqqqq", out: true},
		{in: "qqqqqqqqqqqqqqqq", out: false},
		{in: "我我我我我我我我我我我我我😊我", out: true},
		{in: "我我我我我我我我我我我我我😊我我", out: false},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			assert.Equal(t, test.out, twitter.IsShortEnough(test.in))
		})
	}
}

func TestTableContainsNoIllegalPattern(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{in: "twitter", out: true},
		{in: "atwitterb", out: true},
		{in: "asdfadsf", out: false},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			assert.Equal(t, test.out, twitter.ContainsNoIllegalPattern(test.in))
		})
	}
}

func TestTableOnlyContainsLegalRunes(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{in: "qqqqqqqqqqqqqqq", out: true},
		{in: "qqqqqqqqqqqqqqqq", out: true},
		{in: "我我我我我我我我我我我我我😊我", out: false},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			assert.Equal(t, test.out, twitter.OnlyContainsLegalRunes(test.in))
		})
	}
}

func TestTableIsLongEnough(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{in: "q", out: true},
		{in: "", out: false},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			assert.Equal(t, test.out, twitter.IsLongEnough(test.in))
		})
	}
}

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
			twitter.Validate(test.in)
			assert.Equal(t, test.out, twitter.Validate(test.in))
		})
	}
}
