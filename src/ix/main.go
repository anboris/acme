package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(0)

	_, err := exec.LookPath("plumb")
	if err != nil {
		log.Fatal(err)
	}

	var global bool
	var file string // .index path

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get current working dir: %v", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("cannot get user home dir: %v", err)
	}

	flag.BoolVar(&global, "g", false, "open global index in home")
	flag.Parse()

	if global {
		file = filepath.Join(home, ".index")
	} else if cwd == home {
		file = filepath.Join(home, ".index")
	} else if gitRoot, err := gitRoot(cwd); err == nil {
		file = filepath.Join(gitRoot, ".index")
	} else {
		file = filepath.Join(cwd, ".index")
	}

	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			log.Fatalf("cannot create index: %v", err)
		}
		f.Close()
	} else if err != nil {
		log.Fatalf("error checking file: %v", err)
	}

	if err = open(file); err != nil {
		log.Fatalf("cannot sent to plumber: %v", err)
	}
}

func gitRoot(cwd string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = cwd
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func open(path string) error {
	pl := exec.Command("/usr/local/plan9/bin/plumb", "-d", "edit", path)
	if err := pl.Run(); err != nil {
		return err
	}
	return nil
}
