package helper

const (
	Add int = iota
	Remove
	Change
)

type Delta struct {
	Changes []string
}

func NewDelta() Delta {
	return Delta{
		Changes: []string{},
	}
}

type file struct {
	size    int
	content []byte
}
