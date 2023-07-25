// build unit

package main

import (
	"testing"
)

func TestGetSensorValue(t *testing.T) {

	t.Run("test the value for sensor of type H H1", func(t *testing.T) {
		got, _ := getSensorValue("H1")
		want := 25
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("test the value for sensor of type V V2", func(t *testing.T) {
		got, _ := getSensorValue("V2")
		want := 39
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("test the value for sensor of type V V1", func(t *testing.T) {
		got, _ := getSensorValue("H1")
		want := 31
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("test the value for sensor of type H H2", func(t *testing.T) {
		got, _ := getSensorValue("V2")
		want := 35
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("test the unknown sensor scenario", func(t *testing.T) {
		got, err := getSensorValue("B1")
		if got == 0 && err == nil {
			t.Errorf("The sensor value should not be present")
		}
	})
}

func TestCheckSensorRepairCost(t *testing.T) {

	t.Run("test the cost of repair for sensor of type H", func(t *testing.T) {
		got, _ := checkSensorRepairCost("H")
		if got != 2500 {
			t.Errorf("didn't expect %d", got)
		}
	})

	t.Run("test the cost of repair for sensor of type V", func(t *testing.T) {
		got, _ := checkSensorRepairCost("V")

		if got != 5000 {
			t.Errorf("didn't expect %d", got)
		}
	})

	t.Run("test if sensor type is not found", func(t *testing.T) {
		got, err := checkSensorRepairCost("B")

		if got == 0 && err == nil {
			t.Errorf("didn't expect %d", got)
		}
	})

}

func TestGetSensorsForMachines(t *testing.T) {

	t.Run("test the number of sensors to type M1", func(t *testing.T) {
		got, _ := getSensorsForMachines("M1")
		array := []SensorSpecs{{sensorType: "H", sensorName: "H1"}, {sensorType: "V", sensorName: "V1"}}

		for i, obj := range got {
			if obj.sensorName != array[i].sensorName {
				t.Errorf("didn't expect this sensor")
			}
		}
	})

	t.Run("test the number of sensors to type M2", func(t *testing.T) {
		got, _ := getSensorsForMachines("M2")
		array := []SensorSpecs{{sensorType: "H", sensorName: "H2"}, {sensorType: "V", sensorName: "V2"}}

		for i, obj := range got {
			if obj.sensorName != array[i].sensorName {
				t.Errorf("didn't expect this sensor")
			}
		}
	})

	t.Run("test the scenario if the machine isn't available", func(t *testing.T) {
		got, err := getSensorsForMachines("M2")

		if got == nil && err != nil {
			t.Errorf("erro should be thrown")
		}
	})

}

func TestGetMinMaxSensorValues(t *testing.T) {
	t.Run("test the min and max values for sensor of type H", func(t *testing.T) {
		got, _ := getMinMaxSensorValues("H")
		min := got[0]
		max := got[1]

		if min != 22 && max != 37 {
			t.Errorf("the range for the type of sensor does not match")
		}
	})

	t.Run("test the min and max values for sensor of type V", func(t *testing.T) {
		got, _ := getMinMaxSensorValues("V")
		min := got[0]
		max := got[1]

		if min != 25 && max != 50 {
			t.Errorf("the range for the type of sensor does not match")
		}
	})

	t.Run("test the error scenario if sensor type is not found", func(t *testing.T) {
		got, err := getMinMaxSensorValues("B")

		if got != nil && err == nil {
			t.Errorf("the type of sensor was not found")
		}
	})
}

func TestCheckMasterSensorRange(t *testing.T) {

	t.Run("test the cost of repair for sensor of type H", func(t *testing.T) {
		got := checkMasterSensorRange("H", 27)
		if got == false {
			t.Errorf("the type of sensor should be in range")
		}
	})

	t.Run("test the cost of repair for sensor of type H", func(t *testing.T) {
		got := checkMasterSensorRange("H", 100)
		if got == true {
			t.Errorf("the type of sensor should be out of range")
		}
	})

	t.Run("test the cost of repair for sensor of type V", func(t *testing.T) {
		got := checkMasterSensorRange("V", 32)
		if got == false {
			t.Errorf("the type of sensor should be in range")
		}
	})

	t.Run("test the cost of repair for sensor of type V", func(t *testing.T) {
		got := checkMasterSensorRange("V", 110)
		if got == true {
			t.Errorf("the type of sensor should be out of range")
		}
	})

}

func TestGetAllSensors(t *testing.T) {

	t.Run("test if all sensors are returned", func(t *testing.T) {

		m := make(map[string][]SensorSpecs)

		m["M1"] = []SensorSpecs{{sensorType: "H", sensorName: "H1"}, {sensorType: "V", sensorName: "V1"}}
		m["M2"] = []SensorSpecs{{sensorType: "H", sensorName: "H2"}, {sensorType: "V", sensorName: "V2"}}

		got := getAllSensors(m)

		arr := []string{"H1", "V1", "H2", "V2"}

		for i, sensorVal := range got {
			if sensorVal.sensorName != arr[i] {

				t.Errorf("the machine sensor was not found")
			}
		}
	})

}

func TestGetSensorTypes(t *testing.T) {
	t.Run("test if all sensors types are returned", func(t *testing.T) {

		m := make(map[string][]SensorSpecs)

		m["M1"] = []SensorSpecs{{sensorType: "H", sensorName: "H1"}, {sensorType: "V", sensorName: "V1"}}
		m["M2"] = []SensorSpecs{{sensorType: "H", sensorName: "H2"}, {sensorType: "V", sensorName: "V2"}}

		got := getSensorTypes(m)

		arr := []string{"H", "V"}

		for i, sensorType := range got {
			if sensorType != arr[i] {
				t.Errorf("the machine sensor was not found")
			}
		}
	})
}
