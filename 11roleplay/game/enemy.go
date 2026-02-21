package game

type Enemy struct {
	Name      string
	HP        int
	Strength  int
	HitPart   string
	BlockPart string

	Item Item
}

func (e *Enemy) Hit(target Character) {

	if e.HitPart == target.GetBlockPart() {
		return
	}

	damage := e.Strength
	target.TakeDamage(damage)
}

func (e *Enemy) Block(part string) {
	e.BlockPart = part
}

func (e *Enemy) IsAlive() bool {
	return e.HP > 0
}

func (e *Enemy) GetName() string {
	return e.Name
}

func (e *Enemy) GetHP() int {
	return e.HP
}

func (e *Enemy) GetBlockPart() string {
	return e.BlockPart
}

func (e *Enemy) TakeDamage(damage int) {
	e.HP -= damage
}
