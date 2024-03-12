package weapon

import (
	"decorator-vs-adapter-vs-facade/darksouls"
)

// Damager represents a type that can damage others.
type Damager interface {
	Damage() int
}

// Enhancer enhances basic weapons with various effects such as
// elemental damages (e.g. fire), non elemental damages (e.g. poison)
// and many other damage types as required.
type Enhancer struct {
	damager Damager
}

// NewEnhancer returns a new Enhancer.
func NewEnhancer(damager Damager) *Enhancer {
	return &Enhancer{
		damager: damager,
	}
}

// Enhance adds various damage altering effects to the base weapons that
// correspond to the provided enhancer functions.
func (e *Enhancer) Enhance(enhancerFns ...EnhancerFn) Damager {
	for _, enhancerFn := range enhancerFns {
		e.damager = enhancerFn(e.damager)
	}

	return e.damager
}

// EnhancerFn represents a function that takes a damager, and strengthens it.
type EnhancerFn func(damager Damager) Damager

// WithThunder adds thunder damage.
func WithThunder(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateElementalDamage(damager, &elementalDamage{
			element: thunder,
			dmg:     dmg,
		})

		return damager
	}
}

// WithFire adds fire damage.
func WithFire(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateElementalDamage(damager, &elementalDamage{
			element: fire,
			dmg:     dmg,
		})

		return damager
	}
}

// WithWater adds water damage.
func WithWater(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateElementalDamage(damager, &elementalDamage{
			element: water,
			dmg:     dmg,
		})

		return damager
	}
}

// WithEarth adds earth damage.
func WithEarth(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateElementalDamage(damager, &elementalDamage{
			element: earth,
			dmg:     dmg,
		})

		return damager
	}
}

// WithDark adds dark damage.
func WithDark(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateNonElementalDamage(damager, &nonElementalDamage{
			nonElement: dark,
			dmg:        dmg,
		})

		return damager
	}
}

// WithLight adds light damage.
func WithLight(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateNonElementalDamage(damager, &nonElementalDamage{
			nonElement: light,
			dmg:        dmg,
		})

		return damager
	}
}

// WithPoison adds poison damage.
func WithPoison(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateNonElementalDamage(damager, &nonElementalDamage{
			nonElement: poison,
			dmg:        dmg,
		})

		return damager
	}
}

// WithAcid adds acid damage.
func WithAcid(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateNonElementalDamage(damager, &nonElementalDamage{
			nonElement: acid,
			dmg:        dmg,
		})

		return damager
	}
}

// WithChaos adds chaos damage.
func WithChaos() EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateChaosDamage(damager, &chaosDamage{})

		return damager
	}
}

// WithSlash adds slash damage.
func WithSlash(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		wp := darksouls.NewWeapon("slash-weapon-id")
		adapter := newDarkSoulsAdapter(wp)

		return decorateSlashDamage(damager, adapter, dmg)
	}
}

// WithPierce adds pierce damage.
func WithPierce(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		wp := darksouls.NewWeapon("pierce-weapon-id")
		adapter := newDarkSoulsAdapter(wp)

		return decoratePierceDamage(damager, adapter, dmg)
	}
}

// WithStrike adds strike damage.
func WithStrike(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		wp := darksouls.NewWeapon("strike-weapon-id")
		adapter := newDarkSoulsAdapter(wp)

		return decorateStrikeDamage(damager, adapter, dmg)
	}
}
