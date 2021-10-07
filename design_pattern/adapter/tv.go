package main

type tv interface {
	switchChannel(channel int)
	switchVol(vol int)
	tvOn(power bool)
}
