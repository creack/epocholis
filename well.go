package main

import "fmt"

type well struct {
	base

	worker *worker

	requiredEmployeeCount int
	employees             int

	waitTick int
}

func (c *controller) newWell(x, y int) *well {
	we := &well{
		base:                  c.newBase(x, y),
		requiredEmployeeCount: 1,
	}
	we.set(we)
	return we
}

func (w *well) render() string {
	if w.employees < w.requiredEmployeeCount {
		return "\033[38;2;255;0;0;1;4m☂\033[0m"
	}
	return "\033[38;2;0;255;0;1;4m☂\033[0m"
}

func (w *well) tick() {
	if w.employees < w.requiredEmployeeCount {
		return
	}
	if w.waitTick > 0 {
		w.waitTick--
		return
	}
	if w.worker != nil {
		w.worker.tick()
		return
	}
	w.waitTick = 30

	r := randElem(filterType[*road](w.neighs()))
	if r == nil {
		panic(fmt.Errorf("no road found"))
	}
	w.worker = newWorker(r)
	r.worker = w.worker

	w.worker.color = "175;238;238;1"
	w.worker.resourceType = resourceTypeWater
	w.worker.resourceCount = 1
}

func (w *well) consumeResource(wo *worker) {
	if w.employees >= w.requiredEmployeeCount {
		return
	}
	if wo.resourceType != resourceTypeEmployee {
		return
	}
	if wo.resourceCount < 1 {
		return
	}
	wo.resourceCount--
	w.employees++
}
