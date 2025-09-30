package tests

import (
	"math"
	"testing"

	"github.com/paulmach/orb"
	"github.com/tije-syntra/geosegment"
	"github.com/tije-syntra/geosegment/utils"
)

// helper untuk membandingkan float dengan toleransi
func floatEquals(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestDistance(t *testing.T) {
	a := orb.Point{106.875279, -6.191511}
	b := orb.Point{106.878737, -6.166889} // sekitar 80 m di lintang

	dist := utils.Haversine(a, b)
	if !floatEquals(dist, 2764, 5.0) {
		t.Errorf("Expected ~2764 m, got %v", dist)
	}
}

func TestLength(t *testing.T) {
	ls := orb.LineString{
		{106.875279, -6.191511},
		{106.878737, -6.166889},
	}

	lenMeter := utils.LengthMeters(ls)

	if !floatEquals(lenMeter, 2764, 1) { // jarak vertikal + horizontal ~ 111+111 km
		t.Errorf("Expected ~90 m, got %v", lenMeter)
	}
}

func TestSliceLine(t *testing.T) {
	ls := orb.LineString{
		{0, 0},
		{0, 1},
		{1, 1},
	}

	a := orb.Point{0, 0.2}
	b := orb.Point{1, 1}
	sliced := geosegment.SliceLine(a, b, ls)

	if len(sliced) < 2 {
		t.Errorf("Expected sliced line length >= 2, got %d", len(sliced))
	}
}

func TestNearestPointOnLine(t *testing.T) {
	ls := orb.LineString{
		{0, 0},
		{0, 1},
		{1, 1},
	}

	pt := orb.Point{0.1, 0.5}
	res := geosegment.NearestPointOnLine(pt, ls)

	if res.Type != "Feature" {
		t.Errorf("Expected Type 'Feature', got %s", res.Type)
	}
	if res.Properties["index"].(int) != 0 {
		t.Errorf("Expected index 0 for nearest segment, got %v", res.Properties["index"])
	}
}

func TestSnapToRoad(t *testing.T) {
	ls := orb.LineString{
		{0, 0},
		{0, 1},
		{1, 1},
	}

	pt := orb.Point{0.1, 0.5}
	prevPt := orb.Point{}
	currPt := orb.Point{0, 0}
	nextPt := orb.Point{0, 1}

	res := geosegment.SnapToRoad(prevPt, nextPt, currPt, pt, ls)

	if res.Distance <= 0 {
		t.Errorf("Expected distance > 0, got %v", res.Distance)
	}
	if res.Direction < 0 {
		t.Errorf("Expected direction >= 0, got %v", res.Direction)
	}
	if len(res.Geometry) != 2 {
		t.Errorf("Expected 2D point, got %v", res.Geometry)
	}
}
