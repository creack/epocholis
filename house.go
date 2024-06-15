package main

import "fmt"

type house struct {
	base
	worker *worker
}

func newHouse(base base) *house {
	h := &house{base: base}
	h.set(h)
	return h
}

func (h *house) render() string {
	return "H"
}

func (h *house) tick() {
	if h.worker != nil {
		h.worker.tick()
		return
	}

	r := randElem(filter[*road](h.neighs()))
	if r == nil {
		panic(fmt.Errorf("no road found"))
	}
	h.worker = newWorker(r.base)
	r.worker = h.worker
}
