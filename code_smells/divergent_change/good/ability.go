package main

type Ability struct {
	Main         string
	SkillShoot   string
	SkillDefense string
}

func (a *Ability) SetMainAbility(mainAbility string) {
	a.Main = mainAbility
}

func (a *Ability) SetSkillShoot(skillShoot string) {
	a.SkillShoot = skillShoot
}

func (a *Ability) SetSkillDefense(skillDefense string) {
	a.SkillDefense = skillDefense
}
