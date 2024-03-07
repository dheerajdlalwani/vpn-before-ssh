package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	vpnStatus  = "warp-cli status"
	sshCommand = "ssh"
)

func isVPNEnabled() bool {
	output, err := exec.Command("bash", "-c", vpnStatus).Output()

	if err != nil {
		fmt.Println("Error checking VPN status:", err)
		return false
	}

	return strings.Contains(string(output), "Enabled")
}

func startDaemon() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()

		if strings.HasPrefix(command, sshCommand) {
			if !isVPNEnabled() {
				fmt.Println("ssssh. VPN not on!")
				continue
			}
		}
		fmt.Println(command)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
	}
}

func main() {
	startDaemon()
}
