package darksouls

type Physical int

const (
	Slash Physical = iota
	Pierce
	Strike
)

// Weapon represents a generic physical weapon.
type Weapon struct {
	ID           string
	slashDamage  int
	pierceDamage int
	strikeDamage int
}

// NewWeapon constructs a new generic physical weapon with a given ID.
func NewWeapon(ID string) *Weapon {
	return &Weapon{
		ID: ID,
	}
}

// AddSlash increments total slash damage.
func (w *Weapon) AddSlash(value int) {
	w.slashDamage += value
}

// AddPierce increments total pierce damage.
func (w *Weapon) AddPierce(value int) {
	w.pierceDamage += value
}

// AddStrike increments total strike damage.
func (w *Weapon) AddStrike(value int) {
	w.strikeDamage += value
}

// Hit returns the sum of all physical damages of weapon.
func (w *Weapon) Hit() int {
	return w.slashDamage + w.pierceDamage + w.strikeDamage
}
