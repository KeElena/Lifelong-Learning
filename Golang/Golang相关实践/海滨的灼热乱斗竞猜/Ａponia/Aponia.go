package Aponia

import "math/rand"

type Aponia struct {
	Hp     int
	At     int
	Def    int
	Speed  int
	Shield int
}

func (aponia *Aponia) InitObj() {
	aponia.Hp = 100
	aponia.At = 21
	aponia.Def = 10
	aponia.Speed = 30
}

func (aponia *Aponia) Active() (blood int) {
	blood = int(float32(aponia.At) * 1.7)
	return
}

func (aponia *Aponia) ActiveSeal() int {
	return 2
}

func (aponia *Aponia) Passive() {
	return
}

func (aponia *Aponia) PassiveSilent() int {
	if rand.Intn(100) < 30 {
		return 1
	}
	return 0
}