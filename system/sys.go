package system

import (
	"fmt"
	"os"
	"strings"
)

func THPEnabled() (bool, error) {
	const thp string = "/sys/kernel/mm/transparent_hugepage/enabled"
	data, err := os.ReadFile(thp)
	if err != nil {
		return false, fmt.Errorf("cannot read file '%s'", data)
	}
	if strings.Contains(string(data), "[never]") {
		return false, nil
	}
	return true, nil
}
