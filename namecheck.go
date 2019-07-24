package namecheck

import (
	"github.com/GGYaX/namecheck/twitter-with-method"
)

type Validater interface {
	Validate(s string) []string
}

func SocialNetworks() []Validater {
	return []Validater{&twitter.Twitter{}}
}
