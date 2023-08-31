package model

type UpdateChannels struct {
	Location chan *Geolocation
}

func InitalizeUpdators() *UpdateChannels {
	location := make(chan *Geolocation)
	return &UpdateChannels{Location: location}
}

func (u *UpdateChannels) UpdateLocation(new_location *Geolocation) {
	u.Location <- new_location
}
