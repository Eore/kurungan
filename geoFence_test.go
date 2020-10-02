package geofence

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type geoFenceTest struct {
	Point          Point
	Polygon        []Point
	ExpectedResult State
}

var geoFenceTestList = []geoFenceTest{
	{
		Point: NewPoint(1, 1),
		Polygon: []Point{
			NewPoint(2, 0), NewPoint(3, 0), NewPoint(3, 2), NewPoint(2, 2), NewPoint(2, 0),
		},
		ExpectedResult: Outside,
	},
	{
		Point: NewPoint(2, 2),
		Polygon: []Point{
			NewPoint(1, 0), NewPoint(3, 0), NewPoint(3, 3), NewPoint(1, 3), NewPoint(1, 0),
		},
		ExpectedResult: Inside,
	},
	{
		Point: NewPoint(2, 2),
		Polygon: []Point{
			NewPoint(30, 10), NewPoint(40, 40), NewPoint(20, 40), NewPoint(10, 20), NewPoint(30, 10),
		},
		ExpectedResult: Outside,
	},
	{
		Point: NewPoint(106.9683837890625, -6.38490666680392),
		Polygon: []Point{
			NewPoint(106.82144165039061, -6.685067300513248),
			NewPoint(107.12081909179688, -6.685067300513248),
			NewPoint(107.12081909179688, -6.504992503919256),
			NewPoint(106.82144165039061, -6.504992503919256),
			NewPoint(106.82144165039061, -6.685067300513248),
		},
		ExpectedResult: Outside,
	},
	{
		Point: NewPoint(106.9683837890625, -6.38490666680392),
		Polygon: []Point{
			NewPoint(106.71157836914061, -6.114611328564151),
			NewPoint(106.66763305664062, -6.2361249830392875),
			NewPoint(106.72531127929688, -6.32348821770752),
			NewPoint(106.78573608398438, -6.383541892152094),
			NewPoint(106.91070556640625, -6.3548807854976666),
			NewPoint(106.94366455078125, -6.272983134659207),
			NewPoint(106.98348999023438, -6.189707330332176),
			NewPoint(106.93130493164062, -6.1159768049268965),
			NewPoint(106.90109252929688, -6.0968598188879355),
			NewPoint(106.820068359375, -6.130996814611725),
			NewPoint(106.76101684570312, -6.1091493882593895),
			NewPoint(106.71157836914061, -6.114611328564151),
		},
		ExpectedResult: Outside,
	},
	{
		Point: NewPoint(106.79878234863281, -6.359657747688895),
		Polygon: []Point{
			NewPoint(106.89147949218749, -6.196533712254142),
			NewPoint(106.74110412597656, -6.171275658817273),
			NewPoint(106.65870666503906, -6.18834204335982),
			NewPoint(106.60652160644531, -6.2920937437552285),
			NewPoint(106.58660888671875, -6.340549632917069),
			NewPoint(106.56532287597656, -6.398554213119635),
			NewPoint(106.59622192382812, -6.453823039876365),
			NewPoint(106.644287109375, -6.481796323398132),
			NewPoint(106.75003051757812, -6.5015813679646),
			NewPoint(106.87431335449219, -6.506356951818031),
			NewPoint(106.93679809570312, -6.45859907588465),
			NewPoint(106.93817138671874, -6.410154340898689),
			NewPoint(106.87156677246094, -6.407424922788262),
			NewPoint(106.80152893066406, -6.43403612558649),
			NewPoint(106.73149108886719, -6.4272128728934375),
			NewPoint(106.67106628417969, -6.378765152234489),
			NewPoint(106.67106628417969, -6.33236022397594),
			NewPoint(106.67861938476562, -6.276395794877724),
			NewPoint(106.70814514160156, -6.2374901460653245),
			NewPoint(106.78092956542969, -6.275030733472354),
			NewPoint(106.79397583007812, -6.331677767365065),
			NewPoint(106.80976867675781, -6.364434665490979),
			NewPoint(106.76925659179688, -6.378765152234489),
			NewPoint(106.7596435546875, -6.400601313644121),
			NewPoint(106.79191589355469, -6.413566092994113),
			NewPoint(106.820068359375, -6.407424922788262),
			NewPoint(106.86676025390625, -6.392412862343618),
			NewPoint(106.88667297363281, -6.382859503462269),
			NewPoint(106.92924499511719, -6.346691604277319),
			NewPoint(106.93473815917969, -6.275030733472354),
			NewPoint(106.91757202148438, -6.204042630374284),
			NewPoint(106.89147949218749, -6.196533712254142),
		},
		ExpectedResult: Outside,
	},
	{
		Point: NewPoint(107.02391624450684, -6.352449635920422),
		Polygon: []Point{
			NewPoint(106.89147949218749, -6.196533712254142),
			NewPoint(106.74110412597656, -6.171275658817273),
			NewPoint(106.65870666503906, -6.18834204335982),
			NewPoint(106.60652160644531, -6.2920937437552285),
			NewPoint(106.58660888671875, -6.340549632917069),
			NewPoint(106.56532287597656, -6.398554213119635),
			NewPoint(106.59622192382812, -6.453823039876365),
			NewPoint(106.644287109375, -6.481796323398132),
			NewPoint(106.75003051757812, -6.5015813679646),
			NewPoint(106.87431335449219, -6.506356951818031),
			NewPoint(106.93679809570312, -6.45859907588465),
			NewPoint(106.99516296386719, -6.427895202283447),
			NewPoint(107.00477600097656, -6.382859503462269),
			NewPoint(107.03086853027344, -6.344644288623892),
			NewPoint(106.95877075195312, -6.347374041020426),
			NewPoint(106.87156677246094, -6.407424922788262),
			NewPoint(106.80152893066406, -6.43403612558649),
			NewPoint(106.73149108886719, -6.4272128728934375),
			NewPoint(106.67106628417969, -6.378765152234489),
			NewPoint(106.67106628417969, -6.33236022397594),
			NewPoint(106.67861938476562, -6.276395794877724),
			NewPoint(106.70814514160156, -6.2374901460653245),
			NewPoint(106.78092956542969, -6.275030733472354),
			NewPoint(106.79397583007812, -6.331677767365065),
			NewPoint(106.80976867675781, -6.364434665490979),
			NewPoint(106.76925659179688, -6.378765152234489),
			NewPoint(106.7596435546875, -6.400601313644121),
			NewPoint(106.79191589355469, -6.413566092994113),
			NewPoint(106.820068359375, -6.407424922788262),
			NewPoint(106.86676025390625, -6.392412862343618),
			NewPoint(106.88667297363281, -6.382859503462269),
			NewPoint(106.92924499511719, -6.346691604277319),
			NewPoint(106.93473815917969, -6.275030733472354),
			NewPoint(106.91757202148438, -6.204042630374284),
			NewPoint(106.89147949218749, -6.196533712254142),
		},
		ExpectedResult: Inside,
	},
	{
		Point: NewPoint(106.88804626464844, -6.296871286298249),
		Polygon: []Point{
			NewPoint(106.89147949218749, -6.196533712254142),
			NewPoint(106.74110412597656, -6.171275658817273),
			NewPoint(106.65870666503906, -6.18834204335982),
			NewPoint(106.60652160644531, -6.2920937437552285),
			NewPoint(106.58660888671875, -6.340549632917069),
			NewPoint(106.56532287597656, -6.398554213119635),
			NewPoint(106.59622192382812, -6.453823039876365),
			NewPoint(106.644287109375, -6.481796323398132),
			NewPoint(106.75003051757812, -6.5015813679646),
			NewPoint(106.87431335449219, -6.506356951818031),
			NewPoint(106.93679809570312, -6.45859907588465),
			NewPoint(106.93817138671874, -6.410154340898689),
			NewPoint(106.87156677246094, -6.407424922788262),
			NewPoint(106.80152893066406, -6.43403612558649),
			NewPoint(106.73149108886719, -6.4272128728934375),
			NewPoint(106.67106628417969, -6.378765152234489),
			NewPoint(106.67106628417969, -6.33236022397594),
			NewPoint(106.67861938476562, -6.276395794877724),
			NewPoint(106.70814514160156, -6.2374901460653245),
			NewPoint(106.78092956542969, -6.275030733472354),
			NewPoint(106.79397583007812, -6.331677767365065),
			NewPoint(106.80976867675781, -6.364434665490979),
			NewPoint(106.76925659179688, -6.378765152234489),
			NewPoint(106.7596435546875, -6.400601313644121),
			NewPoint(106.79191589355469, -6.413566092994113),
			NewPoint(106.820068359375, -6.407424922788262),
			NewPoint(106.86676025390625, -6.392412862343618),
			NewPoint(106.88667297363281, -6.382859503462269),
			NewPoint(106.92924499511719, -6.346691604277319),
			NewPoint(106.93473815917969, -6.275030733472354),
			NewPoint(106.91757202148438, -6.204042630374284),
			NewPoint(106.89147949218749, -6.196533712254142),
		},
		ExpectedResult: Inside,
	},
}

func TestGeoFenceCheck(t *testing.T) {
	for i, testCase := range geoFenceTestList {
		res, err := GeoFenceCheck(testCase.Point, testCase.Polygon)
		if err != nil {
			t.Errorf("For test case %d error : %s", i, err.Error())
		} else if res != testCase.ExpectedResult {
			t.Errorf("For test case %d, expected result is %v, got %v", i, testCase.ExpectedResult, res)
		}
	}
}

type geoPolyFileTest struct {
	Point          Point
	PolyFilePath   string
	ExpectedResult State
}

var geoPolyFileTestList = []geoPolyFileTest{
	{
		Point:          NewPoint(106.91173553466797, -6.269015889070834),
		PolyFilePath:   "./polygon-test/jakarta.json",
		ExpectedResult: Outside,
	},
	{
		Point:          NewPoint(106.92212104797363, -6.254682363360378),
		PolyFilePath:   "./polygon-test/jakarta.json",
		ExpectedResult: Inside,
	},
	{
		Point:          NewPoint(106.91931612789631, -6.250678317605168),
		PolyFilePath:   "./polygon-test/jakarta.json",
		ExpectedResult: Outside,
	},
}

func TestGeoPolyFile(t *testing.T) {
	for i, testCase := range geoPolyFileTestList {
		data, errRead := readGeoJSON(testCase.PolyFilePath)
		if errRead != nil {
			t.Errorf("For test case %d error : %s", i, errRead.Error())
		}
		res, err := GeoFenceCheck(testCase.Point, data)
		if err != nil {
			t.Errorf("For test case %d error : %s", i, err.Error())
		} else if res != testCase.ExpectedResult {
			t.Errorf("For test case %d, expected result is %v, got %v", i, testCase.ExpectedResult, res)
		}
	}
}

func readGeoJSON(path string) ([]Point, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return []Point{}, err
	}
	var data [][2]float64
	json.Unmarshal(b, &data)
	var polygon []Point
	for _, value := range data {
		polygon = append(polygon, NewPoint(value[0], value[1]))
	}
	return polygon, nil
}
