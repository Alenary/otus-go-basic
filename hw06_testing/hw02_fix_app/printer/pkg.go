package printer

import (
	"fmt"

	"github.com/Alenary/otus-go-basic/hw06_testing/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
