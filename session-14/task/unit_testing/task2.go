package unit_testing

type UserProfile struct {
	Name    string
	Age     int
	IsMinor bool
}

func NewUserProfile(name string, age int) UserProfile {
	var isMinor bool
	if age < 18 {
		isMinor = true
	} else {
		isMinor = false
	}
	return UserProfile{
		Name:    name,
		Age:     age,
		IsMinor: isMinor,
	}
}
