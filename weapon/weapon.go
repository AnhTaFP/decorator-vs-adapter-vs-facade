package weapon

// Sword represents a basic sword.
type Sword struct {
	BaseDamage int
	// other fields that describe a sword
}

// Damage returns the damage value of the sword.
func (s *Sword) Damage() int {
	return s.BaseDamage
}

// Bow represents a basic bow.
type Bow struct {
	BaseDamage int
	// other fields that describe a bow
}

// Damage returns the damage value of the bow.
func (b *Bow) Damage() int {
	return b.BaseDamage
}

// Staff represents a basic staff.
type Staff struct {
	BaseDamage int
	// other fields that describe a staff
}

// Damage returns the damage value of the staff.
func (s *Staff) Damage() int {
	return s.BaseDamage
}
