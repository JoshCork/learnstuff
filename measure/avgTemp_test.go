package measure

import (
	"testing"
	"time"
)

type Parameters struct {
	Input			[]TempMeasurement	// testing input
	ExOutput 	float64						// expected output
}

var avgTempTests []Parameters

func makeSample(countMeasures int, sampleTemp1 float64, sampleTemp2 float64) []TempMeasurement{

	measurements := []TempMeasurement{}
	currentMeasure := TempMeasurement{}
	now := time.Now()

	i := 1
	for i <= countMeasures {
	if (i < countMeasures/2) {
		currentMeasure = TempMeasurement{Timestamp: now.Add(time.Second * time.Duration(-i)), Measurement: sampleTemp1}
	} else {
		currentMeasure = TempMeasurement{Timestamp: now.Add(time.Second * time.Duration(-i)), Measurement: sampleTemp2}
	}

	measurements = append(measurements, currentMeasure)
	i = i + 1
	}

	return measurements

}

func TestAvgTemp(t *testing.T) {
	avgTempTests = append(avgTempTests, Parameters{Input: makeSample(240, 100,200), ExOutput: 100})
	for _,testCase := range avgTempTests {
		actual := AvgTemp(testCase.Input)
		if actual != testCase.ExOutput {
			t.Errorf("Fail! For input: %v | Expected: %v | Received: %v", testCase.Input, testCase.ExOutput, actual)
		}
	}
}