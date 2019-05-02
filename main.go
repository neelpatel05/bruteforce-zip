package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

func bruteforce(filename string) {
	defer cleanup()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	words, err := ioutil.ReadAll(file)
	if err!=nil {
		panic(err.Error())
	}

	wordlist := strings.Split(string(words), "\n")

	for _,i = range wordlist {

	}
}

func one(reader *bufio.Reader) {
	defer cleanup()

	fmt.Print("$: Enter the string: ")
	string, err := reader.ReadString('\n')
	string = strings.TrimSpace(string)
	if err!=nil {
		panic(err.Error())
	}
	fmt.Print("$: Enter the length of password: ")
	length, err := reader.ReadString('\n')
	if err!=nil {
		panic(err.Error())
	}

	cmd := exec.Command("python", "generate_word.py", string, length)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err!=nil {
		panic(err.Error())
	}
	wordListStat, _ := os.Stat("word-list.txt")
	fmt.Println("-----------------------------\n",wordListStat.Name(),"\n",wordListStat.Size()/(1024),"KB\n")

	filename := "word-list.txt"
	bruteforce(filename)
}

func two(reader *bufio.Reader) {
	defer cleanup()

	fmt.Println("$: Enter the filename")
	_, err := reader.ReadString('\n')
	if err!=nil {
		panic(err.Error())
	}
}

func main() {
	defer cleanup()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$: 1. Create wordlist and attack 2. Use own wordlist and attact 3. Exit $: ")
		choice, err := reader.ReadString('\n')
		if err!=nil {
			panic(err.Error())
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
		}
	}
}
