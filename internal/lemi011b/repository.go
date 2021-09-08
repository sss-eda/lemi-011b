package lemi011b

// DeviceRepository TODO
type DeviceRepository interface {
	Load(DeviceID) (*Device, error)
	Save(*Device) (DeviceID, error)
}
