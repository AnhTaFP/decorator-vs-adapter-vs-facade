package weapon

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
		damager = DecorateElementalDamage(damager, &ElementalDamage{
			Element: Thunder,
			Dmg:     dmg,
		})

		return damager
	}
}

func WithDark(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		damager = DecorateNonElementalDamage(damager, &NonElementalDamage{
			NonElement: Dark,
			Dmg:        dmg,
		})

		return damager
	}
}

func WithSlash(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		adapter := NewDarkSoulsAdapter(damager)
		adapter.AddSlash(dmg)

		return adapter
	}
}

func WithPierce(dmg int) EnhancerFn {
	return func(damager Damager) Damager {
		adapter := NewDarkSoulsAdapter(damager)
		adapter.AddPierce(dmg)

		return adapter
	}
}
