package main

type samsungAdapter struct {
	tv *SamSung
}

func (s *samsungAdapter) switchChannel(channel int) {
	s.tv.currentChannel = channel
}

func (s *samsungAdapter) switchVol(vol int) {
	s.tv.currentVol = vol
}

func (s *samsungAdapter) tvOn(power bool) {
	if power == true {
		s.tv.turnTvOn()
	} else {
		s.tv.turnTvOff()
	}
}
