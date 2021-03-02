package structs

import "testing"

func TestShapes(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 10, Width: 3}, hasArea: 30.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
		got := shape.Area()
		if got != hasArea {
			t.Errorf("%#v got %g but want %g", shape, got, hasArea)
		}
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.hasArea)
		})
	}
}
