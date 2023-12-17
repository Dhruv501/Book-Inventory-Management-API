package model

import (
	"encoding/json"
	"testing"
)

func TestBookSerialization(t *testing.T) {

	book := Book{
		ID:          1,
		Title:       "Test Book",
		Author:      "Test Author",
		Genre:       "Test Genre",
		Data:        "Test Data",
		Price:       9.99,
		Description: "Test Description",
	}

	jsonData, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Failed to marshal book to JSON: %v", err)
	}

	var decodedBook Book
	err = json.Unmarshal(jsonData, &decodedBook)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to book: %v", err)
	}

	if decodedBook != book {
		t.Errorf("Expected decoded book to match original book, but they differ")
	}
}
