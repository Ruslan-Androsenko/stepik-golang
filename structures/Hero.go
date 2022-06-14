package structures

type Hero struct {
	On    bool
	Ammo  int
	Power int
}

func (this *Hero) Shoot() bool {
	answer := false

	if this.On && this.Ammo > 0 {
		answer = true
		this.Ammo--
	}

	return answer
}

func (this *Hero) Drive() bool {
	answer := false

	if this.On && this.Power > 0 {
		answer = true
		this.Power--
	}

	return answer
}
