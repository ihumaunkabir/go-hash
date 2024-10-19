package hash

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGenerateHashFromText(t *testing.T) {
	text := "text"
	hash, err := GenerateHashFromText(text)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text)); err != nil {
		t.Fatalf("Hash does not match text: %v", err)
	}
}

func TestCompareHashAndText(t *testing.T) {
	text := "text"
	hash, err := GenerateHashFromText(text)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = CompareHashAndText(hash, text)
	if err != nil {
		t.Fatalf("Expected no error for matching hash and text, got %v", err)
	}

	wrongText := "wrong_text"
	err = CompareHashAndText(hash, wrongText)
	if err == nil {
		t.Fatal("Expected an error for non-matching hash and text, got none")
	}
}
