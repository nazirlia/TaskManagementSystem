package unit_testing

import "testing"

func TestNewUserProfileIsMinorTrue(t *testing.T) {
	result := NewUserProfile("Gulsenem", 17)
	if !result.IsMinor {
		t.Errorf("IsMinor is false, expected true")
	}
	if result.Name != "Gulsenem" {
		t.Errorf("Name is %s, expected Gulsenem", result.Name)
	}
	if result.Age != 17 {
		t.Errorf("Age is %d, expected 17", result.Age)
	}
}
func TestNewUserProfileIsMinorFalse(t *testing.T) {
	result := NewUserProfile("Gulsenem", 19)
	if result.IsMinor {
		t.Errorf("IsMinor is true, expected false")
	}
	if result.Name != "Gulsenem" {
		t.Errorf("Name is %s, expected Gulsenem", result.Name)
	}
	if result.Age != 19 {
		t.Errorf("Age is %d, expected 17", result.Age)
	}
}
