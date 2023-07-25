package main

func getMachinesIds() []string {
	return []string{"M1", "M2"}
}

func findMachineForSensor(m map[string][]SensorSpecs, sensorq string) (string, bool) {
	for machine, sensors := range m {
		for _, sensor := range sensors {
			if sensor.sensorName == sensorq {
				return machine, true
			}
		}
	}

	return "", false
}
