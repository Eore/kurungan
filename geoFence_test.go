package geofence

import (
	"testing"
)

type geoFenceTest struct {
	Point          point
	Polygon        []point
	ExpectedResult State
}

var geoFenceTestList = []geoFenceTest{
	{
		Point: toPoint(1, 1),
		Polygon: []point{
			toPoint(2, 0), toPoint(3, 0), toPoint(3, 2), toPoint(2, 2), toPoint(2, 0),
		},
		ExpectedResult: Outside,
	},
	{
		Point: toPoint(2, 2),
		Polygon: []point{
			toPoint(1, 0), toPoint(3, 0), toPoint(3, 3), toPoint(1, 3), toPoint(1, 0),
		},
		ExpectedResult: Inside,
	},
	{
		Point: toPoint(2, 2),
		Polygon: []point{
			toPoint(30, 10), toPoint(40, 40), toPoint(20, 40), toPoint(10, 20), toPoint(30, 10),
		},
		ExpectedResult: Outside,
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
		ExpectedResult: Outside,
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
		ExpectedResult: Outside,
	},
	{
		Point: toPoint(106.79878234863281, -6.359657747688895),
		Polygon: []point{
			toPoint(106.89147949218749, -6.196533712254142),
			toPoint(106.74110412597656, -6.171275658817273),
			toPoint(106.65870666503906, -6.18834204335982),
			toPoint(106.60652160644531, -6.2920937437552285),
			toPoint(106.58660888671875, -6.340549632917069),
			toPoint(106.56532287597656, -6.398554213119635),
			toPoint(106.59622192382812, -6.453823039876365),
			toPoint(106.644287109375, -6.481796323398132),
			toPoint(106.75003051757812, -6.5015813679646),
			toPoint(106.87431335449219, -6.506356951818031),
			toPoint(106.93679809570312, -6.45859907588465),
			toPoint(106.93817138671874, -6.410154340898689),
			toPoint(106.87156677246094, -6.407424922788262),
			toPoint(106.80152893066406, -6.43403612558649),
			toPoint(106.73149108886719, -6.4272128728934375),
			toPoint(106.67106628417969, -6.378765152234489),
			toPoint(106.67106628417969, -6.33236022397594),
			toPoint(106.67861938476562, -6.276395794877724),
			toPoint(106.70814514160156, -6.2374901460653245),
			toPoint(106.78092956542969, -6.275030733472354),
			toPoint(106.79397583007812, -6.331677767365065),
			toPoint(106.80976867675781, -6.364434665490979),
			toPoint(106.76925659179688, -6.378765152234489),
			toPoint(106.7596435546875, -6.400601313644121),
			toPoint(106.79191589355469, -6.413566092994113),
			toPoint(106.820068359375, -6.407424922788262),
			toPoint(106.86676025390625, -6.392412862343618),
			toPoint(106.88667297363281, -6.382859503462269),
			toPoint(106.92924499511719, -6.346691604277319),
			toPoint(106.93473815917969, -6.275030733472354),
			toPoint(106.91757202148438, -6.204042630374284),
			toPoint(106.89147949218749, -6.196533712254142),
		},
		ExpectedResult: Outside,
	},
}

func TestGeoGengeCheck(t *testing.T) {
	for i, testCase := range geoFenceTestList {
		res, err := GeoFenceCheck(testCase.Point, testCase.Polygon)
		if err != nil {
			t.Errorf("For test case %d error : %s", i, err.Error())
		} else if res != testCase.ExpectedResult {
			t.Errorf("For test case %d, expected result is %v, got %v", i, testCase.ExpectedResult, res)
		}
	}
}
