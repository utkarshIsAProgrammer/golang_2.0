package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt") // open a file
	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat() // describe the file structure
	if err != nil {
		panic(err)
	}

	fmt.Println(fileInfo.Name())    // print file name
	fmt.Println(fileInfo.IsDir())   // check isDir
	fmt.Println(fileInfo.Size())    // check file size (1 char = 1 byte)
	fmt.Println(fileInfo.Mode())    // check file mode
	fmt.Println(fileInfo.ModTime()) // check file modification time

	defer file.Close() // close file

	/*
		buff := make([]byte, 12)
		buffSize, err := file.Read(buff)
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(buff); i++ {
			fmt.Print(string(buff[i]))
		}
		fmt.Println("\n", buffSize) */

	f1, err := os.ReadFile("example.txt") // reading entire file
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f1))

	dir, err := os.Open("../") // open folder
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	info, err := dir.ReadDir(-1) // read directory
	for _, fi := range info {
		fmt.Println(fi.Name())
		// fmt.Println(fi.IsDir())
	}

	nf, err := os.Create("examples/example.txt")
	if err != nil {
		panic(err)
	}

	defer nf.Close()

	// nf.WriteString("Learning Golang is fun!\n")

	// replace the previous content
	byts := []byte("Hello Golang!")
	nf.Write(byts)

	// remove a file
	err = os.Remove("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file deleted!")

}
