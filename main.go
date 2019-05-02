package main

import (
	"os/exec"
)

func main() {

	_ = exec.Command("python","generate_word.py").Run()
}
