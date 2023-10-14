package main

type Character struct {
	Name         string
	Face         string
	Skin         string
	MainAbility  string
	SkillShoot   string
	SkillDefense string
	Health       int
	Power        int
}

func (c *Character) SetName(name string) {
	c.Name = name
}

func (c *Character) SetFace(face string) {
	c.Face = face
}

func (c *Character) SetSkin(skin string) {
	c.Skin = skin
}

func (c *Character) SetMainAbility(mainAbility string) {
	c.MainAbility = mainAbility
}

func (c *Character) SetSkillShoot(skillShoot string) {
	c.SkillShoot = skillShoot
}

func (c *Character) SetSkillDefense(skillDefense string) {
	c.SkillDefense = skillDefense
}

func (c *Character) SetHealth(health int) {
	c.Health = health
}

func (c *Character) SetPower(power int) {
	c.Power = power
}
