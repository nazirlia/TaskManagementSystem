package table_driven

import "testing"

func TestReverseStringEmpty(t *testing.T) {
	result := ReverseString("")
	if result != "" {
		t.Errorf("Reverse string is %s, expected empty string", result)
	}
}

func TestReverseStringPalindrome(t *testing.T) {
	result := ReverseString("aziza")
	if result != "aziza" {
		t.Errorf("Reverse string is %s, expected 12321", result)
	}
}

func TestReverseStringSpecialCharacter(t *testing.T) {
	result := ReverseString("Niy@medd!n")
	if result != "n!ddem@yiN" {
		t.Errorf("Reverse string is %s, expected Niy@medd!n", result)
	}
}
