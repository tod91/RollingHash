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

func (d *Delta) Add(cmd string) {
	d.Changes = append(d.Changes, cmd)
}

type file struct {
	size    int
	content []byte
}
