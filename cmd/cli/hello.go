package main

import (
	"fmt"

	"github.com/JohnL4/StarsWithoutNumber/internal/character"
)

func main() {
	c := character.Roll()
	fmt.Printf("%+v\n", c)
}
