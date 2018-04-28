package chap01

// Dollar ...
type Dollar struct {
	Ammount int
}

// Times ...
func (d *Dollar) Times(num int) {
	// FIXME
	d.Ammount = 10
}
