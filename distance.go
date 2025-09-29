package geosegment

import "math"

const earthRadiusKm = 6371.0

// Haversine menghitung jarak antar dua koordinat (km).
func Haversine(p1, p2 Point) float64 {
	dLat := (p2.Lat - p1.Lat) * math.Pi / 180.0
	dLng := (p2.Lng - p1.Lng) * math.Pi / 180.0

	lat1 := p1.Lat * math.Pi / 180.0
	lat2 := p2.Lat * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLng/2)*math.Sin(dLng/2)*math.Cos(lat1)*math.Cos(lat2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}
