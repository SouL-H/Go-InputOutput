package main

import (
	"bufio"
	"fmt"
	"os"
)

//Satır satır okuma

func main() {
	//Yeni bir txt oluşturdul içine veri yazdık.
	os.WriteFile("satir.txt", []byte("line 1\nline 2\nline 3"), os.ModePerm)
	testFile, err := os.Open("satir.txt")
	if err != nil {
		fmt.Println(err)
	}
	//scanner değeri okunan veri bilgisi dönüyor.
	scanner := bufio.NewScanner(testFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
