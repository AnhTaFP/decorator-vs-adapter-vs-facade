package main

import (
	"fmt"

	"decorator-vs-adapter-vs-facade/weapon"
)

func main() {
	baseSword := weapon.Sword{
		BaseDamage: 10,
	}

	fireSword := weapon.DecorateElementalDamage(&baseSword, &weapon.ElementalDamage{
		Element: weapon.Fire,
		Dmg:     3,
	})

	firePoisonSword := weapon.DecorateNonElementalDamage(fireSword, &weapon.NonElementalDamage{
		NonElement: weapon.Poison,
		Dmg:        1,
	})

	excalibur := weapon.DecorateNonElementalDamage(firePoisonSword, &weapon.NonElementalDamage{
		NonElement: weapon.Light,
		Dmg:        10,
	})

	fmt.Printf("excalibur now does %d damage\n", excalibur.Damage())

	// our warrior equips a standard claymore
	// claymore := weapon.Sword{
	// 	BaseDamage: 1,
	// }

	// claymoreAdapter := weapon.NewDarkSoulsAdapter(&claymore)

	// our warrior adds a point in slashing damage mastery
	// claymoreAdapter.AddSlash(10)

	// our warrior inserts a thunder gemstone
	// decoratedClaymore := weapon.DecorateElementalDamage(claymoreAdapter, &weapon.ElementalDamage{
	// Element: weapon.Thunder,
	// Dmg:     5,
	// })

	// our warrior adds a point in piercing damage mastery
	// claymoreAdapter.AddPierce(10)

	// our warrior inserts a dark gemstone
	// decoratedClaymore = weapon.DecorateNonElementalDamage(decoratedClaymore, &weapon.NonElementalDamage{
	// 	NonElement: weapon.Dark,
	// 	Dmg:        5,
	// })

	// fmt.Printf("claymore now does %d damage\n", decoratedClaymore.Damage())

	claymore := weapon.Sword{
		BaseDamage: 1,
	}

	enhancer := weapon.NewEnhancer(&claymore)
	enhancedClaymore := enhancer.Enhance(
		weapon.WithSlash(10),
		weapon.WithThunder(5),
		weapon.WithPierce(10),
		weapon.WithDark(5),
	)

	fmt.Printf("claymore now does %d damage\n", enhancedClaymore.Damage())
}
