package geosegment

import "testing"

func TestLength(t *testing.T) {
	s := Segment{
		Start: Point{Lat: -6.200000, Lng: 106.816666}, // Jakarta
		End:   Point{Lat: -6.914744, Lng: 107.609810}, // Bandung
	}

	got := s.Length()
	if got < 100 || got > 200 {
		t.Errorf("expected length between 100-200 km, got %.2f km", got)
	}
}
