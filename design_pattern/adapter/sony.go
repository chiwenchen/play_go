package main

type Sony struct {
	channel int
	vol     int
	power   bool
}

func (s *Sony) switchChannel(channel int) {
	s.channel = channel
}

func (s *Sony) switchVol(vol int) {
	s.vol = vol
}

func (s *Sony) tvOn(power bool) {
	s.power = power
}
