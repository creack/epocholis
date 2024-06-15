package main

import "fmt"

type direction int

const (
	_ direction = iota
	directionNorth
	directionSouth
	directionEast
	directionWest
)

//nolint:gochecknoglobals // Expected global.
var allDirections = []direction{directionNorth, directionSouth, directionEast, directionWest}

type worker struct {
	base
	direction direction

	house    *house
	curRoad  *road
	prevRoad *road

	waitTick int
}

func newWorker(road *road, house *house) *worker {
	return &worker{
		base:      road.base,
		direction: directionEast,
		curRoad:   road,
		prevRoad:  nil,
	}
}

// NOTE: The caller must check that the move is valid before calling.
func (w *worker) move(dir direction) {
	w.prevRoad = w.get().(*road)
	w.prevRoad.worker = nil
	w.direction = dir
	switch dir {
	case directionNorth:
		w.y--
	case directionSouth:
		w.y++
	case directionEast:
		w.x++
	case directionWest:
		w.x--
	default:
		panic(fmt.Errorf("unknown direction %d", dir))
	}
	w.curRoad = w.get().(*road)
	w.curRoad.worker = w
}

func (w *worker) lookupDirection(neigh *road) direction {
	for _, dir := range allDirections {
		r, ok := w.getNeigh(dir).(*road)
		if ok && r == neigh {
			return dir
		}
	}
	return 0
}

func (w *worker) tick() {
	if w.waitTick > 0 {
		w.waitTick--
		return
	}
	w.waitTick = 1

	for _, dir := range allDirections {
		if dir != w.direction {
			continue
		}
		neighRoads := filterOutSingleElem(filterType[*road](w.neighs()), w.prevRoad)
		if len(neighRoads) == 0 {
			neighRoads = append(neighRoads, w.prevRoad)
		}
		w.move(w.lookupDirection(randElem(neighRoads)))
	}
}
