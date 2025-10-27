package main

import "fmt"

type Thermometer interface {
	GetTemperatureInCelsius() float64
}

type KelvinSensor struct{}

func (ks *KelvinSensor) GetTemperatureInKelvin() float64 {
	return 300.15

}

type CelsiusAdapter struct {
	kelvinSensor *KelvinSensor
}

func (ca *CelsiusAdapter) GetTemperatureInCelsius() float64 {
	kelvin := ca.kelvinSensor.GetTemperatureInKelvin()

	celsius := kelvin - 273.15

	return celsius
}

func DisplayWeather(t Thermometer) {
	tempC := t.GetTemperatureInCelsius()
	fmt.Printf("Погода: %.2f °C\n", tempC)
}

func main() {
	kelvinSensor := &KelvinSensor{}

	adapter := &CelsiusAdapter{kelvinSensor: kelvinSensor}

	DisplayWeather(adapter)
}
