package geosegment

// Length menghitung panjang segmen (dalam kilometer).
func (s Segment) Length() float64 {
	return Haversine(s.Start, s.End)
}

// Midpoint menghitung titik tengah dari segmen.
func (s Segment) Midpoint() Point {
	return Point{
		Lat: (s.Start.Lat + s.End.Lat) / 2,
		Lng: (s.Start.Lng + s.End.Lng) / 2,
	}
}
