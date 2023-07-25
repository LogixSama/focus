package main

import (
	"errors"
	"fmt"
)

type SensorSpecs struct {
	sensorType string
	sensorName string
}

func getSensorTypes(machinemap map[string][]SensorSpecs) []string {
	sensorTypes := []string{}

	for _, values := range machinemap {
		for _, value := range values {
			if !isAvailable(value.sensorType, sensorTypes) {
				sensorTypes = append(sensorTypes, value.sensorType)
			}
		}
	}
	return sensorTypes
}

func getAllSensors(machinemap map[string][]SensorSpecs) []SensorSpecs {

	SensorsArray := []SensorSpecs{}

	for _, values := range machinemap {
		for _, value := range values {
			SensorsArray = append(SensorsArray, value)
		}
	}
	return SensorsArray
}

func checkMasterSensorRange(sensorType string, val int) bool {
	// H 22 37 2500
	// V 25 50 5000

	sensorValues, err := getMinMaxSensorValues(sensorType)
	if err != nil {
		fmt.Println(err)
	}

	minValue := sensorValues[0]
	maxValue := sensorValues[1]

	if val >= minValue && val <= maxValue {
		return true
	} else {
		return false
	}
}

func getMinMaxSensorValues(sensorType string) ([]int, error) {
	if sensorType == "H" {
		values := []int{22, 37}
		return values, nil
	} else if sensorType == "V" {
		values := []int{25, 50}
		return values, nil
	} else {
		var err = errors.New("Sensor Type not found")
		return nil, err
	}
}

func getSensorsForMachines(machineId string) ([]SensorSpecs, error) {
	if machineId == "M1" {
		sensors := []SensorSpecs{{sensorType: "H", sensorName: "H1"}, {sensorType: "V", sensorName: "V1"}}
		return sensors, nil
	} else if machineId == "M2" {
		sensors := []SensorSpecs{{sensorType: "H", sensorName: "H2"}, {sensorType: "V", sensorName: "V2"}}
		return sensors, nil
	} else {
		return nil, errors.New("Machine ID not found")
	}
}

func constructMachineSensorMap() map[string][]SensorSpecs {
	machineSensorMap := make(map[string][]SensorSpecs)
	machineIds := getMachinesIds()

	for _, machineId := range machineIds {
		sensors, err := getSensorsForMachines(machineId)
		if err != nil {
			fmt.Println(err)
		}
		machineSensorMap[machineId] = sensors
	}
	return machineSensorMap
}

func getSensorValue(sensor string) (int, error) {

	switch sensor {
	case "H1":
		return random(20, 38), nil

	case "V1":
		return random(24, 51), nil

	case "H2":
		return random(20, 38), nil

	case "V2":
		return random(24, 51), nil

	default:

		return 0, errors.New("SensorId not found")
	}

}

func checkSensorRepairCost(sensorType string) (int, error) {

	if sensorType == "H" {
		return 2500, nil
	} else if sensorType == "V" {
		return 5000, nil
	} else {
		return 0, errors.New("sensor type not found")
	}
}
