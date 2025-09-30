package main

import (
	"fmt"

	"github.com/paulmach/orb"
	"github.com/tije-syntra/geosegment"
)

func main() {
	// ----------------------
	// Contoh 1: Distance
	// ----------------------
	a := orb.Point{106.876466, -6.253991} // lon, lat
	b := orb.Point{106.876458, -6.253983}

	dist := geosegment.Distance(a, b)
	fmt.Printf("Distance between a and b: %.2f meters\n", dist)

	// ----------------------
	// Contoh 2: Length of LineString
	// ----------------------
	ls := orb.LineString{
		{106.876466, -6.253991},
		{106.876684, -6.253689},
		{106.876487, -6.253432},
	}
	length := geosegment.Length(ls)
	fmt.Printf("Length of LineString: %.2f meters\n", length)

	// ----------------------
	// Contoh 3: SliceLine
	// ----------------------
	startPt := orb.Point{106.876466, -6.253991}
	endPt := orb.Point{106.876487, -6.253432}
	slice := geosegment.SliceLine(startPt, endPt, ls)
	fmt.Println("Sliced LineString:")
	for i, p := range slice {
		fmt.Printf("  Point %d: %v\n", i, p)
	}

	// ----------------------
	// Contoh 4: NearestPointOnLine
	// ----------------------
	pt := orb.Point{106.876500, -6.253800}
	nearest := geosegment.NearestPointOnLine(pt, ls)
	fmt.Printf("Nearest point on line to pt: %v\n", nearest.Geometry)
	fmt.Printf("Distance: %.2f meters, Segment index: %v, Location along line: %.2f meters\n",
		nearest.Properties["dist"].(float64),
		nearest.Properties["index"],
		nearest.Properties["location"].(float64),
	)

	// ----------------------
	// Contoh 5: SnapToRoad
	// ----------------------
	prevPt := orb.Point{} // kosong, start point
	currPt := orb.Point{106.876466, -6.253991}
	nextPt := orb.Point{106.876487, -6.253432}
	toSnap := orb.Point{106.876500, -6.253800}

	snap := geosegment.SnapToRoad(prevPt, nextPt, currPt, toSnap, ls)
	fmt.Println("SnapToRoad result:")
	fmt.Printf("  Geometry: %v\n  Distance to road: %.2f meters\n  Direction along road: %.2f meters\n",
		snap.Geometry, snap.Distance, snap.Direction)
}
