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

	curRoad  *road
	prevRoad *road

	waitTick int

	resourceType  resourceType
	resourceCount int

	color string
}

func newWorker(road *road) *worker {
	return &worker{
		base:      road.base,
		direction: directionEast,
		curRoad:   road,
		prevRoad:  nil,
	}
}

func (w *worker) render() string {
	if w.color == "" {
		return "W"
	}
	return "\033[38;2;" + w.color + "mW\033[0m"
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
	if w.resourceCount == 0 {
		w.get().(*road).worker = nil
		return
	}
	// Move.
	w.actionMove()
	// Perform actions.
	for _, rc := range filterType[resourceConsumer](w.neighs()) {
		rc.consumeResource(w)
		if w.resourceCount <= 0 {
			break
		}
	}
}

func (w *worker) actionMove() {
	w.waitTick = 3
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
