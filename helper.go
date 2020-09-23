package geofence

import (
	"math"
	"math/big"
)

type direction float64

const (
	CW  direction = 1
	EQ  direction = 0
	CCW direction = -1
)

type point [2]*big.Float
type vector [2]point

var zero = big.NewFloat(0)

func toPoint(x, y float64) point {
	return point{big.NewFloat(x), big.NewFloat(y)}
}

func vectorDirection(p point, v vector) direction {
	x1, y1, x2, y2 := v[0][0], v[0][1], v[1][0], v[1][1]
	xn, yn := p[0], p[1]
	n := new(big.Float).Sub(
		new(big.Float).Mul(
			new(big.Float).Sub(yn, y1),
			new(big.Float).Sub(x2, x1),
		),
		new(big.Float).Mul(
			new(big.Float).Sub(xn, x1),
			new(big.Float).Sub(y2, y1),
		),
	)
	switch {
	case n.Cmp(zero) == -1:
		return CW
	case n.Cmp(zero) == 1:
		return CCW
	}
	return EQ
}

func normVector(v vector) vector {
	x1, y1, x2, y2 := v[0][0], v[0][1], v[1][0], v[1][1]
	n, m := new(big.Float).Sub(x1, zero), new(big.Float).Sub(y1, zero)
	return vector{
		{new(big.Float).Sub(x1, n), new(big.Float).Sub(y1, m)},
		{new(big.Float).Sub(x2, n), new(big.Float).Sub(y2, m)},
	}
}

func findAngle(normVectorA, normVectorB vector) float64 {
	xA, yA, xB, yB := normVectorA[1][0], normVectorA[1][1], normVectorB[1][0], normVectorB[1][1]
	dotProduct := new(big.Float).Add(new(big.Float).Mul(xA, xB), new(big.Float).Mul(yA, yB))
	magA := new(big.Float).Sqrt(
		new(big.Float).Add(
			new(big.Float).Mul(xA, xA),
			new(big.Float).Mul(yA, yA),
		),
	)
	magB := new(big.Float).Sqrt(
		new(big.Float).Add(
			new(big.Float).Mul(xB, xB),
			new(big.Float).Mul(yB, yB),
		),
	)
	cosTeta := new(big.Float).Quo(dotProduct, new(big.Float).Mul(magA, magB))
	cosTetaF64, _ := cosTeta.Float64()
	teta := math.Acos(cosTetaF64)
	// return teta * 180 / math.Pi
	return teta
}
