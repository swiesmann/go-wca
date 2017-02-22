package main

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioEndpointVolume struct {
	ole.IUnknown
}

type IAudioEndpointVolumeVtbl struct {
	ole.IUnknownVtbl
	RegisterControlChangeNotify   uintptr
	UnregisterControlChangeNotify uintptr
	GetChannelCount               uintptr
	SetMasterVolumeLevel          uintptr
	SetMasterVolumeLevelScalar    uintptr
	GetMasterVolumeLevel          uintptr
	GetMasterVolumeLevelScalar    uintptr
	SetChannelVolumeLevel         uintptr
	SetChannelVolumeLevelScalar   uintptr
	GetChannelVolumeLevel         uintptr
	GetChannelVolumeLevelScalar   uintptr
	SetMute                       uintptr
	GetMute                       uintptr
	GetVolumeStepInfo             uintptr
	VolumeStepUp                  uintptr
	VolumeStepDown                uintptr
	QueryHardwareSupport          uintptr
	GetVolumeRange                uintptr
}

func (v *IAudioEndpointVolume) VTable() *IAudioEndpointVolumeVtbl {
	return (*IAudioEndpointVolumeVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioEndpointVolume) GetChannelCount(channelCount *uint32) (err error) {
	err = getChannelCount(v, channelCount)
	return
}
func (v *IAudioEndpointVolume) GetMasterVolumeLevelScalar(level *float32) (err error) {
	err = getMasterVolumeLevelScalar(v, level)
	return
}

func (v *IAudioEndpointVolume) SetMasterVolumeLevelScalar(level float32, eventContextGUID *ole.GUID) (err error) {
	err = setMasterVolumeLevelScalar(v, level, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) VolumeStepUp(eventContextGUID *ole.GUID) (err error) {
	err = volumeStepUp(v, eventContextGUID)
	return
}
