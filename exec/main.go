// Executing external commands
package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("ls", "-l")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))

	stdout, _ = exec.Command("/bin/zsh", "-c", "ps -ef | grep zsh").Output()
	fmt.Printf("%s", stdout)

	stdout, _ = exec.Command("hostname").Output()
	if strings.Contains(string(stdout), "local") {
		fmt.Println("hostname has local")
	}
}
