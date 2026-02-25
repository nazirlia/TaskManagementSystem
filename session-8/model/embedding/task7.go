package embedding

import "fmt"

type Mover interface {
	Move()
}

type Sayer interface {
	Say()
}

type Robot interface {
	Mover
	Sayer
}

type AutoBot struct{}

func (AutoBot) Move() {
	fmt.Println("Moving forward")
}

func (AutoBot) Say() {
	fmt.Println("I'm a robot")
}
