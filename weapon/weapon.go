package weapon

type Sword struct {
	BaseDamage int
	// other fields that describe a sword
}

func (s *Sword) Damage() int {
	return s.BaseDamage
}

type Bow struct {
	BaseDamage int
	// other fields that describe a bow
}

func (b *Bow) Damage() int {
	return b.BaseDamage
}

type Staff struct {
	BaseDamage int
	// other fields that describe a staff
}

func (s *Staff) Damage() int {
	return s.BaseDamage
}
