package main

import (
	"fmt"
	"math/rand/v2"
)

type base struct {
	grid grid
	x, y int
}

func (c *controller) newBase(x, y int) base {
	return base{grid: c.grid, x: x, y: y}
}

func (b *base) get() gridCase { return b.grid[b.y][b.x] }

func (b *base) getNeigh(dir direction) gridCase {
	switch dir {
	case directionNorth:
		if b.y-1 < 0 {
			return nil
		}
		return b.grid[b.y-1][b.x]
	case directionSouth:
		if b.y+1 >= len(b.grid) {
			return nil
		}
		return b.grid[b.y+1][b.x]
	case directionEast:
		if b.x+1 >= len(b.grid[0]) {
			return nil
		}
		return b.grid[b.y][b.x+1]
	case directionWest:
		if b.x-1 < 0 {
			return nil
		}
		return b.grid[b.y][b.x-1]
	default:
		panic(fmt.Errorf("unknown direction %d", dir))
	}
}

func (b *base) set(elem gridCase) { b.grid[b.y][b.x] = elem }

func (b *base) neighs() []gridCase {
	var out []gridCase
	if b.x-1 >= 0 {
		out = append(out, b.grid[b.y][b.x-1])
	}
	if b.y-1 >= 0 {
		out = append(out, b.grid[b.y-1][b.x])
	}
	if b.x+1 < len(b.grid[0]) {
		out = append(out, b.grid[b.y][b.x+1])
	}
	if b.y+1 < len(b.grid) {
		out = append(out, b.grid[b.y+1][b.x])
	}
	return out
}

func filterType[T any](in []gridCase) []T {
	var out []T
	for _, elem := range in {
		if tmp, ok := elem.(T); ok {
			out = append(out, tmp)
		}
	}
	return out
}

func randElem[T any](in []*T) *T {
	if len(in) == 0 {
		return nil
	}
	if len(in) == 1 {
		return in[0]
	}
	return in[randRange(0, len(in))]
}

func filterOutSingleElem[T any](in []*T, f *T) []*T {
	var out []*T
	for _, elem := range in {
		if elem != f {
			out = append(out, elem)
		}
	}
	return out
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
