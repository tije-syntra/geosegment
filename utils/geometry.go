package utils

import (
	"math"

	"github.com/paulmach/orb"
)

const EarthRadius = 6371000.0

// DegToMeter converts an angle from degrees to meters.
func DegToMeter(deg float64) float64 {
	return deg * 111_000 // 1 derajat ≈ 111 km
}

// DegToRad converts an angle from degrees to radians.
//
// Parameters:
//   - d: angle in degrees
//
// Returns:
//   - float64: angle in radians
func DegToRad(d float64) float64 {
	return d * math.Pi / 180
}

// RadToDeg converts an angle from radians to degrees.
//
// Parameters:
//   - r: angle in radians
//
// Returns:
//   - float64: angle in degrees
func RadToDeg(r float64) float64 {
	return r * 180 / math.Pi
}

// ClosestPointOnLine finds the closest point on a LineString to a given point.
//
// Parameters:
//   - line: an orb.LineString representing a polyline or route (sequence of points)
//   - point: the orb.Point to which the closest point on the line is sought
//
// Returns:
//   - orb.Point: the point on the line that is closest to the given point
//   - int: the index of the segment's starting point in the LineString where the closest point lies
//
// Notes:
//   - The function iterates through each segment of the line and projects the point
//     onto the segment using ClosestPointOnSegment.
//   - Distance is measured using the Haversine formula to account for Earth's curvature.
//   - Useful for finding the nearest location on a route or path to a GPS coordinate.
func ClosestPointOnLine(line orb.LineString, point orb.Point) (orb.Point, int) {
	minDist := math.MaxFloat64
	closestPoint := line[0]
	closestIndex := 0

	for i := 0; i < len(line)-1; i++ {
		a := line[i]
		b := line[i+1]

		proj := ClosestPointOnSegment(point, a, b)
		dist := Haversine(point, proj)

		if dist < minDist {
			minDist = dist
			closestPoint = proj
			closestIndex = i
		}
	}

	return closestPoint, closestIndex
}

// ClosestPointOnSegment returns the point on the line segment defined by points
// a and b that is closest to the point p.
//
// If the projection of p onto the line defined by a-b falls outside the segment,
// the closest endpoint (a or b) is returned.
//
// Parameters:
//   - p: the point to find the closest point to
//   - a: the starting point of the segment
//   - b: the ending point of the segment
//
// Returns:
//   - orb.Point: the closest point on the segment to p

func ClosestPointOnSegment(p, a, b orb.Point) orb.Point {
	ax, ay := a.X(), a.Y()
	bx, by := b.X(), b.Y()
	px, py := p.X(), p.Y()

	dx, dy := bx-ax, by-ay
	d := dx*dx + dy*dy

	if d == 0 {
		return a // a == b
	}

	t := ((px-ax)*dx + (py-ay)*dy) / d
	if t < 0 {
		return a
	} else if t > 1 {
		return b
	}

	return orb.Point{ax + t*dx, ay + t*dy}
}

// Haversine calculates the great-circle distance between two points on the Earth
// specified by latitude and longitude using the Haversine formula.
//
// Parameters:
//   - a: the first point (longitude, latitude) in degrees
//   - b: the second point (longitude, latitude) in degrees
//
// Returns:
//   - float64: the distance between the two points in the same unit as EarthRadius
//
// Notes:
//   - This function accounts for the Earth's curvature and is suitable for GPS coordinates.
//   - Ensure that the coordinates are in degrees; they will be converted to radians internally.
func Haversine(a, b orb.Point) float64 {
	lat1 := DegToRad(a[1])
	lon1 := DegToRad(a[0])
	lat2 := DegToRad(b[1])
	lon2 := DegToRad(b[0])

	dlat := lat2 - lat1
	dlon := lon2 - lon1

	h := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)

	return 2 * EarthRadius * math.Asin(math.Sqrt(h))
}

func LengthMeters(ls orb.LineString) float64 {
	total := 0.0
	for i := 1; i < len(ls); i++ {
		total += Haversine(ls[i-1], ls[i])
	}

	return total
}
