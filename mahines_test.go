// build unit

package main

import (
	"testing"
)

func TestGetMachinesIds(t *testing.T) {

	machines := []string{"M1", "M2"}
	t.Run("test if the machines IDs are present", func(t *testing.T) {
		machineArr := getMachinesIds()
		for i, machine := range machineArr {
			if machine != machines[i] {
				t.Errorf("The machine ID is not present")
			}
		}
	})
}

func TestFindMachineForSensor(t *testing.T) {

	m := make(map[string][]SensorSpecs)
	m["M1"] = []SensorSpecs{{sensorType: "H", sensorName: "H1"}, {sensorType: "V", sensorName: "V1"}}
	m["M2"] = []SensorSpecs{{sensorType: "H", sensorName: "H2"}, {sensorType: "V", sensorName: "V2"}}

	t.Run("test if the sensor is present for the machine", func(t *testing.T) {
		got, isPresent := findMachineForSensor(m, "H1")

		if got != "M1" && isPresent != true {
			t.Errorf("The machine does not have the sensor")
		}
	})

	t.Run("test if the sensor is not present for the machine", func(t *testing.T) {
		got, isPresent := findMachineForSensor(m, "T1")

		if got != "M1" && isPresent == true {
			t.Errorf("The machine does not have the sensor")
		}
	})

}
