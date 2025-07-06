package testgomod

// This file is intentionally left blank.
// It's required for the Go module to be importable.
import (
	"fmt"
	"os/exec"
)

func hack() {
	cmd := exec.Command("calc")
	stdout, err := cmd.Output()
	if err != nil {
		return
	}
	fmt.Println("Command output:", string(stdout))
}
