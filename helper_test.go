package geofence

import "testing"

type pointInTriagleTest struct {
	Point          point
	Triangle       triangle
	ExpectedResult bool
}

var pointInTriagleTestList = []pointInTriagleTest{
	{
		Point: toPoint(82.85888671875, 32.7872745269555),
		Triangle: [3]vector{
			{
				toPoint(77.67333984375, 27.31321389856826),
				toPoint(74.92675781249999, 25.898761936567023),
			},
			{
				toPoint(74.92675781249999, 25.898761936567023),
				toPoint(80.88134765625, 24.16680208530324),
			},
			{
				toPoint(80.88134765625, 24.16680208530324),
				toPoint(77.67333984375, 27.31321389856826),
			},
		},
		ExpectedResult: false,
	},
	{
		Point: toPoint(77.71728515624999, 28.188243641850313),
		Triangle: [3]vector{
			{
				toPoint(77.080078125, 31.541089879585808),
				toPoint(69.697265625, 26.509904531413927),
			},
			{
				toPoint(69.697265625, 26.509904531413927),
				toPoint(82.529296875, 24.10664717920179),
			},
			{
				toPoint(82.529296875, 24.10664717920179),
				toPoint(77.080078125, 31.541089879585808),
			},
		},
		ExpectedResult: true,
	},
	{
		Point: toPoint(7.7728271484375, 35.56016058260795),
		Triangle: [3]vector{
			{
				toPoint(7.7398681640625, 35.572448615622804),
				toPoint(7.207031249999999,
					35.21420969483077),
			},
			{
				toPoint(7.207031249999999, 35.21420969483077),
				toPoint(8.28369140625, 35.23664622093195),
			},
			{
				toPoint(8.28369140625, 35.23664622093195),
				toPoint(7.7398681640625, 35.572448615622804),
			},
		},
		ExpectedResult: false,
	},
}

func TestPointInTriangle(t *testing.T) {
	for i, testCase := range pointInTriagleTestList {
		if inTri := pointInTriangle(testCase.Point, testCase.Triangle); inTri != testCase.ExpectedResult {
			t.Errorf("For test case %d => Expect %t, got %t", i, testCase.ExpectedResult, inTri)
		}
	}
}

// import "testing"

// type vectorDirectionTest struct {
// 	Point             point
// 	Vector            vector
// 	ExpectedDirection direction
// }

// var vectorDirectionTestList = []vectorDirectionTest{
// 	{
// 		Point: toPoint(106.9683837890625, -6.38490666680392),
// 		Vector: vector{
// 			toPoint(106.94366455078125, -6.272983134659207),
// 			toPoint(106.98348999023438, -6.189707330332176),
// 		},
// 		ExpectedDirection: CCW,
// 	},
// 	{
// 		Point:             toPoint(6, -1),
// 		Vector:            vector{toPoint(3, -1), toPoint(6, 2)},
// 		ExpectedDirection: CW,
// 	},
// 	{
// 		Point:             toPoint(1, 1),
// 		Vector:            vector{toPoint(2, 0), toPoint(3, 0)},
// 		ExpectedDirection: CCW,
// 	},
// 	{
// 		Point:             toPoint(1, 1),
// 		Vector:            vector{toPoint(3, 0), toPoint(3, 2)},
// 		ExpectedDirection: CCW,
// 	},
// 	{
// 		Point:             toPoint(1, 1),
// 		Vector:            vector{toPoint(3, 2), toPoint(2, 2)},
// 		ExpectedDirection: CCW,
// 	},
// 	{
// 		Point:             toPoint(1, 1),
// 		Vector:            vector{toPoint(2, 2), toPoint(2, 0)},
// 		ExpectedDirection: CW,
// 	},
// }

// func TestVectorDirection(t *testing.T) {
// 	for _, testCase := range vectorDirectionTestList {
// 		if dir := vectorDirection(testCase.Point, testCase.Vector); dir != testCase.ExpectedDirection {
// 			t.Errorf("%+v;%+v => Expect %f, got %f", testCase.Point, testCase.Vector, testCase.ExpectedDirection, dir)
// 		}
// 	}
// }

// func TestFindAngle(t *testing.T) {
// 	// t.Error(findAngle(vector{{0, 0}, {4, 3}}, vector{{0, 0}, {3, 5}}))
// 	// t.Error(Ge)
// }
