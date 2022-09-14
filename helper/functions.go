package helper

func SanityCheck(first, second string) bool {
	return first == second
}

func (d *Delta) Add(cmd string) {
	d.changes = append(d.changes, cmd)
}
