package Griseo

import "math/rand"

type Griseo struct {
	Hp      int
	At      int
	Def     int
	Speed   int
	Shield  int
	State   int
	BackDef int
}

func (griseo *Griseo) InitObj() {
	griseo.Hp = 100
	griseo.At = 16
	griseo.Def = 11
	griseo.Speed = 18
	griseo.Shield = 0
	griseo.BackDef = 10
}

func (griseo *Griseo) Active() (blood int) {
	if griseo.Shield > 0 {
		blood = griseo.Def
	}
	griseo.Shield = 15
	return
}

func (griseo *Griseo) BreakShield() (blood int) {
	blood = griseo.Def * (200 + rand.Intn(200)) / 100
	return
}

func (griseo *Griseo) Passive() {
	if rand.Intn(100) < 40 {
		griseo.BackDef -= 2
		if griseo.BackDef < 0 {
			return
		}
		griseo.Def += 2
	}
}
