package main

type Character struct {
	Name   string
	Health int
	Power  int
}

func (c *Character) SetName(name string) {
	c.Name = name
}

func (c *Character) SetHealth(health int) {
	c.Health = health
}

func (c *Character) SetPower(power int) {
	c.Power = power
}
