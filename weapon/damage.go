package weapon

import (
	"math/rand"
)

type Element int

const (
	Fire Element = iota
	Water
	Thunder
	Earth
)

type NonElement int

const (
	Poison NonElement = iota
	Acid
	Dark
	Light
)

type Damager interface {
	Damage() int
}

type Difficulty int

const (
	Normal Difficulty = iota
	Easy
	Hard
)

var globalDifficulty Difficulty

type ElementalDamage struct {
	Damager    Damager
	Element    Element
	Dmg        int
	Difficulty Difficulty
}

func (d *ElementalDamage) Damage() int {
	dmg := d.Dmg

	if d.Difficulty == Easy {
		dmg = dmg + (dmg * 10 / 100)
	} else if d.Difficulty == Hard {
		dmg = dmg - (dmg * 10 / 100)
	}

	return dmg + d.Damager.Damage()
}

func DecorateElementalDamage(damager Damager, d *ElementalDamage) Damager {
	d.Damager = damager
	return d
}

type NonElementalDamage struct {
	Damager    Damager
	NonElement NonElement
	Dmg        int
}

func (d *NonElementalDamage) Damage() int {
	return d.Dmg + d.Damager.Damage()
}

func DecorateNonElementalDamage(damager Damager, d *NonElementalDamage) Damager {
	d.Damager = damager
	return d
}

type ChaosDamage struct {
	Damager Damager
}

func (d *ChaosDamage) Damage() int {
	dmg := rand.Intn(6)
	return dmg + d.Damager.Damage()
}

func DecorateChaosDamage(damager Damager, d *ChaosDamage) Damager {
	d.Damager = damager
	return d
}
