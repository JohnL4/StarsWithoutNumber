package main

import(
	"fmt"
	"math/rand/v2"
) 

type Character struct {
	ST int	// Strength
	DX int	// Dexterity
	CO int	// Constitution
	IQ int	// Intelligence
	WI int	// Wisdom
	CH int	// Charisma
}

func main() {
	// fmt.Println("Hello, 世界")
	// var n int;
	var c Character;
	c.ST = dice(3, 6)
	c.DX = dice(3, 6)
	c.CO = dice(3, 6)
	c.IQ = dice(3, 6)
	c.WI = dice(3, 6)
	c.CH = dice(3, 6)
	fmt.Printf( "%+v\n", c)
}

func dice( count int, sides int) int {
	total := 0
	for i := 0; i < count; i++ {
		total += rand.IntN(sides) + 1
	}
	return total
}
