package bridges

import "fmt"

type RemoteControlInterface interface {
	TogglePower()
}

type VolumeControlInterface interface {
	VolumeUp()
	VolumeDown()
}

type RemoteControl struct {
	device Device
}

func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{
		device: device,
	}
}

func (r *RemoteControl) TogglePower() {
	r.device.setPower(!r.device.getPower())
	fmt.Println("Power toggled. Current power state:", r.device.getPower())
}

type RemoteControlWithVolume struct {
	RemoteControl
}

func NewRemoteControlWithVolume(device Device) *RemoteControlWithVolume {
	return &RemoteControlWithVolume{
		RemoteControl: RemoteControl{device: device},
	}
}

func (r *RemoteControlWithVolume) VolumeUp() {
	oldVolume := r.device.getVolume()
	newVolume := oldVolume + 1
	r.device.setVolume(newVolume)
	fmt.Println("The device", r.device.getName(), "volume increased from", oldVolume, "to", newVolume)
}

func (r *RemoteControlWithVolume) VolumeDown() {
	oldVolume := r.device.getVolume()
	newVolume := oldVolume - 1
	r.device.setVolume(newVolume)
	fmt.Println("The device", r.device.getName(), "volume decreased from", oldVolume, "to", newVolume)
}
