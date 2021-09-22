package main

import (
	"fmt"
	"os"
)

func main() {
	writeFile()
}
func writeOsFile() {
	err := os.WriteFile(
		"testFile.txt", []byte("test"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
func writeFile() {
	f, err := os.Create("testFileWrite.txt")
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte("1\n"))
	f.Write([]byte("2\n"))
	f.Write([]byte("3\n"))
	f.Write([]byte("4\n"))
	f.Write([]byte("5\n"))
	f.Write([]byte("6\n"))
	f.Write([]byte("7\n"))
	f.Write([]byte("8\n"))
	f.Write([]byte("9\n"))
	f.Write([]byte("10\n"))
	f.Close()

}
