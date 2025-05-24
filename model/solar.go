package model

import (
	"math/rand"
	"time"
)

type SolarData struct {
	TimeOfDay      time.Time // timestamp?
	SolarGenerated float64   // kWh
	BatteryStorage float64   // %
	GridPulled     float64   // kWh
	ShouldUseGrid  bool      //
}

func GenerateFakeData() SolarData {
	now := time.Now()
	hour := now.Hour()

	var solarGen float64
	if hour >= 6 && hour <= 18 {
		solarGen = rand.Float64() * 5 // up to 5 kWh during day
	} else {
		solarGen = 0 // night
	}

	battery := rand.Float64() * 100
	grid := rand.Float64() * 2 // small fallback use
	useGrid := solarGen < 1 && battery < 20

	return SolarData{
		TimeOfDay:      now,
		SolarGenerated: solarGen,
		BatteryStorage: battery,
		GridPulled:     grid,
		ShouldUseGrid:  useGrid,
	}
}
