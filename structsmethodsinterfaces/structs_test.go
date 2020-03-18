package structsmethodsinterfaces_test

import "testing"
import "github.com/devaof/learn-go-with-tests/structsmethodsinterfaces"

type Shape interface {
	Area() float64
}

func TestPerimeters(t *testing.T) {
	rec := structsmethodsinterfaces.Rectangle{10.0, 10.0}
	got := structsmethodsinterfaces.Perimeters(rec)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{structsmethodsinterfaces.Rectangle{12, 6}, 72.0},
		{structsmethodsinterfaces.Circle{10}, 314.1592653589793},
		{structsmethodsinterfaces.Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}

	}
}
