package namecheck

type Validater interface {
	Validate(s string) []string
}

type IsAvailabler interface {
	IsAvailable(s string) (bool, IsAvailablerError)
}

type SocialNetworkChecker interface {
	Validater
	IsAvailabler
}

//func SocialNetworks() []Validater {
//	return []Validater{&twitter.Twitter{}}
//}

type IsAvailablerError struct {
	E error
}
