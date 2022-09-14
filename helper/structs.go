package helper

const (
	Add int = iota
	Remove
	Change
)

type Delta struct {
	changes []string
}

type file struct {
	size    int
	content []byte
}
