package helpers

import "testing"

type TestRegisterUserStruct struct {
	Name     string `validate:"required,min=3"`
	Email    string `validate:"required"`
	Password string `validate:"required,min=8"`
}

// Test 1: Valid struct, should return nil
func TestValidate_ValidStruct(t *testing.T) {
	input := TestRegisterUserStruct{
		Name:     "Arza",
		Email:    "arza@email.com",
		Password: "12345678",
	}

	result := Validate(input)

	if result != nil {
		t.Errorf("Validate() = %v, want nil", result)
	}
}

// Test 2: Empty field, should return error
func TestValidate_EmptyRequiredField(t *testing.T) {
	input := TestRegisterUserStruct{
		Name:     "",
		Email:    "",
		Password: "",
	}

	result := Validate(input)

	if result == nil {
		t.Fatal("Expected errors, got nil", result)
	}

	// Check error message
	expected := "Name wajib diisi."
	if result["Name"] != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result["Name"])
	}
}
