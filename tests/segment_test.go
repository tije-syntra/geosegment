package tests

import (
	"testing"

	"github.com/tije-syntra/geosegment"
	. "github.com/tije-syntra/geosegment"
)

func TestSegmentLength(t *testing.T) {
	s := geosegment.Segment{
		Start: geosegment.Point{Lat: -6.200000, Lng: 106.816666}, // Jakarta
		End:   Point{Lat: -7.8013953, Lng: 110.3641204},          // Yogyakarta
	}

	got := s.Length()
	if got < 100 || got > 500 {
		t.Errorf("expected 100-200 km, got %.2f km", got)
	}
}

func TestSegmentMidpoint(t *testing.T) {
	s := geosegment.Segment{
		Start: geosegment.Point{Lat: 0, Lng: 0},
		End:   geosegment.Point{Lat: 10, Lng: 10},
	}
	mid := s.Midpoint()

	if mid.Lat != 5 || mid.Lng != 5 {
		t.Errorf("expected midpoint (5,5), got (%.2f, %.2f)", mid.Lat, mid.Lng)
	}
}
