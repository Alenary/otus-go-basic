package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeBooks(t *testing.T) {
	tests := []struct {
		name    string
		books   []*Book
		wantErr bool
	}{
		{
			name: "simple case",
			books: []*Book{
				{Id: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925, Size: 180, Rate: 4.5},
				{Id: 2, Title: "1984", Author: "George Orwell", Year: 1949, Size: 328, Rate: 4.8},
			},
			wantErr: false,
		},
		{
			name:    "empty slice",
			books:   []*Book{},
			wantErr: false,
		},
		{
			name:    "nil slice",
			books:   nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := SerializeBooks(tt.books)
			if (err != nil) != tt.wantErr {
				t.Errorf("SerializeBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, data, "Serialized data should not be nil")
			}
		})
	}
}

func TestDeserializeBooks(t *testing.T) {
	tests := []struct {
		name    string
		books   []*Book
		wantErr bool
	}{
		{
			name: "simple case",
			books: []*Book{
				{Id: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925, Size: 180, Rate: 4.5},
				{Id: 2, Title: "1984", Author: "George Orwell", Year: 1949, Size: 328, Rate: 4.8},
			},
			wantErr: false,
		},
		{
			name:    "empty slice",
			books:   []*Book{},
			wantErr: false,
		},
		{
			name:    "nil slice",
			books:   nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := SerializeBooks(tt.books)
			if err != nil {
				t.Fatalf("Failed to serialize books: %v", err)
			}

			deserializedBooks, err := DeserializeBooks(data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeserializeBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if tt.books == nil {
					assert.Nil(t, deserializedBooks, "Deserialized books should be nil")
				} else {
					assert.Equal(t, len(tt.books), len(deserializedBooks),
						"Deserialized books should have the same length as the original")
					for i := range tt.books {
						assert.Equal(t, tt.books[i].Id, deserializedBooks[i].Id)
						assert.Equal(t, tt.books[i].Title, deserializedBooks[i].Title)
						assert.Equal(t, tt.books[i].Author, deserializedBooks[i].Author)
						assert.Equal(t, tt.books[i].Year, deserializedBooks[i].Year)
						assert.Equal(t, tt.books[i].Size, deserializedBooks[i].Size)
						assert.Equal(t, tt.books[i].Rate, deserializedBooks[i].Rate)
					}
				}
			}
		})
	}
}

func TestSerializeBooksJSON(t *testing.T) {
	tests := []struct {
		name    string
		books   []*Book
		wantErr bool
	}{
		{
			name: "simple case",
			books: []*Book{
				{Id: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925, Size: 180, Rate: 4.5},
				{Id: 2, Title: "1984", Author: "George Orwell", Year: 1949, Size: 328, Rate: 4.8},
			},
			wantErr: false,
		},
		{
			name:    "empty slice",
			books:   []*Book{},
			wantErr: false,
		},
		{
			name:    "nil slice",
			books:   nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := SerializeBooksJSON(tt.books)
			if (err != nil) != tt.wantErr {
				t.Errorf("SerializeBooksJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, data, "Serialized data should not be nil")
			}
		})
	}
}

func TestDeserializeBooksJSON(t *testing.T) {
	tests := []struct {
		name    string
		books   []*Book
		wantErr bool
	}{
		{
			name: "simple case",
			books: []*Book{
				{Id: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925, Size: 180, Rate: 4.5},
				{Id: 2, Title: "1984", Author: "George Orwell", Year: 1949, Size: 328, Rate: 4.8},
			},
			wantErr: false,
		},
		{
			name:    "empty slice",
			books:   []*Book{},
			wantErr: false,
		},
		{
			name:    "nil slice",
			books:   nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := SerializeBooksJSON(tt.books)
			if err != nil {
				t.Fatalf("Failed to serialize books to JSON: %v", err)
			}

			deserializedBooks, err := DeserializeBooksJSON(data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeserializeBooksJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if tt.books == nil {
					assert.Nil(t, deserializedBooks, "Deserialized books should be nil")
				} else {
					assert.Equal(t, len(tt.books), len(deserializedBooks),
						"Deserialized books should have the same length as the original")
					for i := range tt.books {
						assert.Equal(t, tt.books[i].Id, deserializedBooks[i].Id)
						assert.Equal(t, tt.books[i].Title, deserializedBooks[i].Title)
						assert.Equal(t, tt.books[i].Author, deserializedBooks[i].Author)
						assert.Equal(t, tt.books[i].Year, deserializedBooks[i].Year)
						assert.Equal(t, tt.books[i].Size, deserializedBooks[i].Size)
						assert.Equal(t, tt.books[i].Rate, deserializedBooks[i].Rate)
					}
				}
			}
		})
	}
}
