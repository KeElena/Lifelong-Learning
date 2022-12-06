package Pardofelis

import "math/rand"

type Pardofelis struct {
	Hp     int
	At     int
	Def    int
	Speed  int
	MaxHp  int
}

func (pardofelis *Pardofelis) InitObj() {
	pardofelis.Hp = 100
	pardofelis.At = 17
	pardofelis.Def = 10
	pardofelis.Speed = 24
	pardofelis.MaxHp =100
}

func (pardofelis *Pardofelis) Active() (blood int) {
	blood=30
	return
}

func (pardofelis *Pardofelis) Passive() {
	return
}

func (Pardofelis *Pardofelis)PassiveAt()int{
	if rand.Intn(100)<30{
		return 30
	}
	return 0
}