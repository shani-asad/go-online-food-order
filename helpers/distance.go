package helpers

import (
	"math"
)

type DistanceHelper struct {
}

func NewDistanceHelper() DistanceHelperInterface {
	return &DistanceHelper{}
}

func (h *DistanceHelper) GetHaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth's radius in kilometers

	// Convert degrees to radians
	lat1 = lat1 * math.Pi / 180
	lon1 = lon1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180
	lon2 = lon2 * math.Pi / 180

	// Haversine formula
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	circumference := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * circumference

	return distance
}
