package weapon

import "decorator-vs-adapter-vs-facade/darksouls"

type DarkSoulsAdapter struct {
	*darksouls.Weapon
	damager Damager
}

func NewDarkSoulsAdapter(damager Damager) *DarkSoulsAdapter {
	weapon := darksouls.NewWeapon("weapon-id-1")

	return &DarkSoulsAdapter{
		Weapon:  weapon,
		damager: damager,
	}
}

func (a *DarkSoulsAdapter) Damage() int {
	return a.Hit() + a.damager.Damage()
}
