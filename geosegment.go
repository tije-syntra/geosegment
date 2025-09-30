package geosegment

import (
	"math"

	"github.com/paulmach/orb"
	"github.com/tije-syntra/geosegment/utils"
)

type Point = orb.Point
type LineString = orb.LineString

type NearestPoint struct {
	Type       string         `json:"type"`
	Geometry   orb.Point      `json:"geometry"`
	Properties map[string]any `json:"properties"`
}

type SnapPoint struct {
	Geometry  orb.Point `json:"geometry"`
	Distance  float64   `json:"distance"`
	Direction float64   `json:"direction"`
}

// Distance returns the distance between two points in meters.
//
// Parameters:
//   - a: the first point (longitude, latitude) in degrees
//   - b: the second point (longitude, latitude) in degrees
//
// Returns:
//   - float64: the distance between the two points in meters
func Distance(a, b orb.Point) float64 {
	return utils.Haversine(a, b)
}

// Length returns the length of a LineString in meters.
//
// Parameters:
//   - ls: the orb.LineString representing the line (polyline or route)
//
// Returns:
//   - float64: the length of the line in meters
func Length(ls orb.LineString) float64 {
	return utils.LengthMeters(ls)
}

// SliceLine returns a subsection of a LineString between two points.
//
// Parameters:
//   - ls: the original orb.LineString (polyline) to be sliced
//   - a: the starting point of the slice
//   - b: the ending point of the slice
//
// Returns:
//   - orb.LineString: a new LineString representing the segment of ls
//     from the closest point to start to the closest point to end
//
// Notes:
//   - The function finds the closest points on the line to the provided
//     start and end points using ClosestPointOnLine.
//   - The order of points is maintained; if start comes after end in the
//     original LineString, the slice is reversed automatically.
//   - Useful for extracting a route segment from a larger path based on
//     GPS coordinates.
func SliceLine(a, b orb.Point, ls orb.LineString) orb.LineString {
	startPt, _ := utils.ClosestPointOnLine(ls, a)
	endPt, _ := utils.ClosestPointOnLine(ls, b)

	startFound := false
	sliced := orb.LineString{}

	for _, p := range ls {
		// Add point if startPt found
		if !startFound && utils.PointEquals(p, startPt) {
			startFound = true
		}
		if startFound {
			sliced = append(sliced, p)
		}
		// Stop when endPt found
		if utils.PointEquals(p, endPt) {
			break
		}
	}

	// Make sure start and end points are included
	if len(sliced) == 0 || !utils.PointEquals(sliced[len(sliced)-1], endPt) {
		sliced = append(sliced, endPt)
	}

	return sliced
}

// NearestPointOnLine finds the nearest point on a LineString to a given point
// and returns detailed information about the result.
//
// Parameters:
//   - ls: the orb.LineString representing the line (polyline or route)
//   - pt: the orb.Point to find the nearest point to
//
// Returns:
//   - NearestPoint: a struct containing the nearest point, distance,
//     segment index, and cumulative distance along the line.
//
// Notes:
//   - The function iterates through each segment of the LineString and projects
//     the point onto each segment using ClosestPointOnSegment.
//   - Distance calculations are done using Haversine, so the results are
//     accurate for geographic coordinates (lat/lon).
//   - The returned NearestPoint.Properties map includes:
//   - "dist": distance in meters from the point to the nearest point on the line
//   - "index": the index of the segment's starting point where the nearest point lies
//   - "location": cumulative distance along the line to the nearest point
//   - Useful for snapping GPS points to a route or path and calculating
//     distances along it.
func NearestPointOnLine(pt orb.Point, ls orb.LineString) NearestPoint {
	var nearest orb.Point
	minDist := math.MaxFloat64
	nearestIndex := -1
	totalLengthBefore := 0.0
	snappedLength := 0.0

	for i := 0; i < len(ls)-1; i++ {
		a := ls[i]
		b := ls[i+1]

		projected := utils.ClosestPointOnSegment(pt, a, b)
		dist := utils.Haversine(pt, projected)
		segLength := utils.Haversine(a, projected)

		if dist < minDist {
			minDist = dist
			nearest = projected
			nearestIndex = i
			snappedLength = totalLengthBefore + segLength
		}

		totalLengthBefore += utils.Haversine(a, b)
	}

	return NearestPoint{
		Type:     "Feature",
		Geometry: nearest,
		Properties: map[string]any{
			"dist":     minDist,
			"index":    nearestIndex,
			"location": snappedLength,
		},
	}
}

// SnapToRoad snaps a point to the nearest location on a road represented by a LineString,
// taking into account the context of previous, current, and next points along the path.
//
// Parameters:
//   - prevPt: the previous point along the path (can be empty for the start of the route)
//   - currPt: the current point along the path (can be empty for departure points)
//   - nextPt: the next point along the path (can be empty at the end of the route)
//   - pt: the point to be snapped to the road
//   - ls: the orb.LineString representing the road or route
//
// Returns:
//   - SnapPoint: a struct containing the snapped geometry point, the distance
//     from the original point to the road, and the direction (cumulative distance along the road)
//
// Notes:
//   - The function slices the road between appropriate points based on the route context
//     (start, departure, arrival, or end-to-end).
//   - Uses SliceLine to extract the relevant segment and NearestPointOnLine to find the
//     closest point.
//   - The returned SnapPoint provides:
//   - Geometry: the snapped orb.Point on the road
//   - Distance: distance in meters from the original point to the road
//   - Direction: cumulative distance along the road to the snapped point
//   - Useful for GPS point snapping in navigation, routing, or map-matching applications.
func SnapToRoad(prevPt, nextPt, currPt orb.Point, pt orb.Point, ls orb.LineString) SnapPoint {
	sliceLine := orb.LineString{}
	result := SnapPoint{}

	if len(prevPt) == 0 && len(currPt) != 0 && len(nextPt) != 0 { // start from first point
		startSlicePt := currPt
		endSlicePt := nextPt
		sliceLine = SliceLine(startSlicePt, endSlicePt, ls)

		snap := NearestPointOnLine(pt, sliceLine)
		result = SnapPoint{
			Geometry:  snap.Geometry,
			Distance:  snap.Properties["dist"].(float64),
			Direction: snap.Properties["location"].(float64),
		}
	} else if len(prevPt) != 0 && len(currPt) == 0 && len(nextPt) != 0 { // departure
		startSlicePt := prevPt
		endSlicePt := nextPt
		sliceLine = SliceLine(startSlicePt, endSlicePt, ls)

		snap := NearestPointOnLine(pt, sliceLine)
		result = SnapPoint{
			Geometry:  snap.Geometry,
			Distance:  snap.Properties["dist"].(float64),
			Direction: snap.Properties["location"].(float64),
		}
	} else if len(prevPt) != 0 && len(currPt) != 0 && len(nextPt) != 0 { // arrival
		startSlicePt := currPt
		endSlicePt := nextPt
		sliceLine = SliceLine(startSlicePt, endSlicePt, ls)

		snap := NearestPointOnLine(pt, sliceLine)
		result = SnapPoint{
			Geometry:  snap.Geometry,
			Distance:  snap.Properties["dist"].(float64),
			Direction: snap.Properties["location"].(float64),
		}
	} else if len(prevPt) != 0 && len(currPt) != 0 && len(nextPt) == 0 { // end to end
		startSlicePt := prevPt
		endSlicePt := currPt
		sliceLine = SliceLine(startSlicePt, endSlicePt, ls)

		snap := NearestPointOnLine(pt, sliceLine)
		result = SnapPoint{
			Geometry:  snap.Geometry,
			Distance:  snap.Properties["dist"].(float64),
			Direction: snap.Properties["location"].(float64),
		}
	}

	return result
}
