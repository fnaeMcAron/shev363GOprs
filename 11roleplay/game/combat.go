package game

import "math/rand"

var Parts = []string{"голова", "тело", "ноги"}

func RandomPart() string {
	return Parts[rand.Intn(len(Parts))]
}

func ProcessTurn(p1 Character, p2 Character) {

	p1.Hit(p2)
	p2.Hit(p1)
}
