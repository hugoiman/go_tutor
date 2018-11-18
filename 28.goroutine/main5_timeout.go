// Timeout: mengontrol penerimaan data dari channel berdasarkan waktu diterimanya, dengan durasi timeout bisa ditentukan sendiri.
//          Ketika tidak ada aktivitas penerimaan data dalam durasi yang sudah ditentukan, callback akan dijalankan.

package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main()  {
  rand.Seed(time.Now().Unix())

  var messages = make(chan int)

  go sendData(messages)
  retreiveData(messages)
}

// goroutine mengirim data pada interval yg random
func sendData(ch chan<- int) {
  for i := 0; true; i++ {
    ch <- i
    time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
  }
}

func retreiveData(ch <-chan int) {
  loop:           //  infinity loop
  for {
    select {
      case data := <-ch:                                        //  akan terpenuhi ketika ada serah terima data pada channel messages
        fmt.Print(`receive data "`, data, `"`, "\n")
      case <-time.After(time.Second * 5):                       //  akan terpenuhi ketika tidak ada aktivitas penerimaan data dari channel dalam durasi 5 detik. Blok inilah yang kita sebut sebagai callback.
        fmt.Println("timeout. no activities under 5 seconds")
        break loop
    }
  }
}
