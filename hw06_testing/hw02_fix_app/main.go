package main

import (
	"fmt"

	"github.com/Alenary/otus-go-basic/hw06_testing/hw02_fix_app/printer"
	"github.com/Alenary/otus-go-basic/hw06_testing/hw02_fix_app/reader"
	"github.com/Alenary/otus-go-basic/hw06_testing/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	} else {
		fmt.Println("passed")
	}

	staff, err = reader.ReadJSON(path)

	if err == nil {
		printer.PrintStaff(staff)
	}
}
