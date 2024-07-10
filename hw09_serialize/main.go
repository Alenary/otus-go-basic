package main

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

// Реализация интерфейса json.Marshaler для Book
func (b *Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	})
}

// Реализация интерфейса json.Unmarshaler для Book
func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux)
}

// SerializeBooks сериализует слайс объектов Book в байтовый массив (Protobuf).
func SerializeBooks(books []*Book) ([]byte, error) {
	bookList := &BookList{
		Books: books,
	}
	return proto.Marshal(bookList)
}

// DeserializeBooks десериализует байтовый массив в слайс объектов Book (Protobuf).
func DeserializeBooks(data []byte) ([]*Book, error) {
	bookList := &BookList{}
	err := proto.Unmarshal(data, bookList)
	if err != nil {
		return nil, err
	}
	return bookList.Books, nil
}

// SerializeBooksJSON сериализует слайс объектов Book в JSON.
func SerializeBooksJSON(books []*Book) ([]byte, error) {
	return json.Marshal(books)
}

// DeserializeBooksJSON десериализует JSON в слайс объектов Book.
func DeserializeBooksJSON(data []byte) ([]*Book, error) {
	var books []*Book
	err := json.Unmarshal(data, &books)
	return books, err
}

func main() {
	books := []*Book{
		{Id: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925, Size: 180, Rate: 4.5},
		{Id: 2, Title: "1984", Author: "George Orwell", Year: 1949, Size: 328, Rate: 4.8},
	}

	data, err := SerializeBooks(books)
	if err != nil {
		log.Fatalf("Error serializing books: %v", err)
	}

	fmt.Println("Serialized books (Protobuf):", data)

	deserializedBooks, err := DeserializeBooks(data)
	if err != nil {
		log.Fatalf("Error deserializing books: %v", err)
	}

	fmt.Println("Deserialized books (Protobuf):", deserializedBooks)

	jsonData, err := SerializeBooksJSON(books)
	if err != nil {
		log.Fatalf("Error serializing books to JSON: %v", err)
	}

	fmt.Println("Serialized books (JSON):", string(jsonData))

	deserializedBooksJSON, err := DeserializeBooksJSON(jsonData)
	if err != nil {
		log.Fatalf("Error deserializing books from JSON: %v", err)
	}

	fmt.Println("Deserialized books (JSON):", deserializedBooksJSON)
}
