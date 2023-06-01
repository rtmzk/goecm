package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func Exec(name string, args ...string) ([]byte, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("command args is none")
	}
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()

	return output, err
}

func ToString(data []string, sep string) string {
	var out = ""
	for _, item := range data {
		out = out + item + sep
	}
	return strings.TrimRight(out, sep)
}
