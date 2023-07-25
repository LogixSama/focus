package main

import (
	"fmt"
	"sync"
)

type sensorData struct {
	sensorType  string
	sensorName  string
	sensorValue int
}

func main() {

	machineSensorMap := constructMachineSensorMap()
	sensors := getAllSensors(machineSensorMap)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	// create a sensor and error channel

	sensorch := make(chan sensorData, len(sensors))
	errorch := make(chan sensorData, len(sensors))
	done := make(chan bool)

	// insert values in the channel
	go func() {

		// fetch all sensors

		for _, sensor := range sensors {

			sensorValue, err := getSensorValue(sensor.sensorName)

			if err != nil {

				fmt.Println(err)
			}

			sensorch <- sensorData{
				sensorName:  sensor.sensorName,
				sensorType:  sensor.sensorType,
				sensorValue: sensorValue,
			}
		}

		wg.Done()

	}()

	go func() {
		// read and check values inside the channel

		loop := cap(sensorch)

		for loop > 0 {
			sensorchData := <-sensorch
			if checkMasterSensorRange(sensorchData.sensorType, sensorchData.sensorValue) {
			} else {
				errorch <- sensorchData
			}
			loop--
		}

		done <- true

		wg.Done()

	}()

	go func() {

		<-done
		numerrors := len(errorch)

		switch {
		case numerrors >= 2:
			fmt.Println("Status: RED")

			totalCost := 0

			for len(errorch) > 0 {
				sensor := <-errorch
				machine, ok := findMachineForSensor(machineSensorMap, sensor.sensorName)
				if !ok {
					panic("value does not exist in map")
				}
				fmt.Printf("Sensor %s of machine %s \n", sensor.sensorName, machine)
				val, err := checkSensorRepairCost(sensor.sensorType)
				if err != nil {
					fmt.Println(err)
				}
				totalCost += val
			}

			fmt.Println("Cost of replacement: ", totalCost)
		case numerrors == 1:
			fmt.Println("Status: WARNING")

			msg := <-errorch
			machine, ok := findMachineForSensor(machineSensorMap, msg.sensorName)
			if !ok {
				panic("value does not exist in map")
			}
			fmt.Printf("Sensor %s of machine %s \n", msg.sensorName, machine)

			fmt.Println((""))

			cost, err := checkSensorRepairCost(msg.sensorType)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Cost of replacement: ", cost)

		default:
			fmt.Println("Status: GREEN")
		}

		wg.Done()

	}()

	wg.Wait()
}
