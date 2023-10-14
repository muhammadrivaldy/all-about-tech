package main

type Appearance struct {
	Face string
	Skin string
}

func (a *Appearance) SetFace(face string) {
	a.Face = face
}

func (a *Appearance) SetSkin(skin string) {
	a.Skin = skin
}
