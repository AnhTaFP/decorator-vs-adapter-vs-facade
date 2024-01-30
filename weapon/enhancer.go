package weapon

type Damager interface {
	Damage() int
}

type Enhancer struct {
	damager Damager
}

func NewEnhancer(damager Damager) *Enhancer {
	return &Enhancer{
		damager: damager,
	}
}

func (e *Enhancer) Enhance(enhancerFns ...EnhancerFn) Damager {
	for _, enhancerFn := range enhancerFns {
		e.damager = enhancerFn(e.damager)
	}

	return e.damager
}

type EnhancerFn func(damager Damager) Damager

func WithThunder(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateElementalDamage(damager, &elementalDamage{
			element: thunder,
			dmg:     dmg,
		})

		return damager
	}
}

func WithFire(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateElementalDamage(damager, &elementalDamage{
			element: fire,
			dmg:     dmg,
		})

		return damager
	}
}

func WithWater(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateElementalDamage(damager, &elementalDamage{
			element: water,
			dmg:     dmg,
		})

		return damager
	}
}

func WithEarth(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateElementalDamage(damager, &elementalDamage{
			element: earth,
			dmg:     dmg,
		})

		return damager
	}
}

func WithDark(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateNonElementalDamage(damager, &nonElementalDamage{
			nonElement: dark,
			dmg:        dmg,
		})

		return damager
	}
}

func WithLight(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateNonElementalDamage(damager, &nonElementalDamage{
			nonElement: light,
			dmg:        dmg,
		})

		return damager
	}
}

func WithPoison(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateNonElementalDamage(damager, &nonElementalDamage{
			nonElement: poison,
			dmg:        dmg,
		})

		return damager
	}
}

func WithAcid(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateNonElementalDamage(damager, &nonElementalDamage{
			nonElement: acid,
			dmg:        dmg,
		})

		return damager
	}
}

func WithChaos() EnhancerFn {
	return func(damager Damager) Damager {
		damager = decorateChaosDamage(damager, &chaosDamage{})

		return damager
	}
}

func WithSlash(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		adapter := newDarkSoulsAdapter(damager)
		adapter.AddSlash(dmg)

		return adapter
	}
}

func WithPierce(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		adapter := newDarkSoulsAdapter(damager)
		adapter.AddPierce(dmg)

		return adapter
	}
}

func WithStrike(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		adapter := newDarkSoulsAdapter(damager)
		adapter.AddStrike(dmg)

		return adapter
	}
}
