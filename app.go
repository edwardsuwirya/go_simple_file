package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "data.txt"
	readFile(fileName)
	//writeFile(fileName)
}

func writeFile(fileName string) {
	// If file exist, truncate the content, else create a file
	//f, err := os.Create("data.txt")
	//defer f.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//f.WriteString("Hi enigmanians")

	// Create file if not exist, else append the string
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("Hah 2\n")
}

func readFile(fileName string) {
	// Manual way to read a file,
	//it is developer's responsibility to close the file
	//file, err := os.Open(fileName)
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
	//// Alternative using ioutil (deprecated already)
	//content, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(content))

	// Go provide built in package to simplify read a file
	// 1. Read all file content, be careful when the size of file too big
	// because it will store the content in memory
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	strContent := string(content)
	fmt.Println("Read content:", strContent)

	// 2. Read line by line
	f, err := os.Open(fileName)
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
