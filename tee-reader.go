package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//Stream okuyalım hem bir yere yazalım
//Bu konuda io.teeReader işimize yarar.
//Reader,Writer alır ve geriye de Reader döner
func main() {

	sReader := strings.NewReader("Test String")
	//Çıktı olarak 2 tane aynı veri basılır.
	//teereader hem okur hem ekrana basar.
	tReader := io.TeeReader(sReader, os.Stdout)
	fmt.Println("Start")
	//Streamdaki bütün datayı okumamızı sağlar.
	readedBytes, _ := io.ReadAll(tReader)
	fmt.Println("\nString okundu")
	fmt.Println(string(readedBytes))

}
