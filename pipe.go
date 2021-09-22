package main

import (
	"fmt"
	"io"
	"time"
)

//Veriyi bir noktadan bir noktaya taşımak için
//2 farklı go routinden biri okuma diğeri yazma işlemi yapıyor
//İkisinin ortak haberleşmesi için
//Memory buffer olmuyor maliyet oluşturmaaz
func main() {

	pReader, pWriter := io.Pipe()
	//create channel
	done := make(chan struct{})
	go read(pReader, done)
	go write(pWriter)
	<-done
}
func read(reader io.Reader, done chan struct{}) {
	//1024 byte oku her defasında
	buff := make([]byte, 1024)
	for {
		//okunanın ne kadar veri varsa int değerini  ve hata dönderir.
		//Hata döndüğü halde hata dönebilir.
		//Bu yüzden EOF kontrolü yapılır (End of file) okuma yeri kapanmış mı
		readed, err := reader.Read(buff)
		if readed == 0 {
			//okunan yeri dönder
			if err == io.EOF {
				//Memory alan kaplamamak için anonim empty struct kullandık.
				done <- struct{}{}
				break
			}
			//Hata alınan yeri dönder.
			if err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		} else {
			//int değer döndüğü için okuduğu kadarını yaz diyoruz
			fmt.Println(buff[:readed])
			//Tekrar bağlantı ya da okuma durumunda hata olmaması için kontrol koyduk.
			if err == io.EOF || err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		}
	}

}

//Hedefe yazma
func write(writer *io.PipeWriter) {
	i := 0
	for {
		if i == 10 {
			writer.Close()
			break
		}
		//Hedefe yazdırma
		writer.Write([]byte(string(i)))
		i++
		time.Sleep(time.Millisecond * 100)
	}
}
