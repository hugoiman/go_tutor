//  Default Channel = Un-Buffered Channel: jika goroutine ingin mengirimkan data (sender), maka HARUS ada goroutine yg bertugas menerima data tsb (receiver)

//  Buffered Channel = Tidak peduli receiver sudah ada/belum. jumlah data yg dikirim ditentukan jumlah buffer-nya
package main

import (
  "fmt"
)

func main() {

  // Channel Buffer
  hello := make(chan string, 2)
  hello <- "Hello"
  hello <- "Golang"
  close(hello)      //  u/ loop channel, maka channel harus di close dahulu. menandakan tidak ada data yg dikirimkan lagi

  // fmt.Println(<-hello)
  // fmt.Println(<-hello)

  for v := range hello {
    fmt.Println(v)
  }
}
