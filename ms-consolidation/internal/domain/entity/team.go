package entity

type Team struct {
	ID      string
	Name    string
	Players []*Player
}

func NewTeam(id, name string) *Team {
	return &Team{
		ID:   id,
		Name: name,
	}
}

func (t *Team) AddPlayer(p *Player) {
	t.Players = append(t.Players, p)
}

func (t *Team) RemovePlayer(p *Player) {
	for i, tp := range t.Players {
		if p.ID == tp.ID {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			return
		}
	}
}
