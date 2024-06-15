package main

import "fmt"

type house struct {
	base
	worker *worker

	water         int
	requiredWater int
}

func (c *controller) newHouse(x, y int) *house {
	h := &house{
		base:          c.newBase(x, y),
		requiredWater: 1,
	}
	h.set(h)
	return h
}

func (h *house) render() string {
	if h.water < h.requiredWater {
		return "\033[38;2;255;0;0;1;4m⌂\033[0m"
	}
	return "\033[38;2;175;238;238;1;4m⌂\033[0m"
}

func (h *house) consumeResource(wo *worker) {
	if h.water >= h.requiredWater {
		return
	}
	if wo.resourceType != resourceTypeWater {
		return
	}
	if wo.resourceCount < 1 {
		return
	}
	wo.resourceCount--
	h.water++
}

func (h *house) tick() {
	if h.worker != nil {
		h.worker.tick()
		return
	}

	r := randElem(filterType[*road](h.neighs()))
	if r == nil {
		panic(fmt.Errorf("no road found"))
	}
	h.worker = newWorker(r)
	r.worker = h.worker

	h.worker.resourceType = resourceTypeEmployee
	h.worker.resourceCount = 1
}
