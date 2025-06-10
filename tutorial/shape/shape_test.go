package shape

import "testing"

func TestPerimeter(t *testing.T) {
	assertCollectMessage := func(t testing.TB, calculated, expected float64) {
		t.Helper()
		if calculated != expected {
			t.Errorf("got %.2f want %.2f", calculated, expected)
		}
	}

	rectangle := Rectangle{4.0, 5.0};
	calculated := Perimeter(rectangle);
	expected := 18.0;

	assertCollectMessage(t, calculated, expected)
}


func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "長方形の面積計算", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "円形の面積計算", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "三角形の面積計算", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			calculated := tt.shape.Area();
			if(calculated != tt.hasArea) {
				t.Errorf("got %g want %g", calculated, tt.hasArea)
			}
		})
	}
}