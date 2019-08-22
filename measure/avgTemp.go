package measure

import (
	"time"
	"sort"
)

type TempMeasurement struct {
	Timestamp	time.Time
	Measurement float64
	}

func AvgTemp(measurements []TempMeasurement) float64 {
	timeLimit := time.Minute * time.Duration(-2)
	z := 0
	mCount  := float64(z)
	tempSum := 0.00
	avgTemp := 0.00

  // Sort the slice so that correct number of measurements can be averaged
	sort.Slice(measurements, func(x, y int) bool {
		return measurements[x].Timestamp.After(measurements[y].Timestamp)
	})


	for z < len(measurements) {
	if(measurements[z].Timestamp.Before(measurements[0].Timestamp.Add(timeLimit))) {
				mCount = float64(z)
				avgTemp = (tempSum-measurements[z].Measurement)/mCount
				break
		}
		tempSum = tempSum + measurements[z].Measurement
	 z = z + 1
	}

	return avgTemp

}

