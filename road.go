package main

type road struct {
	base
	worker *worker
}

func (c *controller) newRoad(x, y int) *road {
	r := &road{base: c.newBase(x, y)}
	r.set(r)
	return r
}

func (r *road) render() string {
	if r.worker != nil {
		return r.worker.render()
	}
	var (
		_, north = r.getNeigh(directionNorth).(*road)
		_, south = r.getNeigh(directionSouth).(*road)
		_, east  = r.getNeigh(directionEast).(*road)
		_, west  = r.getNeigh(directionWest).(*road)
	)
	switch {
	case (east || west) && !north && !south:
		return "─"
	case (north || south) && !east && !west:
		return "│"
	case !north && south && east && !west:
		return "┌"
	case !north && south && !east && west:
		return "┐"
	case north && !south && east && !west:
		return "└"
	case north && !south && !east && west:
		return "┘"
	case north && !south && east && west:
		return "┴"
	case !north && south && east && west:
		return "┬"
	case north && south && !east && west:
		return "┤"
	case north && south && east && !west:
		return "├"
	case north && south && east && west:
		return "┼"
	default:
		return "?"
	}
}
