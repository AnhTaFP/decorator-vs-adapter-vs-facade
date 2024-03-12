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

type nonElement int

const (
	poison nonElement = iota
	acid
	dark
	light
)

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

type darksoulsDamage struct {
	adapter *darkSoulsAdapter
	damager Damager
}

func (d *darksoulsDamage) Damage() int {
	return d.adapter.Damage() + d.damager.Damage()
}

func decorateSlashDamage(damager Damager, adapter *darkSoulsAdapter, dmg int) Damager {
	adapter.wp.AddSlash(dmg)
	return &darksoulsDamage{
		adapter: adapter,
		damager: damager,
	}
}

func decoratePierceDamage(damager Damager, adapter *darkSoulsAdapter, dmg int) Damager {
	adapter.wp.AddPierce(dmg)
	return &darksoulsDamage{
		adapter: adapter,
		damager: damager,
	}
}

func decorateStrikeDamage(damager Damager, adapter *darkSoulsAdapter, dmg int) Damager {
	adapter.wp.AddStrike(dmg)
	return &darksoulsDamage{
		adapter: adapter,
		damager: damager,
	}
}
