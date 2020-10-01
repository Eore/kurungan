package geofence

import (
	"math/big"
)

type point [2]*big.Float
type vector [2]point
type triangle [3]vector

var zero = big.NewFloat(0)

func pointInTriangle(p point, tri triangle) bool {
	sign := func(p point, v vector) *big.Float {
		x, y := p[0], p[1]
		x1, y1 := v[0][0], v[0][1]
		x2, y2 := v[1][0], v[1][1]

		return new(big.Float).Sub(
			new(big.Float).Mul(
				new(big.Float).Sub(y, y1),
				new(big.Float).Sub(x2, x1),
			),
			new(big.Float).Mul(
				new(big.Float).Sub(x, x1),
				new(big.Float).Sub(y2, y1),
			),
		)
	}

	check := func(p point, tri triangle) bool {
		c1 := sign(p, tri[0]).Cmp(zero)
		c2 := sign(p, tri[1]).Cmp(zero)
		c3 := sign(p, tri[2]).Cmp(zero)
		return !(c1 < 0 || c2 < 0 || c3 < 0)
	}

	return check(p, tri)
}

func toPoint(x, y float64) point {
	return point{big.NewFloat(x), big.NewFloat(y)}
}
