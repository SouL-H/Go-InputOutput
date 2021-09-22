package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//Aynı veriyi iki yere birden yazmak istediğimizde
	//2 dosyaya ve konsola çıktı için
	txtFile, _ := os.Create("textDosya.txt")
	txtFile2, _ := os.Create("textDosya2.txt")
	mWriter := io.MultiWriter(os.Stdout, txtFile, txtFile2)

	//Üç hedefe yazma işlemi burada gerçekleşir.
	n, err := mWriter.Write([]byte("Çoklu dosya yazma\n"))
	if err != nil {
		fmt.Println(err)
	}
	txtFile.Close()
	txtFile2.Close()
	//Kaç byte veri yazdığını döner.
	fmt.Println(n)
}
