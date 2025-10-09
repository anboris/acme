package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type line struct {
	winid  int
	action string
	fname  string
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("editinacme: ")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: editinacme file\n")
		os.Exit(2)
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
	}

	file := flag.Arg(0)

	fullpath, err := filepath.Abs(file)
	if err != nil {
		log.Fatal(err)
	}
	file = fullpath

	log.Printf("editing %s", file)

	out, err := exec.Command("plumb", "-d", "edit", file).CombinedOutput()
	if err != nil {
		log.Fatalf("executing plumb: %v\n%s", err, out)
	}

	cmd := exec.Command("9p", "read", "acme/log")
	stdout, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(stdout)
	cmd.Start()

	for scanner.Scan() {
		m := scanner.Text()
		parts := strings.Fields(m)

		if len(parts) < 3 {
			continue
		}

		winid, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting winid to int:", err)
		}

		l := line{winid: winid, action: parts[1], fname: parts[2]}

		if l.action == "del" && l.fname == file {
			break
		}
	}
}
