package helpers

import (
	"log"
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

func Permute(destinations [][2]float64) [][][2]float64 {
    var helper func([][2]float64, int)
    res := [][][2]float64{}

    helper = func(arr [][2]float64, n int) {
        if n == 1 {
            tmp := make([][2]float64, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++ {
                helper(arr, n-1)
                if n%2 == 1 {
                    arr[0], arr[n-1] = arr[n-1], arr[0]
                } else {
                    arr[i], arr[n-1] = arr[n-1], arr[i]
                }
            }
        }
    }

    helper(destinations, len(destinations))
    return res
}

// func CalculateShortestTime(userLat, userLon, startLat, startLon float64, destinations [][2]float64) int {
func CalculateShortestTime(startLat, startLon float64, destinations [][2]float64) float64 {
    h := &DistanceHelper{}
    permutations := Permute(destinations)
    minDistance := math.MaxFloat64

    for _, perm := range permutations {
        totalDistance := h.GetHaversineDistance(startLat, startLon, perm[0][0], perm[0][1])
        for i := 0; i < len(perm)-1; i++ {
            totalDistance += h.GetHaversineDistance(perm[i][0], perm[i][1], perm[i+1][0], perm[i+1][1])
        }
        // totalDistance += h.GetHaversineDistance(perm[len(perm)-1][0], perm[len(perm)-1][1], userLat, userLon)

		if totalDistance < minDistance {
			minDistance = totalDistance
		}
    }
    log.Println("minDistance", minDistance)
	speed := 40.0
	totalTimeHours := minDistance / speed
    log.Println("totalTimeHours", totalTimeHours)
    totalTimeMinutes := totalTimeHours * 60
    log.Println("totalTimeMinutes", totalTimeMinutes)

    return totalTimeMinutes
}