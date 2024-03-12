package weapon

import "decorator-vs-adapter-vs-facade/darksouls"

type darkSoulsAdapter struct {
	wp *darksouls.Weapon
}

func newDarkSoulsAdapter(wp *darksouls.Weapon) *darkSoulsAdapter {
	return &darkSoulsAdapter{
		wp: wp,
	}
}

func (a *darkSoulsAdapter) Damage() int {
	return a.wp.Hit()
}
