package main

import "fmt"

type workerDirection int

const (
	_ workerDirection = iota
	workerDirectionNorth
	workerDirectionSouth
	workerDirectionEast
	workerDirectionWest
)

type worker struct {
	base
	direction workerDirection

	waitTick int
}

func newWorker(base base) *worker {
	return &worker{
		base:      base,
		direction: workerDirectionEast,
	}
}

func (w *worker) tick() {
	if w.waitTick > 0 {
		w.waitTick--
		return
	}
	w.waitTick = 5
	w.get().(*road).worker = nil
	switch w.direction {
	case workerDirectionEast:
		if w.x+1 < len(w.grid[0]) {
			r, ok := w.grid[w.y][w.x+1].(*road)
			if ok {
				r.worker = w
				w.x++
				return
			}
		}
		w.direction = workerDirectionWest
	case workerDirectionWest:
		if w.x-1 >= 0 {
			r, ok := w.grid[w.y][w.x-1].(*road)
			if ok {
				r.worker = w
				w.x--
				return
			}
		}
		w.direction = workerDirectionEast
	default:
		panic(fmt.Errorf("not implemented: %d", w.direction))

	}
}
