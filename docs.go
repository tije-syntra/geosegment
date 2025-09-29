// Package geosegment provides utilities for working with
// geographical points, segments, and distance calculations.
//
// Example usage:
//
//	seg := geosegment.Segment{
//	    Start: geosegment.Point{Lat: -6.2, Lng: 106.816666}, // Jakarta
//	    End:   geosegment.Point{Lat: -6.914744, Lng: 107.609810}, // Bandung
//	}
//
//	fmt.Printf("Length: %.2f km\n", seg.Length())
//
// This package supports:
// - Haversine distance calculation
// - Segment length & midpoint
// - Extensible unit conversion (km, miles, nautical miles in v2+)
package geosegment
