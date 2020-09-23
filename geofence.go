package geofence

import "fmt"

func GeoFenceCheck(p point, polygon []point) (bool, error) {
	var total float32 = 0
	nPolygonPoints := len(polygon)
	if !isPointSame(polygon[0], polygon[nPolygonPoints-1]) {
		return false, fmt.Errorf("Polygon its not closed")
	}
	for i := 0; i < nPolygonPoints; i++ {
		if i > 0 {
			total = total + findAngleWithDirection(p, polygon[i-1], polygon[i])
		}
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}

func findAngleWithDirection(pn, p1, p2 point) float32 {
	rad := findAngle(
		normVector(vector{pn, p1}),
		normVector(vector{pn, p2}),
	)
	dir := vectorDirection(pn, vector{p1, p2})
	return float32(rad * float64(dir))
}

func isPointSame(p1, p2 point) bool {
	return p1[0].Cmp(p2[0]) == 0 && p1[1].Cmp(p2[1]) == 0
}
