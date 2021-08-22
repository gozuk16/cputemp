package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

func main() {
	h1, h2, h3, _ := host.PlatformInformation()
	fmt.Println(h1, h2, h3)
	t, _ := host.SensorsTemperatures()
	for _, v := range t {
		fmt.Printf("sensorKey: %s, temperature: %2.0f, sensorHigh: %2.0f, sensorCritical: %2.0f\n", v.SensorKey, v.Temperature, v.High, v.Critical)
	}
}
