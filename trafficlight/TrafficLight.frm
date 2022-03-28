```
package trafficlight

import (
	"encoding/json"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)
```
#[derive(Managed,Marshal)]
#[mom="TrafficLightMom"]
#TrafficLight

    -interface-

    start @(|>>|)
    stop 
    tick
    systemError
    systemRestart
    changeColor[color:string]
    getColor:string

    -machine-

    $Begin
        |>>|
            startWorkingTimer()
            -> $Red ^

    $Red => $Working
        |>|
            enterRed() ^
        |tick|
            -> $Green ^

    $Green => $Working
        |>|
            enterGreen() ^
        |tick|
            -> $Yellow ^

    $Yellow => $Working
        |>|
            enterYellow() ^
        |tick|
            -> $Red ^

    $FlashingRed
        |>|
            enterFlashingRed()
            stopWorkingTimer()
            startFlashingTimer() ^
        |<|
            exitFlashingRed()
            stopFlashingTimer()
            startWorkingTimer() ^
        |tick|
            changeFlashingAnimation() ^
        |changeColor|[color:string]
            flashColor = color ^
        |systemRestart|
            -> $Red  ^
        |stop|
            -> $End ^
        |getColor|:string
            @^ = getColor() ^

    $End
        |>| stopWorkingTimer() ^

    $Working
        |stop|
            -> $End ^
        |systemError|
            -> $FlashingRed ^
        |changeColor|[color:string]
            flashColor = color ^
        |getColor|:string
            @^ = getColor() ^

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
    getColor:string
    log [msg:string]

    -domain-

    var flashColor:string = ""

##