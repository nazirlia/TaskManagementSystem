package mocking

import (
	"errors"
	"testing"
)

type MockDatabase struct{}

func (m MockDatabase) GetUserByID(id int) (User, error) {
	if id == 1 {
		return User{ID: 1, Name: "Ehmediyye"}, nil
	}
	return User{}, errors.New("user not found")
}

func TestGetUserValidID(t *testing.T) {
	db := &MockDatabase{}
	result, err := db.GetUserByID(1)
	if err != nil {
		t.Errorf("GetUser returned unexpected error: %v", err)
	}
	if result.ID != 1 {
		t.Errorf("GetUser returned ID %d, expected 1", result.ID)
	}
}

func TestGetUserInvalidID(t *testing.T) {
	db := &MockDatabase{}
	_, err := db.GetUserByID(0)
	if err == nil {
		t.Errorf("GetUser expected error for ID 0, got nil")

	}
}
