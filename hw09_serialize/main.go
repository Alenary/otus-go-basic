package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
)

// SerializeBooks сериализует слайс объектов Book в байтовый массив.
func SerializeBooks(books []*Book) ([]byte, error) {
	bookList := &BookList{
		Books: books,
	}
	return proto.Marshal(bookList)
}

// DeserializeBooks десериализует байтовый массив в слайс объектов Book.
func DeserializeBooks(data []byte) ([]*Book, error) {
	bookList := &BookList{}
	err := proto.Unmarshal(data, bookList)
	if err != nil {
		return nil, err
	}
	return bookList.Books, nil
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

	fmt.Println("Serialized books:", data)

	deserializedBooks, err := DeserializeBooks(data)
	if err != nil {
		log.Fatalf("Error deserializing books: %v", err)
	}

	fmt.Println("Deserialized books:", deserializedBooks)
}
