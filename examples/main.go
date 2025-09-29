package main

import (
	"fmt"

	"github.com/tije-syntra/geosegment"
)

func main() {
	seg := geosegment.Segment{
		Start: geosegment.Point{Lat: -6.200000, Lng: 106.816666},   // Jakarta
		End:   geosegment.Point{Lat: -7.8013953, Lng: 110.3641204}, // Yogyakarta
	}

	fmt.Printf("Panjang segmen: %.2f km\n", seg.Length())
	fmt.Printf("Titik tengah: %+v\n", seg.Midpoint())
}
