package main

import (
	"fmt"
)

type HornSounder interface {
	SoundHorn()
}
type Vehicle struct {
	List [2]HornSounder
}
type Car struct {
	Sound string
}
type Bike struct {
	Sound string
}
type Rikshaw struct {
	Sound string
}

func (c Car) SoundHorn() {
	fmt.Println(c.Sound)
}

func (b Bike) SoundHorn() {
	fmt.Println(b.Sound)
}
func (r Rikshaw) SoundHorn() {
	fmt.Println(r.Sound)
}
func PressHorn(vehicle HornSounder) {
	vehicle.SoundHorn()
}
