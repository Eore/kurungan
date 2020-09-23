package geofence

import "testing"

type vectorDirectionTest struct {
	Point             point
	Vector            vector
	ExpectedDirection direction
}

var vectorDirectionTestList = []vectorDirectionTest{
	{
		Point: toPoint(106.9683837890625, -6.38490666680392),
		Vector: vector{
			toPoint(106.94366455078125, -6.272983134659207),
			toPoint(106.98348999023438, -6.189707330332176),
		},
		ExpectedDirection: CCW,
	},
	{
		Point:             toPoint(6, -1),
		Vector:            vector{toPoint(3, -1), toPoint(6, 2)},
		ExpectedDirection: CW,
	},
	{
		Point:             toPoint(1, 1),
		Vector:            vector{toPoint(2, 0), toPoint(3, 0)},
		ExpectedDirection: CCW,
	},
	{
		Point:             toPoint(1, 1),
		Vector:            vector{toPoint(3, 0), toPoint(3, 2)},
		ExpectedDirection: CCW,
	},
	{
		Point:             toPoint(1, 1),
		Vector:            vector{toPoint(3, 2), toPoint(2, 2)},
		ExpectedDirection: CCW,
	},
	{
		Point:             toPoint(1, 1),
		Vector:            vector{toPoint(2, 2), toPoint(2, 0)},
		ExpectedDirection: CW,
	},
}

func TestVectorDirection(t *testing.T) {
	for _, testCase := range vectorDirectionTestList {
		if dir := vectorDirection(testCase.Point, testCase.Vector); dir != testCase.ExpectedDirection {
			t.Errorf("%+v;%+v => Expect %f, got %f", testCase.Point, testCase.Vector, testCase.ExpectedDirection, dir)
		}
	}
}

func TestFindAngle(t *testing.T) {
	// t.Error(findAngle(vector{{0, 0}, {4, 3}}, vector{{0, 0}, {3, 5}}))
	// t.Error(Ge)
}
