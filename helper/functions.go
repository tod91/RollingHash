package helper

func SanityCheck(first, second string) bool {
	return first == second
}

func (d *Delta) Add(cmd string) {
	d.Changes = append(d.Changes, cmd)
}
