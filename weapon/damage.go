package weapon

import (
	"math/rand"
)

type element int

const (
	fire element = iota
	water
	thunder
	earth
)

type nonElement int

const (
	poison nonElement = iota
	acid
	dark
	light
)

type elementalDamage struct {
	damager Damager
	element element
	dmg     int
}

func (d *elementalDamage) Damage() int {
	return d.dmg + d.damager.Damage()
}

func decorateElementalDamage(damager Damager, d *elementalDamage) Damager {
	d.damager = damager
	return d
}

type nonElementalDamage struct {
	damager    Damager
	nonElement nonElement
	dmg        int
}

func (d *nonElementalDamage) Damage() int {
	return d.dmg + d.damager.Damage()
}

func decorateNonElementalDamage(damager Damager, d *nonElementalDamage) Damager {
	d.damager = damager
	return d
}

type chaosDamage struct {
	Damager Damager
}

func (d *chaosDamage) Damage() int {
	dmg := rand.Intn(6)
	return dmg + d.Damager.Damage()
}

func decorateChaosDamage(damager Damager, d *chaosDamage) Damager {
	d.Damager = damager
	return d
}
