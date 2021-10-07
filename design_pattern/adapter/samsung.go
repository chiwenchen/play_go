package main

type SamSung struct {
	currentChannel int
	currentVol     int
	isTvOn         bool
}

func (s *SamSung) changeCurrentChannel(channel int) {
	s.currentChannel = channel
}

func (s *SamSung) changeCurrentVol(vol int) {
	s.currentVol = vol
}

func (s *SamSung) turnTvOn() {
	s.isTvOn = true
}

func (s *SamSung) turnTvOff() {
	s.isTvOn = false
}
