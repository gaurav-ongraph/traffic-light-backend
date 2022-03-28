```
package trafficlight

import (
	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)
```

#TrafficLightMom

    -interface-

    start @(|>>|)
    stop @(|<<|)
    tick    
    enterRed
    enterGreen
    enterYellow
    enterFlashingRed
    exitFlashingRed
    startWorkingTimer
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer
    startFlashing
    stopFlashing
    changeFlashingAnimation
    systemError
    systemRestart
    log [msg:string]
    -machine-

    $New => $TrafficLightApi
        |>>| 
            trafficLight = NewTrafficLight(#)
            trafficLight.Start()
            -> "Traffic Light\nStarted" $Saving ^
 
    $Saving 
        |>|
            data = trafficLight.Marshal() 
            trafficLight = nil 
            -> "Saved" $Persisted ^

    $Persisted 
        |tick| -> "Tick" $Working("tick") ^
        |systemError| -> "System Error" $Working("systemError") ^
        |systemRestart| -> "System Restart" $Working("systemRestart") ^
        |<<| -> "Stop" $End ^

    $Working[trafficLightEvent:string color:string] => $TrafficLightApi
        |>|
            trafficLight = LoadTrafficLight(# data)
            trafficLightEvent ?~
                /tick/ trafficLight.Tick() -> "Done" $Saving :>
                /systemError/ trafficLight.SystemError() -> "Done" $Saving :>
                /systemRestart/ trafficLight.SystemRestart() -> "Done" $Saving :: ^

    $TrafficLightApi
        |enterRed| enterRed() ^
        |enterGreen| enterGreen()  ^
        |enterYellow| enterYellow() ^
        |enterFlashingRed| enterFlashingRed() ^
        |exitFlashingRed| exitFlashingRed() ^
        |startWorkingTimer| startWorkingTimer() ^
        |stopWorkingTimer| stopWorkingTimer() ^
        |startFlashingTimer| startFlashingTimer() ^
        |stopFlashingTimer| stopFlashingTimer() ^
        |startFlashing| startFlashing() ^
        |stopFlashing| stopFlashing() ^
        |changeFlashingAnimation| changeFlashingAnimation() ^
        |log| [msg:string] log(msg) ^

    $End => $TrafficLightApi
        |>|
            trafficLight = LoadTrafficLight(# data) 
            trafficLight.Stop() ^

    -actions-

    enterRed
    enterGreen
    enterYellow
    enterFlashingRed
    exitFlashingRed
    startWorkingTimer
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer   
    startFlashing
    stopFlashing
    changeFlashingAnimation
    log [msg:string]

    -domain-

    var trafficLight:TrafficLight = null
    var data:`[]byte` = null

##