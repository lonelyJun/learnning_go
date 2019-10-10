package main

import "fmt"

func main() {
	var attack = 40

	var defence = 20

	var damageRate float32 = 0.17

	var damage = float32(attack-defence) * damageRate

	fmt.Println(damage)

}
