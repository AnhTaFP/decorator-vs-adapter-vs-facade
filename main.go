package main

import (
	"fmt"

	"decorator-vs-adapter-vs-facade/weapon"
)

func main() {
	// standard sword
	baseSword := weapon.Sword{
		BaseDamage: 10,
	}

	swordEnhancer := weapon.NewEnhancer(&baseSword)
	excalibur := swordEnhancer.Enhance(
		weapon.WithFire(3),   // inserts a fire gemstone
		weapon.WithPoison(1), // puts a skill point in poison mastery
		weapon.WithLight(10), // finds a holy shrine and imbues the sword with power of light
	)

	fmt.Printf("sword now does %d damage\n", excalibur.Damage())

	// standard claymore
	claymore := weapon.Sword{
		BaseDamage: 1,
	}

	claymoreEnhancer := weapon.NewEnhancer(&claymore)
	enhancedClaymore := claymoreEnhancer.Enhance(
		weapon.WithSlash(10),  // puts a skill point in slashing mastery
		weapon.WithThunder(5), // inserts a thunder gemstone
		weapon.WithPierce(10), // puts a skill point in piercing master
		weapon.WithDark(5),    // inserts a dark gemstone
	)

	fmt.Printf("claymore now does %d damage\n", enhancedClaymore.Damage())
}
