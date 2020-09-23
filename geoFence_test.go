package geofence

import (
	"testing"
)

type geoFenceTest struct {
	Point          point
	Polygon        []point
	ExpectedResult bool
}

var geoFenceTestList = []geoFenceTest{
	{
		Point: toPoint(1, 1),
		Polygon: []point{
			toPoint(2, 0), toPoint(3, 0), toPoint(3, 2), toPoint(2, 2),
		},
		ExpectedResult: false,
	},
	{
		Point: toPoint(2, 2),
		Polygon: []point{
			toPoint(1, 0), toPoint(3, 0), toPoint(3, 3), toPoint(1, 3),
		},
		ExpectedResult: true,
	},
	{
		Point: toPoint(2, 2),
		Polygon: []point{
			toPoint(30, 10), toPoint(40, 40), toPoint(20, 40), toPoint(10, 20), toPoint(30, 10),
		},
		ExpectedResult: false,
	},
	{
		Point: toPoint(106.9683837890625, -6.38490666680392),
		Polygon: []point{
			toPoint(106.82144165039061, -6.685067300513248),
			toPoint(107.12081909179688, -6.685067300513248),
			toPoint(107.12081909179688, -6.504992503919256),
			toPoint(106.82144165039061, -6.504992503919256),
			toPoint(106.82144165039061, -6.685067300513248),
		},
		ExpectedResult: false,
	},
	{
		Point: toPoint(106.9683837890625, -6.38490666680392),
		Polygon: []point{
			toPoint(106.71157836914061, -6.114611328564151),
			toPoint(106.66763305664062, -6.2361249830392875),
			toPoint(106.72531127929688, -6.32348821770752),
			toPoint(106.78573608398438, -6.383541892152094),
			toPoint(106.91070556640625, -6.3548807854976666),
			toPoint(106.94366455078125, -6.272983134659207),
			toPoint(106.98348999023438, -6.189707330332176),
			toPoint(106.93130493164062, -6.1159768049268965),
			toPoint(106.90109252929688, -6.0968598188879355),
			toPoint(106.820068359375, -6.130996814611725),
			toPoint(106.76101684570312, -6.1091493882593895),
			toPoint(106.71157836914061, -6.114611328564151),
		},
		ExpectedResult: false,
	},
}

func TestGeoGengeCheck(t *testing.T) {
	for i, testCase := range geoFenceTestList {
		res, err := GeoFenceCheck(testCase.Point, testCase.Polygon)
		if err != nil {
			t.Errorf("For test case %d, %s", i, err.Error())
		} else if res != testCase.ExpectedResult {
			t.Errorf("For test case %d, expected result is %v, got %v", i, testCase.ExpectedResult, res)
		}
	}
}
