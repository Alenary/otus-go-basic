package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Alenary/otus-go-basic/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	byte, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(byte, &data)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	res := data
	return res, nil
}