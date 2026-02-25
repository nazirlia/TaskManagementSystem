package defining

type Speaker interface {
	Speak() string
}

type Dog struct {
	Audio string
}

type Person struct {
	Audio string
}

func (d Dog) Speak() string {
	return d.Audio
}

func (p Person) Speak() string {
	return p.Audio
}

func Speak(s Speaker) string {
	return s.Speak()
}
