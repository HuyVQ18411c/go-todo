package utils

import (
	"testing"
)

type TestDto struct {
	Title       string
	Description string
}

func TestAssignValuesToStructFields(t *testing.T) {
	data := map[string]any{"title": "Homework", "description": "Do some home work"}
	var dto TestDto

	AssignValuesToStructFields(data, &dto)
	if dto.Title != "Homework" && dto.Description != "Do some home work" {
		t.Fatal("Failed to assign value")
	}
}
