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
	a := orb.Point{106.876024, -6.254066} // lon, lat
	b := orb.Point{106.875928, -6.254471}

	dist := geosegment.Distance(a, b)
	fmt.Printf("Distance between a and b: %.2f meters\n", dist)
	fmt.Println("---------------------------------------------")

	// ----------------------
	// Contoh 2: Length of LineString
	// ----------------------
	ls := orb.LineString{
		{106.876024, -6.254066},
		{106.875928, -6.254471},
		{106.875976, -6.254615},
		{106.876298, -6.254679},
	}

	fmt.Println("LineString:")
	for i, p := range ls {
		fmt.Printf("  Point %d: %v\n", i, p)
	}
	fmt.Println("---------------------------------------------")
	length := geosegment.Length(ls)
	fmt.Printf("Length of LineString: %.2f meters\n", length)

	// ----------------------
	// Contoh 3: SliceLine
	// ----------------------
	startPt := orb.Point{106.876024, -6.254066}
	endPt := orb.Point{106.875976, -6.254615}
	slice := geosegment.SliceLine(startPt, endPt, ls)
	fmt.Println("Sliced LineString:")
	for i, p := range slice {
		fmt.Printf("  Point %d: %v\n", i, p)
	}
	fmt.Println("---------------------------------------------")
	// ----------------------
	// Contoh 4: NearestPointOnLine
	// ----------------------
	pt := orb.Point{106.875799, -6.254226}
	nearest := geosegment.NearestPointOnLine(pt, ls)
	fmt.Printf("Nearest point on line to pt: %v\n", nearest.Geometry)
	fmt.Printf("Distance: %.2f meters, Segment index: %v, Location along line: %.2f meters\n",
		nearest.Properties["dist"].(float64),
		nearest.Properties["index"],
		nearest.Properties["location"].(float64),
	)
	fmt.Println("---------------------------------------------")

	// ----------------------
	// Contoh 5: SnapToRoad
	// ----------------------
	prevPt := orb.Point{} // kosong, start point
	currPt := orb.Point{106.875928, -6.254471}
	nextPt := orb.Point{106.876298, -6.254679}
	toSnap := orb.Point{106.875874, -6.254178}

	snap := geosegment.SnapToRoad(prevPt, nextPt, currPt, toSnap, ls)
	fmt.Println("SnapToRoad result:")
	fmt.Printf("  Geometry: %v\n  Distance to road: %.2f meters\n  Direction along road: %.2f meters\n",
		snap.Geometry, snap.Distance, snap.Direction)
	fmt.Println("---------------------------------------------")
}
