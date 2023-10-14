package main

type ICar interface {
	turnOnEngine()
	turnOffEngine()
	brakesEngine()
}

type toyota struct{}

func NewToyota() ICar {
	return &toyota{}
}

func (t *toyota) turnOnEngine() {}

func (t *toyota) turnOffEngine() {}

func (t *toyota) brakesEngine() {}

type honda struct{}

func NewHonda() ICar {
	return &honda{}
}

func (h *honda) turnOnEngine() {}

func (h *honda) turnOffEngine() {}

func (h *honda) brakesEngine() {}
