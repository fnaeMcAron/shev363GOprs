package game

type Player struct {
	Name      string
	HP        int
	Strength  int
	HitPart   string
	BlockPart string

	Inventory []Item
	Equipment []Item
}

func (p *Player) Hit(target Character) {
	if p.HitPart == target.GetBlockPart() {
		return
	}

	damage := p.Strength
	for _, item := range p.Equipment {
		damage += item.Attack
	}

	var finalDamage int
	switch t := target.(type) {
	case *Player:
		defence := 0
		for _, item := range t.Equipment {
			defence += item.Defence
		}
		finalDamage = damage - defence
	case *Enemy:
		finalDamage = damage
	default:
		finalDamage = damage
	}

	if finalDamage < 0 {
		finalDamage = 0
	}
	target.TakeDamage(finalDamage)
}

func (p *Player) Block(part string) {
	p.BlockPart = part
}

func (p *Player) IsAlive() bool {
	return p.HP > 0
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GetHP() int {
	return p.HP
}

func (p *Player) GetBlockPart() string {
	return p.BlockPart
}

func (p *Player) TakeDamage(damage int) {

	defence := 0

	for _, item := range p.Equipment {
		defence += item.Defence
	}

	finalDamage := damage - defence
	if finalDamage < 0 {
		finalDamage = 0
	}

	p.HP -= finalDamage
}
