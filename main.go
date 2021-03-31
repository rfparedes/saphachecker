package main

import (
	"fmt"

	"github.com/rfparedes/saphachecker/ha"
	"github.com/rfparedes/saphachecker/system"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting up")
	str, err := ha.GetCorosyncDB()
	if err != nil {
		fmt.Println(err)
	}

	m1 := ha.Ec2CorosyncConfig()
	//m2 := ha.GceCorosyncConfig()
	//fmt.Println(ha.CorosyncCfgCompare(m1, m2))
	//ha.ReadCorosyncCfg("/etc/corosync/corosync.conf")
	cm, _ := ha.ProcessCorosyncDB(str)
	//fmt.Println(cm)
	ha.PrintCorosyncCfgDiff(ha.CorosyncCfgCompare(cm, m1))
	system.CheckLocalDiskSpace()

	system.Main()
}
