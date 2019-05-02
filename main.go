package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/yeka/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func cleanup() {
	if r := recover(); r!=nil {
		log.Fatal(r)
	}
}

func bruteforce(filename string, reader *bufio.Reader) {
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

	fmt.Print("$: Enter the filename: ")
	filename1, err := reader.ReadString('\n')
	filename1 = strings.TrimSpace(filename1)
	if err!=nil {
		panic(err.Error())
	}

	status := false
	t1 := time.Now()
	for index, password := range wordlist {
		status = unzip(filename1, password)
		t2 := time.Now()
		fmt.Println("Password - ",password," | Status - ",status,"| Password left - ",len(wordlist)-index+1)
		if status {
			fmt.Println("-----------------------------------")
			fmt.Println("Password is - ",password)
			fmt.Println("Total time taken - ",t2.Sub(t1))
			fmt.Println("-----------------------------------")
			break
		}
	}
	t2 := time.Now()
	if !status {
		fmt.Println("-----------------------------------")
		fmt.Println("Password is - Not found")
		fmt.Println("Total time taken - ",t2.Sub(t1))
		fmt.Println("-----------------------------------")
	}
}

func unzip(filename string, password string) bool {
	defer cleanup()

	zipfile, err := zip.OpenReader(filename)
	if err!=nil {
		fmt.Println("here")
		panic(err.Error())
	}
	defer  zipfile.Close()

	buffer := new(bytes.Buffer)
	for _, x := range zipfile.File {
		x.SetPassword(password)
		r, err := x.Open()
		if err!=nil{
			return false
		}
		n, err := io.Copy(buffer, r)
		if n == 0 || err != nil {
			return false
		}
		break
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
	bruteforce(filename, reader)
}

func two(reader *bufio.Reader) {
	defer cleanup()

	fmt.Println("$: Enter the filename")
	filename, err := reader.ReadString('\n')
	if err!=nil {
		panic(err.Error())
	}
	filename = strings.TrimSpace(filename)

	bruteforce(filename, reader)
}

func main() {
	defer cleanup()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$: 1. Create wordlist and attack 2. Use own wordlist and attack 3. Exit - press Ctrl + c $: ")
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
