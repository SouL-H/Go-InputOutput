package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//readFile()
	//readFileBufio()
	readFileSeek()
}

func readFile() {
	//Komple dosya içeriğini dönderir.
	readTest, err := os.ReadFile("testFileWrite.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(readTest))
	}
}

func readFileBufio() {
	//Dosyanın pointerine erişilir.
	readTest, err := os.Open("testFileWrite.txt")
	if err != nil {
		fmt.Println(err)
	} else {

		bufReader := bufio.NewReader(readTest)
		//Stdout bi writer implement ediyor.
		io.Copy(os.Stdout, bufReader)
	}
}

//Seek fonksiyonu ile bir kısım okuma işlemi yapabiliriz.
func readFileSeek() {
	f, err := os.Open("testFileWrite.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		//Hangi ofsetten başlayıp kime göre başlanacağı belirtilir.
		f.Seek(2, 1)
		readByte := make([]byte, 16)
		f.Read(readByte)
		fmt.Println(string(readByte))
	}
}
