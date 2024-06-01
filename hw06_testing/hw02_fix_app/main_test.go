package main

import (
	"os"
	"testing"

	"github.com/Alenary/otus-go-basic/hw06_testing/hw02_fix_app/reader"
	"github.com/Alenary/otus-go-basic/hw06_testing/hw02_fix_app/types"
)

func TestReadJSON(t *testing.T) {
	// Временные файлы для тестов.
	validFileContent := `[
        {"UserID": 123, "Age": 20, "Name": "TestA", "DepartmentID": 44},
        {"UserID": 124, "Age": 40, "Name": "TestB", "DepartmentID": 55},
        {"UserID": 125, "Age": 40, "Name": "TestC", "DepartmentID": 66}
    ]`
	validFilePath := "valid_employees.json"
	invalidFilePath := "invalid_employees.json"

	// Валидный JSON.
	err := os.WriteFile(validFilePath, []byte(validFileContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create valid JSON file: %v", err)
	}
	// Удаляем файл после тестов.
	defer os.Remove(validFilePath)

	// Невалидный JSON.
	err = os.WriteFile(invalidFilePath, []byte(`{invalid json}`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create invalid JSON file: %v", err)
	}
	// Удаляем файл после тестов.
	defer os.Remove(invalidFilePath)

	// Определяем таблицу тестов.
	tests := []struct {
		name        string
		filePath    string
		expectError bool
		expected    []types.Employee
	}{
		{
			name:        "Valid JSON file",
			filePath:    validFilePath,
			expectError: false,
			expected: []types.Employee{
				{UserID: 123, Age: 20, Name: "TestA", DepartmentID: 44},
				{UserID: 124, Age: 40, Name: "TestB", DepartmentID: 55},
				{UserID: 125, Age: 40, Name: "TestC", DepartmentID: 66},
			},
		},
		{
			name:        "Invalid JSON file",
			filePath:    invalidFilePath,
			expectError: true,
			expected:    nil,
		},
		{
			name:        "Non-existent file",
			filePath:    "non_existent_file.json",
			expectError: true,
			expected:    nil,
		},
	}

	// Проходим по всем тестам.
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result, readErr := reader.ReadJSON(tt.filePath)
			if (readErr != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, readErr)
			}
			if !tt.expectError && !equal(result, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

// equal - функция для сравнения двух срезов Employee.
func equal(a, b []types.Employee) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].UserID != b[i].UserID ||
			a[i].Name != b[i].Name ||
			a[i].Age != b[i].Age ||
			a[i].DepartmentID != b[i].DepartmentID {
			return false
		}
	}
	return true
}
