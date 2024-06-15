package main

type road struct {
	base
	worker *worker
}

func newRoad(base base) *road {
	r := &road{base: base}
	r.set(r)
	return r
}

func (r *road) render() string {
	if r.worker != nil {
		return "W"
	}
	return "-"
}
