package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatalf("no pattern supplied")
	}

	pattern := strings.TrimSpace(scanner.Text())
	if pattern == "" {
		log.Fatalf("empty pattern")
	}

	root, err := getGitRoot()
	if err != nil {
		log.Fatalf("cannot get git repo root: %v", err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get pwd: %v", err)
	}

	cmd := exec.Command("git", "grep", "-n", "--no-color", pattern)
	cmd.Dir = root
	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("git grep pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("git grep start: %v", err)
	}

	scanner = bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 3)
		if len(parts) < 3 {
			continue
		}

		relPath := parts[0]
		absPath := filepath.Clean(filepath.Join(root, relPath))

		// skip anything outside the repo
		if !strings.HasPrefix(absPath, filepath.Clean(root)) {
			continue
		}

		newPath, err := filepath.Rel(pwd, absPath)
		if err != nil {
			continue
		}

		fmt.Printf("%s:%s:%s\n", newPath, parts[1], parts[2])
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("reading git grep: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		if e, ok := err.(*exec.ExitError); ok && e.ExitCode() == 1 {
			return // no matches
		}
		log.Fatalf("git grep wait: %v", err)
	}
}

func getGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
