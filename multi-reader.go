package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	//Bir stringten reader oluşturmak için yöntem.
	str1 := strings.NewReader("birinci string\n")
	str2 := strings.NewReader("ikinci string")
	//Birden fazla reader'i birleştirmek için kullanılır.
	mReader := io.MultiReader(str1, str2)

	io.Copy(os.Stdout, mReader)
}
