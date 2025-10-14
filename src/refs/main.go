package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tag, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "refs: failed to read tag from stdin: %v\n", err)
		os.Exit(1)
	}
	tag = strings.TrimSpace(tag)

	parts := strings.Fields(tag)
	if len(parts) == 0 {
		fmt.Fprintf(os.Stderr, "refs: tag is empty or malformed.\n")
		os.Exit(1)
	}

	filePath := parts[0]
	dir := filepath.Dir(filePath)

	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stderr, "refs: failed to change directory to %s: %v\n", dir, err)
		os.Exit(1)
	}

	cmd := exec.Command("L", "refs")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "acmerefs: 'L refs' failed to execute: %v\n", err)
		os.Exit(1)
	}
}
