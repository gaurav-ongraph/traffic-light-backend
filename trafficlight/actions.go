package trafficlight

import (
	"time"
)

func (m *trafficLightStruct) enterRed() {
	m.ChangeColor("red")
	m.mom.EnterRed()
}
func (m *trafficLightStruct) enterGreen() {
	m.ChangeColor("green")
	m.mom.EnterGreen()
}
func (m *trafficLightStruct) enterYellow() {
	m.ChangeColor("yellow")
	m.mom.EnterYellow()
}
func (m *trafficLightStruct) enterFlashingRed() {
	m.ChangeColor("red")
	m.mom.EnterFlashingRed()
}

func (m *trafficLightStruct) exitFlashingRed() {
	m.mom.ExitFlashingRed()
}

func (m *trafficLightStruct) startWorkingTimer() {
	m.mom.StartWorkingTimer()
}

func (m *trafficLightStruct) stopWorkingTimer() {
	m.mom.StopWorkingTimer()
}

func (m *trafficLightStruct) startFlashingTimer() {
	m.mom.StartFlashingTimer()
}
func (m *trafficLightStruct) stopFlashingTimer() {
	m.mom.StopFlashingTimer()
}

func (m *trafficLightStruct) startFlashing() {}
func (m *trafficLightStruct) stopFlashing()  {}

func (m *trafficLightStruct) changeFlashingAnimation() {
	flashColor := ""
	if m.flashColor == "red" {
		flashColor = "default"
	} else {
		flashColor = "red"
	}

	m.ChangeColor(flashColor)
	m.mom.ChangeFlashingAnimation()
}

func (m *trafficLightStruct) getColor() string {
	return m.flashColor
}

func (m *trafficLightStruct) systemError() {
	m.mom.SystemError()
}

func (m *trafficLightStruct) systemRestart() {
	m.mom.SystemRestart()
}

func (m *trafficLightStruct) log(msg string) {}

func (m *trafficLightMomStruct) enterRed() {
	color := m.trafficLight.GetColor()
	res := CreateResponse("working", color, "Destroy traffic light", false)
	SendResponse(res)
}

func (m *trafficLightMomStruct) enterGreen() {
	color := m.trafficLight.GetColor()
	res := CreateResponse("working", color, "Destroy traffic light", false)
	SendResponse(res)
}

func (m *trafficLightMomStruct) enterYellow() {
	color := m.trafficLight.GetColor()
	res := CreateResponse("working", color, "Destroy traffic light", false)
	SendResponse(res)
}

func (m *trafficLightMomStruct) enterFlashingRed() {
	color := m.trafficLight.GetColor()
	res := CreateResponse("error", color, "Destroy traffic light", false)
	SendResponse(res)
}
func (m *trafficLightMomStruct) exitFlashingRed() {
}

func (m *trafficLightMomStruct) startWorkingTimer() {
	Stopper = SetInterval(MOM.Tick, 2*time.Second)
}

func (m *trafficLightMomStruct) stopWorkingTimer() {
	Stopper <- true
}

func (m *trafficLightMomStruct) startFlashingTimer() {
	Stopper = SetInterval(MOM.Tick, 1*time.Second)
}

func (m *trafficLightMomStruct) stopFlashingTimer() {
	Stopper <- true
}

func (m *trafficLightMomStruct) startFlashing() {}
func (m *trafficLightMomStruct) stopFlashing()  {}

func (m *trafficLightMomStruct) changeFlashingAnimation() {
	color := m.trafficLight.GetColor()
	res := CreateResponse("error", color, "Destroy traffic light", false)
	SendResponse(res)
}

func (m *trafficLightMomStruct) systemError() {
	// res:= CreateResponse("error", "", "Create traffic light", false)
	// SendResponse(res)
}

func (m *trafficLightMomStruct) systemRestart() {
	// res:= CreateResponse("error", "", "Create traffic light", false)
	// SendResponse(res)
}

func (m *trafficLightMomStruct) log(msg string) {}
