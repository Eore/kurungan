package geofence

import (
	"fmt"
	"math/big"
)

type Point [2]*big.Float
type State int

const (
	Outside State = 0
	Inside  State = 1
)

func NewPoint(x, y float64) Point {
	return Point{big.NewFloat(x), big.NewFloat(y)}
}

func GeoFenceCheck(p Point, polygon []Point) (State, error) {
	l := len(polygon)
	if !isPointSame(polygon[0], polygon[l-1]) {
		return Outside, fmt.Errorf("Polygon not closing")
	}
	if l < 3 {
		return Outside, fmt.Errorf("Is not polygon")
	}
	state := false
	x, y := p[0], p[1]
	for i := 0; i < l-1; i++ {
		x1, y1 := polygon[i][0], polygon[i][1]
		x2, y2 := polygon[i+1][0], polygon[i+1][1]
		if y1.Cmp(y) != y2.Cmp(y) &&
			x.Cmp(
				new(big.Float).Add(
					x1,
					new(big.Float).Mul(
						new(big.Float).Sub(x2, x1),
						new(big.Float).Quo(
							new(big.Float).Sub(y, y1),
							new(big.Float).Sub(y2, y1),
						),
					),
				),
			) == -1 {
			state = !state
		}
	}
	if state {
		return Inside, nil
	}
	return Outside, nil
}

func isPointSame(p1, p2 Point) bool {
	return p1[0].Cmp(p2[0]) == 0 && p1[1].Cmp(p2[1]) == 0
}
