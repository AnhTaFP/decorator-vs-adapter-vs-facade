package main

import (
	"fmt"

	"decorator-vs-adapter-vs-facade/weapon"
)

func main() {
	baseSword := weapon.Sword{
		BaseDamage: 10,
	}

	swordEnhancer := weapon.NewEnhancer(&baseSword)
	excalibur := swordEnhancer.Enhance(
		weapon.WithFire(3),
		weapon.WithPoison(1),
		weapon.WithLight(10),
	)

	fmt.Printf("sword now does %d damage\n", excalibur.Damage())

	claymore := weapon.Sword{
		BaseDamage: 1,
	}

	claymoreEnhancer := weapon.NewEnhancer(&claymore)
	enhancedClaymore := claymoreEnhancer.Enhance(
		weapon.WithSlash(10),
		weapon.WithThunder(5),
		weapon.WithPierce(10),
		weapon.WithDark(5),
	)

	fmt.Printf("claymore now does %d damage\n", enhancedClaymore.Damage())
}
