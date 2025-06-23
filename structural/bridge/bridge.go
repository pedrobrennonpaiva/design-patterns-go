package bridge

import (
	"design-patterns-go/structural/bridge/bridges"
	"fmt"
)

func Run() {
	fmt.Println("")

	tv := bridges.NewTV()
	radio := bridges.NewRadio()

	radioRemoteControl := bridges.NewRemoteControl(radio)
	tvRemoteControl := bridges.NewRemoteControlWithVolume(tv)

	runClient(radioRemoteControl)
	runClient(tvRemoteControl)
}

func runClient(remote bridges.RemoteControlInterface) {
	remote.TogglePower()

	if volumeControl, ok := remote.(bridges.VolumeControlInterface); ok {
		fmt.Println("This remote supports volume control.")
		volumeControl.VolumeUp()
		volumeControl.VolumeUp()
		volumeControl.VolumeDown()
	} else {
		fmt.Println("This remote does not support volume control.")
	}
}
