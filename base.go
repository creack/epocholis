package main

type base struct {
	grid grid
	x, y int
}

func (c *controller) newBase(x, y int) base {
	return base{grid: c.grid, x: x, y: y}
}

func (b *base) get() gridCase {
	return b.grid[b.y][b.x]
}

func (b *base) set(elem gridCase) {
	b.grid[b.y][b.x] = elem
}

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

func filter[T any](in []gridCase) []T {
	var out []T
	for _, elem := range in {
		if tmp, ok := elem.(T); ok {
			out = append(out, tmp)
		}
	}
	return out
}

// TODO: Implement this.
func randElem[T any](in []*T) *T {
	if len(in) == 0 {
		return nil
	}
	return in[0]
}
