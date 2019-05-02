package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func cleanup() {
	if r := recover(); r!=nil {
		log.Fatal(r)
	}
}

func one(reader *bufio.Reader) {
	defer cleanup()

	fmt.Println("$: Enter the string")
	string, err := reader.ReadString('\n')
	string = strings.TrimSpace(string)
	if err!=nil {
		panic(err)
	}
	fmt.Println("$: Enter the length of password")
	length, err := reader.ReadString('\n')
	if err!=nil {
		panic(err)
	}

	err = exec.Command("python", "generate_word.py", string, length).Run()
	if err!=nil {
		panic(err)
	}

}

func two(reader *bufio.Reader) {
	defer cleanup()

	fmt.Println("$: Enter the filename")
	_, err := reader.ReadString('\n')
	if err!=nil {
		panic(err)
	}
}

func main() {
	defer cleanup()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("$: 1. Create wordlist and attack 2. Use own wordlist and attact 3. Exit")
		choice, err := reader.ReadString('\n')
		if err!=nil {
			panic(err)
		}
		choice = strings.TrimSpace(choice)
		fmt.Println(choice)
		switch choice {
		case "1":
			one(reader)
		case "2":
			two(reader)
		case "3":
			break
			os.Exit(0)
		}
	}
}
