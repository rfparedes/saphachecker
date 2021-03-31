package ha

import (
	"bufio"
	"fmt"
	_ "github.com/rfparedes/saphachecker/log"
	"github.com/sirupsen/logrus"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
)

// AWS EC2 default corosync configuration
var Ec2CorosyncConfig = func() map[string]string {
	return map[string]string{
		// if prepended by runtime, then this is not a parameter that is set
		"totem.version":   "2",
		"totem.token":     "30000",
		"totem.consensus": "36000",
		"totem.token_retransmits_before_loss_const": "6",
		"totem.secauth":                       "on",
		"totem.crypto_hash":                   "sha1",
		"totem.crypto_cipher":                 "aes256",
		"totem.clear_node_high_bit":           "yes",
		"totem.interface.0.bindnetaddr":       "VARIABLE",
		"totem.interface.0.mcastport":         "5405",
		"totem.interface.0.ttl":               "1",
		"totem.transport":                     "udpu",
		"logging.fileline":                    "off",
		"logging.to_logfile":                  "yes",
		"logging.to_syslog":                   "yes",
		"logging.logfile":                     "/var/log/cluster/corosync.log",
		"logging.debug":                       "off",
		"logging.timestamp":                   "on",
		"logging.logger_subsys.QUORUM.subsys": "QUORUM",
		"logging.logger_subsys.QUORUM.debug":  "off",
		"nodelist.node.0.nodeid":              "1",
		"nodelist.node.1.nodeid":              "2",
		"nodelist.node.0.ring0_addr":          "VARIABLE",
		"nodelist.node.1.ring0_addr":          "VARIABLE",
		"nodelist.node.0.ring1_addr":          "VARIABLE",
		"nodelist.node.1.ring1_addr":          "VARIABLE",
		"quorum.provider":                     "corosync_votequorum",
		"quorum.expected_votes":               "2",
		"quorum.two_node":                     "1",
	}
}

// Azure Virtual Machines default corosync configuration
var AzureCorosyncConfig = func() map[string]string {
	return map[string]string{
		// if prepended by runtime, then this is not a parameter that is set
		"totem.version":   "2",
		"totem.token":     "30000",
		"totem.consensus": "36000",
		"totem.token_retransmits_before_loss_const": "10",
		"totem.join":                          "60",
		"totem.max_messages":                  "20",
		"totem.cluster_name":                  "VARIABLE",
		"totem.secauth":                       "on",
		"totem.crypto_hash":                   "sha1",
		"totem.crypto_cipher":                 "aes256",
		"totem.clear_node_high_bit":           "yes",
		"totem.interface.0.bindnetaddr":       "VARIABLE",
		"totem.interface.0.mcastport":         "5405",
		"totem.interface.0.ttl":               "1",
		"totem.transport":                     "udpu",
		"logging.fileline":                    "off",
		"logging.to_logfile":                  "yes",
		"logging.to_stderr":                   "no",
		"logging.to_syslog":                   "yes",
		"logging.logfile":                     "/var/log/cluster/corosync.log",
		"logging.debug":                       "off",
		"logging.timestamp":                   "on",
		"logging.logger_subsys.QUORUM.subsys": "QUORUM",
		"logging.logger_subsys.QUORUM.debug":  "off",
		"nodelist.node.0.nodeid":              "1",
		"nodelist.node.1.nodeid":              "2",
		"nodelist.node.0.ring0_addr":          "VARIABLE",
		"nodelist.node.1.ring0_addr":          "VARIABLE",
		"quorum.provider":                     "corosync_votequorum",
		"quorum.expected_votes":               "2",
		"quorum.two_node":                     "1",
	}
}

// Google Cloud Engine default corosync configuration
var GceCorosyncConfig = func() map[string]string {
	return map[string]string{
		// if prepended by runtime, then this is not a parameter that is set
		"totem.version":   "2",
		"totem.token":     "20000",
		"totem.consensus": "24000",
		"totem.token_retransmits_before_loss_const": "10",
		"totem.join":                          "60",
		"totem.max_messages":                  "20",
		"totem.cluster_name":                  "hacluster",
		"totem.secauth":                       "on",
		"totem.crypto_hash":                   "sha1",
		"totem.crypto_cipher":                 "aes256",
		"totem.clear_node_high_bit":           "yes",
		"totem.interface.0.bindnetaddr":       "VARIABLE",
		"totem.interface.0.mcastport":         "5405",
		"totem.interface.0.ttl":               "1",
		"totem.transport":                     "udpu",
		"logging.fileline":                    "off",
		"logging.to_logfile":                  "no",
		"logging.to_stderr":                   "no",
		"logging.to_syslog":                   "yes",
		"logging.logfile":                     "/var/log/cluster/corosync.log",
		"logging.debug":                       "off",
		"logging.timestamp":                   "on",
		"logging.logger_subsys.QUORUM.subsys": "QUORUM",
		"logging.logger_subsys.QUORUM.debug":  "off",
		"nodelist.node.0.nodeid":              "1",
		"nodelist.node.1.nodeid":              "2",
		"nodelist.node.0.ring0_addr":          "VARIABLE",
		"nodelist.node.1.ring0_addr":          "VARIABLE",
		"quorum.provider":                     "corosync_votequorum",
		"quorum.expected_votes":               "2",
		"quorum.two_node":                     "1",
	}
}

// On-premise default corosync configuration
var OnpremiseCorosyncConfig = func() map[string]string {
	return map[string]string{
		// if prepended by runtime, then this is not a parameter that is set
		"totem.version":   "2",
		"totem.token":     "5000",
		"totem.consensus": "6000",
		"totem.token_retransmits_before_loss_const": "10",
		"totem.join":                          "60",
		"totem.max_messages":                  "20",
		"totem.cluster_name":                  "VARIABLE",
		"totem.secauth":                       "on",
		"totem.crypto_hash":                   "sha1",
		"totem.crypto_cipher":                 "aes256",
		"totem.clear_node_high_bit":           "yes",
		"totem.interface.0.mcastport":         "5405",
		"totem.interface.0.ttl":               "1",
		"totem.transport":                     "udpu",
		"logging.fileline":                    "off",
		"logging.to_logfile":                  "no",
		"logging.to_stderr":                   "no",
		"logging.to_syslog":                   "yes",
		"logging.logfile":                     "/var/log/cluster/corosync.log",
		"logging.debug":                       "off",
		"logging.timestamp":                   "on",
		"logging.logger_subsys.QUORUM.subsys": "QUORUM",
		"logging.logger_subsys.QUORUM.debug":  "off",
		"nodelist.node.0.nodeid":              "1",
		"nodelist.node.1.nodeid":              "2",
		"nodelist.node.0.ring0_addr":          "VARIABLE",
		"nodelist.node.1.ring0_addr":          "VARIABLE",
		"quorum.provider":                     "corosync_votequorum",
		"quorum.expected_votes":               "2",
		"quorum.two_node":                     "1",
	}
}

// GetCorosyncDB retrieves corosync configuration of system
func GetCorosyncDB() (string, error) {
	cmdName := "corosync-cmapctl"
	path, err := exec.LookPath(cmdName)
	if err != nil {
		return "", fmt.Errorf("command '%s' not found", cmdName)
	}
	cmdOut, err := exec.Command(path).Output()
	if err != nil {
		return "", fmt.Errorf("command '%s' output error", cmdName)
	}
	return string(cmdOut), nil
}

// ProcessCorosyncDB reads the corosync config into a map
func ProcessCorosyncDB(str string) (map[string]string, error) {
	m := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		var key string
		// skip over these parameters as they are runtime or internal
		var re = regexp.MustCompile(`config|internal|runtime|uidgid|local_node_pos`)
		line := scanner.Text()
		if re.MatchString(line) {
			continue
		}
		for i, p := range strings.Fields(line) {
			// Only read in the param and value, leave out type and =
			if i == 0 {
				key = p
			} else if i == 3 {
				m[key] = p
			}
		}

	}
	return m, nil
}

// CorosyncCfgEqual determines if two maps are equal or not
func CorosyncCfgEqual(m1 map[string]string, m2 map[string]string) bool {
	eq := reflect.DeepEqual(m1, m2)
	return eq
}

// CorosyncCfgCompare will compare two maps and return the differences
func CorosyncCfgCompare(customerMap map[string]string, validatedMap map[string]string) map[string][]string {

	diffMap := make(map[string][]string)
	// Range over the validated Map
	for vk, vv := range validatedMap {
		found := false
		// Range the customer map
		for ck, cv := range customerMap {
			if ck == vk {
				found = true
				//if the values don't match
				if vv != cv && vv != "VARIABLE" {
					diffMap[ck] = []string{cv, vv}
				}
				break
			}
		}
		// The customer config doesn't contain a validated config param
		if !found {
			diffMap[vk] = []string{"-1", vv}
		}
	}
	// Need to search the opposite way also in case customer has values that validated doesn't
	for ck, cv := range customerMap {
		found := false
		for vk := range validatedMap {
			if ck == vk {
				found = true
				break
			}
		}
		if !found {
			diffMap[ck] = []string{"-2", cv}
		}
	}
	return diffMap
}

// PrintCorosyncCfgDiff will print out the differences of diffMap
func PrintCorosyncCfgDiff(diffMap map[string][]string) {

	// Handle if there is no difference
	if len(diffMap) != 0 {
		for k, v := range diffMap {
			if v[0] == "-1" {
				logrus.WithFields(logrus.Fields{
					"value": v[1],
				}).Warning("corosync.conf missing parameter: \033[32m", k)
			} else if v[0] == "-2" {
				logrus.WithFields(logrus.Fields{
					"value": v[1],
				}).Warning("corosync.conf additional parameter not validated: remove \033[32m", k)
			} else {
				logrus.WithFields(logrus.Fields{
					"from": v[0],
					"to":   v[1],
				}).Warning("corosync.conf parameter difference: change \033[32m", k)
			}
		}
	}
}
