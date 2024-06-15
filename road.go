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
