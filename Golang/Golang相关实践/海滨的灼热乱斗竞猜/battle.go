package main

import (
	"bh3_activity_2022_7_18/Aponia"
	"bh3_activity_2022_7_18/Griseo"
	"bh3_activity_2022_7_18/Pardofelis"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	round int
	Ch1   chan int
	Ch2   chan int
	wg    sync.WaitGroup
)

type action interface {
	InitObj()
	Active() (blood int)
	Passive()
}

func AponiaAndGriseo() {
	defer wg.Done()
	//initObj
	var Griseo Griseo.Griseo
	var Aponia Aponia.Aponia
	Griseo.InitObj()
	Aponia.InitObj()
	//init const
	round = 1
	//Go
	for {
		if round%4 == 0 {
			if Griseo.Shield > 0 {
				Griseo.Shield -= Aponia.Active() - Griseo.Def
				if Griseo.Shield < 0 {
					Aponia.Hp -= Griseo.BreakShield() - Aponia.Def
					Griseo.Hp += Griseo.Shield
					Griseo.Shield = 0
				}
			} else {
				Griseo.Hp -= Aponia.Active() - Griseo.Def
			}
			Griseo.State = Aponia.ActiveSeal()
		} else {
			if Griseo.Shield > 0 {
				Griseo.Shield -= Aponia.At - Griseo.Def
				if Griseo.Shield < 0 {
					Aponia.Hp -= Griseo.BreakShield() - Aponia.Def
					Griseo.Hp += Griseo.Shield
					Griseo.Shield = 0
				}
			} else {
				Griseo.Hp -= Aponia.At - Griseo.Def
			}
			Griseo.State = Aponia.PassiveSilent()
		}
		if Griseo.Hp <= 0 && Aponia.Hp <= 0 {
			Ch2 <- 1
			Ch1 <- 1
			break
		} else if Griseo.Hp <= 0 {
			Ch2 <- 1
			break
		} else if Aponia.Hp <= 0 {
			Ch1 <- 1
			break
		}

		switch Griseo.State {
		case 0:
			if round%3 == 0 {
				Aponia.Hp -= Griseo.Active()
			} else {
				Aponia.Hp -= Griseo.At - Aponia.Def
			}
			Griseo.Passive()
		case 1:
			Aponia.Hp -= Griseo.At - Aponia.Def
			Griseo.State = 0
		case 2:
			Griseo.State = 0
		}
		if Aponia.Hp <= 0 {
			Ch1 <- 1
			break
		}
		round++
	}
}

func GeiseoAndPardofelis() {
	//init
	defer wg.Done()
	var Griseo Griseo.Griseo
	var Pardofelis Pardofelis.Pardofelis
	Griseo.InitObj()
	Pardofelis.InitObj()
	//const
	round = 1
	//Go
	for {
		if round%3 == 0 {
			bloodAct := Pardofelis.Active() - Griseo.Def
			if Griseo.Shield > 0 {
				Griseo.Shield -= bloodAct
				if Griseo.Shield < 0 {
					Pardofelis.Hp -= Griseo.BreakShield() - Pardofelis.Def
					Griseo.Hp += Griseo.Shield
					if Pardofelis.MaxHp+Griseo.Shield >= 0 {
						Pardofelis.Hp -= Griseo.Shield
						Pardofelis.MaxHp += Griseo.Shield
					} else {
						Pardofelis.Hp += Pardofelis.MaxHp
						Pardofelis.MaxHp = 0
					}
					Griseo.Shield = 0
				}
			} else {
				Griseo.Hp -= bloodAct
				if Pardofelis.MaxHp-bloodAct >= 0 {
					Pardofelis.Hp += bloodAct
					Pardofelis.MaxHp -= bloodAct
				} else {
					Pardofelis.Hp += Pardofelis.MaxHp
					Pardofelis.MaxHp = 0
				}
			}
			//gt
			bloodPass := Pardofelis.PassiveAt() - Griseo.Def
			if bloodPass > 0 {
				if Griseo.Shield > 0 {
					Griseo.Shield -= bloodPass
					if Griseo.Shield < 0 {
						Pardofelis.Hp -= Griseo.BreakShield() - Pardofelis.Def
						Griseo.Hp += Griseo.Shield
						Griseo.Shield = 0
					}
				} else {
					Griseo.Hp -= bloodPass
				}
			}
			//round%3!=0
		} else {
			var blood int
			if Pardofelis.At-Griseo.Def < 0 {
				blood = 1 + Pardofelis.PassiveAt() - Griseo.Def
			} else {
				blood = Pardofelis.At + Pardofelis.PassiveAt() - Griseo.Def*2
			}
			if Griseo.Shield > 0 {
				Griseo.Shield -= blood
				if Griseo.Shield < 0 {
					Pardofelis.Hp -= Griseo.BreakShield() - Pardofelis.Def
					Griseo.Hp += Griseo.Shield
					Griseo.Shield = 0
				}
			} else {
				Griseo.Hp -= blood
			}
		}
		//check
		if Griseo.Hp <= 0 && Pardofelis.Hp <= 0 {
			Ch1 <- 1
			Ch2 <- 1
			break
		} else if Griseo.Hp <= 0 {
			Ch2 <- 1
			break
		} else if Pardofelis.Hp <= 0 {
			Ch1 <- 1
			break
		}

		if round%3 == 0 {
			if Griseo.Shield > 0 {
				Pardofelis.Hp -= Griseo.Def - Pardofelis.Def
			}
			Griseo.Shield = 15
		} else {
			Pardofelis.Hp -= Griseo.At - Pardofelis.Def
		}
		Griseo.Passive()

		if Pardofelis.Hp <= 0 {
			Ch1 <- 1
			break
		}

		round++
	}

}

func main() {
	//time seed
	rand.Seed(time.Now().UnixNano())
	//init chan
	Ch1 = make(chan int, 1000000)
	Ch2 = make(chan int, 1000000)
	n := 1000000
	//goroutine
	for i := 0; i < n; i++ {
		wg.Add(1)
		go GeiseoAndPardofelis()
	}
	wg.Wait()
	close(Ch1)
	close(Ch2)
	//reduce
	var GriseoSum int
	var PardofelisSum int

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			val, ok := <-Ch1
			if ok == false {
				break
			}
			GriseoSum += val
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			val, ok := <-Ch2
			if ok == false {
				break
			}
			PardofelisSum += val
		}
	}()

	wg.Wait()
	//report
	fmt.Println("\nPardofelis win:", PardofelisSum)
	fmt.Println("Griseo win:", GriseoSum)
	fmt.Printf("Pardofelis=%f\n", float64(PardofelisSum)/float64(n))
	fmt.Printf("Griseo=%f\n", float64(GriseoSum)/float64(n))
}
