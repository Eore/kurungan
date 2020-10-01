package geofence

import "fmt"

type State int

const (
	Outside State = 0
	Inside  State = 1
)

func GeoFenceCheck(p point, polygon []point) (State, error) {
	count := 0
	l := len(polygon)
	if !isPointSame(polygon[0], polygon[l-1]) {
		return Outside, fmt.Errorf("Polygon not closing")
	}
	if l < 3 {
		return Outside, fmt.Errorf("Is not polygon")
	}
	startPoint := polygon[0]
	for i := 0; i < l-3; i++ {
		if pointInTriangle(p, triangle{
			vector{startPoint, polygon[i+1]},
			vector{polygon[i+1], polygon[i+2]},
			vector{polygon[i+2], startPoint},
		}) {
			count++
		}
	}
	if count == 0 || count%2 == 0 {
		return Outside, nil
	}
	return Inside, nil
}

func isPointSame(p1, p2 point) bool {
	return p1[0].Cmp(p2[0]) == 0 && p1[1].Cmp(p2[1]) == 0
}
