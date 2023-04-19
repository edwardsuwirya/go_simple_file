package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//readFile()
	writeFile()
}

func writeFile() {
	// If file exist, truncate the content, else create a file
	//f, err := os.Create("data.txt")
	//defer f.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//f.WriteString("Hi enigmanians")

	// Create file if not exist, else append the string
	f, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("Hah 2\n")
}

func readFile() {
	//file, err := os.Open("hello.txt")
	//defer file.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fileInfo, err := file.Stat()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fileSize := fileInfo.Size()
	//
	//readBuffer := make([]byte, fileSize)
	//_, err = file.Read(readBuffer)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s", string(readBuffer))
	//
	//// Alternative using ioutil
	//content, err := ioutil.ReadFile("hello.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(content))

	// Read line by line
	f, err := os.Open("hello.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(".", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
