package reader

import (
	"encoding/json"
	"os"

	"github.com/Alenary/otus-go-basic/hw06_testing/hw02_fix_app/types"
)

// ReadJSON читает JSON файл и возвращает массив Employee
func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var employees []types.Employee
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&employees); err != nil {
		return nil, err
	}

	return employees, nil
}
