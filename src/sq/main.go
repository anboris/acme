package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	verbose := flag.Bool("v", false, "print which database is used")
	flag.Parse()

	root := gitRoot()
	db, err := findDB(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Fprintln(os.Stderr, "Using:", db)
	}

	if !hasInput() {
		fmt.Fprintln(os.Stderr, "no input detected")
		os.Exit(1)
	}

	cmd := exec.Command("sqlite3", db)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "sqlite3:", err)
		os.Exit(1)
	}
}

func gitRoot() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err == nil {
		return strings.TrimSpace(string(out))
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "."
	}
	return cwd
}

func findDB(root string) (string, error) {
	preferred := []string{"app.db", "main.db", "data.db"}

	for _, name := range preferred {
		path := filepath.Join(root, name)
		if info, err := os.Stat(path); err == nil && info.Mode().IsRegular() {
			return path, nil
		}
	}

	matches, _ := filepath.Glob(filepath.Join(root, "*.db"))
	if len(matches) > 0 {
		return matches[0], nil
	}
	return "", errors.New("no database found in root")
}

func hasInput() bool {
	info, _ := os.Stdin.Stat()
	return (info.Mode() & os.ModeCharDevice) == 0
}
