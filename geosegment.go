package geosegment

// Point merepresentasikan koordinat geografis (latitude, longitude).
type Point struct {
	Lat float64
	Lng float64
}

// Segment merepresentasikan garis dari dua titik.
type Segment struct {
	Start Point
	End   Point
}

// Line merepresentasikan garis dari dua titik.
type Line struct {
	Segments []Segment
}
