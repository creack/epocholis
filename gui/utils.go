package main

func cartesianToIso(tileSize, x, y float64) (float64, float64) {
	ix := (x - y) * float64(tileSize/2)
	iy := (x + y) * float64(tileSize/4)
	return ix, iy
}

func isoToCartesian(tileSize, x, y float64) (float64, float64) {
	cx := (x/float64(tileSize/2) + y/float64(tileSize/4)) / 2
	cy := (y/float64(tileSize/4) - (x / float64(tileSize/2))) / 2
	return cx, cy
}

func lerp(x1, y1, x2, y2, t float64) (x3, y3 float64) {
	x1, y1 = x1*(1.0-t), y1*(1.0-t)
	x2, y2 = x2*t, y2*t
	return x1 + x2, y1 + y2
}
