package game

type Character interface {
	Hit(target Character)
	Block(part string)
	IsAlive() bool
	GetName() string
	GetHP() int
	GetBlockPart() string
	TakeDamage(damage int)
}
