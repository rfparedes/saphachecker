package system

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/sirupsen/logrus"
)

const diskWarningPercent float64 = 95.0

func CheckLocalDiskSpace() {
	parts, _ := disk.Partitions(false)
	for _, p := range parts {
		device := p.Mountpoint
		s, _ := disk.Usage(device)

		if s.Total == 0 {
			continue
		}
		if s.UsedPercent >= diskWarningPercent {
			logrus.Warning("local disk space usage is high: ", p.Mountpoint)
		} else {
			logrus.Info("local disk space usage OK: ", p.Mountpoint)
		}
	}
}
