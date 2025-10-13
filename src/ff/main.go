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
	var pattern string

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		in := bufio.NewScanner(os.Stdin)
		if in.Scan() {
			pattern = strings.TrimSpace(in.Text())
		}
	} else {
		if len(os.Args) < 2 {
			fmt.Fprintln(os.Stderr, "usage: ff <pattern>")
			os.Exit(1)
		}
		pattern = strings.TrimSpace(os.Args[1])
	}

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

	cmd := exec.Command("find", ".", "-type", "f", "-name", pattern+"*")
	cmd.Dir = root
	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("find pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("find start: %v", err)
	}

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		absPath := filepath.Clean(filepath.Join(root, line))
		newPath, err := filepath.Rel(pwd, absPath)
		if err != nil {
			continue
		}
		fmt.Println(newPath)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("reading find output: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalf("find wait: %v", err)
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
