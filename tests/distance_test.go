package tests

import (
	"testing"

	"github.com/tije-syntra/geosegment"
)

func TestHaversine(t *testing.T) {
	jakarta := geosegment.Point{Lat: -6.200000, Lng: 106.816666}
	bandung := geosegment.Point{Lat: -6.914744, Lng: 107.609810}

	d := geosegment.Haversine(jakarta, bandung)
	if d < 100 || d > 200 {
		t.Errorf("expected 100-200 km, got %.2f km", d)
	}
}
