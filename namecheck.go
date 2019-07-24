package namecheck

type Validater interface {
	Validate(s string) []string
}
