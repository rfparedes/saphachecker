package system

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func Main() {

	// Check THP
	enabled, err := THPEnabled()
	if err != nil {
		fmt.Println(err)
	}
	if enabled {
		logrus.Warning("disable transparent hugepages")
	} else {
		logrus.Info("transparent hugepages disabled OK")
	}
}
