package weapon

import "decorator-vs-adapter-vs-facade/darksouls"

type darkSoulsAdapter struct {
	*darksouls.Weapon
	damager Damager
}

func newDarkSoulsAdapter(damager Damager) *darkSoulsAdapter {
	weapon := darksouls.NewWeapon("weapon-id-1")

	return &darkSoulsAdapter{
		Weapon:  weapon,
		damager: damager,
	}
}

func (a *darkSoulsAdapter) Damage() int {
	return a.Hit() + a.damager.Damage()
}
