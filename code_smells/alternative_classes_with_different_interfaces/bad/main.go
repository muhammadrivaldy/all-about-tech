package main

type IToyota interface {
	turnOnEngine()
	turnOffEngine()
	brakesEngine()
}

type toyota struct{}

func NewToyota() IToyota {
	return &toyota{}
}

func (t *toyota) turnOnEngine() {}

func (t *toyota) turnOffEngine() {}

func (t *toyota) brakesEngine() {}

type IHonda interface {
	activateEngine()
	deactivateEngine()
	takeABrakes()
}

type honda struct{}

func NewHonda() IHonda {
	return &honda{}
}

func (h *honda) activateEngine() {}

func (h *honda) deactivateEngine() {}

func (h *honda) takeABrakes() {}
