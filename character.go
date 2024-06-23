package main

type Character struct {
	GameObject

	Name string
}

func NewCharacter(props *Properties) *Character {
	return &Character{
		GameObject: *NewGameObject(props),
	}
}

type Ghost struct {
	Character
}

func NewGhost(props *Properties) *Ghost {
	return &Ghost{
		Character: *NewCharacter(props),
	}
}

func (g Ghost) Draw() {
	println("Drawing Ghost")
}

func (g Ghost) Update() {
	println("Updating Ghost")
}

func (g Ghost) Destroy() {
	println("Destroying Ghost")
}
