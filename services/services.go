package services

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

//IsServiceActive will return if a systemd service is enabled
func IsServiceActive(service string) (bool, error) {
	output, err := exec.Command("systemctl", "is-enabled", service+".service").Output()
	if err != nil {
		return false, log.Error("systemctl cannot determine if service %s is enabled", service)
	}
	if strings.TrimSpace(string(output)) == "enabled" {
		return true, nil
	}
	return false, nil
}
