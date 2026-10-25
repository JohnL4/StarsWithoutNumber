package character

import "math/rand/v2"

type Character struct {
	ST int `json:"ST"` // Strength
	DX int `json:"DX"` // Dexterity
	CO int `json:"CO"` // Constitution
	IQ int `json:"IQ"` // Intelligence
	WI int `json:"WI"` // Wisdom
	CH int `json:"CH"` // Charisma
}

// Roll returns a character with each attribute rolled on 3d6.
func Roll() Character {
	return Character{
		ST: dice(3, 6),
		DX: dice(3, 6),
		CO: dice(3, 6),
		IQ: dice(3, 6),
		WI: dice(3, 6),
		CH: dice(3, 6),
	}
}

func dice(count int, sides int) int {
	total := 0
	for i := 0; i < count; i++ {
		total += rand.IntN(sides) + 1
	}
	return total
}
