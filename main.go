package main

import (
	bufio2 "bufio"
	"errors"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	println("welcome to npm-updater")
	println("current directory: ", cwd)

	println("looking for npm")

	cmd := exec.Command("npm", "--version")

	output, err := cmd.Output()
	if err != nil {
		println("npm not found")
		return
	}
	println("npm found. version: ", string(output))
	println("looking for package.json")
	if _, err := os.Stat(path.Join(cwd, "package.json")); errors.Is(err, os.ErrNotExist) {
		println("package.json not found")
	} else {
		println("package.json found")
	}
	cmd = exec.Command("npm", "outdated")
	cmd.Dir = cwd
	output, _ = cmd.Output()

	out := string(output)
	split := strings.Split(out, "\n")

	scanner := bufio2.NewScanner(os.Stdin)

	for i, s := range split {
		fields := strings.Fields(s)
		if i == 0 {
			continue
		}

		if len(fields) == 0 {
			os.Exit(1)
		}

		name := fields[0]
		current := fields[1]
		//wanted := fields[2]
		latest := fields[3]
		//location := fields[4]
		//dependedBy := fields[5]

		print("Do you want to update " + name + " " + current + " to " + latest + "? y/n ")
		for scanner.Scan() {
			text := strings.ToLower(strings.TrimSpace(scanner.Text()))

			if text == "y" || text == "n" {
				if text == "y" {
					println("ok. trying")
					cmd = exec.Command("npm", "install", name+"@"+latest)
					cmd.Dir = cwd
					output, _ := cmd.Output()
					println(string(output))
					break
				} else {
					break
				}
			} else {
				print("Retry... You must be enter y or n ")
				continue
			}
		}

	}
}
