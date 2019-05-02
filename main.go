package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"github.com/alexmullins/zip"
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

	wordListStat, _ := os.Stat("word-list.txt")
	fmt.Print("-----------------------------\n","File: ",wordListStat.Name(),"\nFile size: ",wordListStat.Size()/(1024),"KB\n")
	fmt.Println("Total words: ",len(wordlist))
	fmt.Println("-----------------------------")

	outfolder := filename
	status := false
	for _, password := range wordlist {
		status = unzip(filename, password, outfolder)
	}
	if status {
		fmt.Println(status)
	}
}

func unzip(filename string, password string, outfolder string) bool {

	zipfile, err := zip.OpenReader(filename)
	if err!=nil {
		panic(err.Error())
	}
	defer  zipfile.Close()

	for _, x := range zipfile.File {
		x.SetPassword(password)
		_, err := x.Open()
		if err!=nil {
			return false
		}
	}
	return true
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
		fmt.Print("$: 1. Create wordlist and attack 2. Use own wordlist and attact 3. Exit - press Ctrl + c $: ")
		choice, err := reader.ReadString('\n')
		if err!=nil {
			panic(err.Error())
		}
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			one(reader)
		case "2":
			two(reader)
		}
	}
}
