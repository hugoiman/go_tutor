//  Goroutine (Mini Thread):  cara mengeksekusi fungsi tanpa menggangu fungsi lain.
//                            misal: 1 fungsi terdapat corrupt, maka fungsi yg lain tidak terganggu eksekusinya

package main

import (
  "fmt"
  // "time"
)

//  channel: digunakan untuk synchronize agar goroutine 1 dan lainnya dapat dijalankan bersama
func main()  {                    //  sebenarnya main() adalah goroutine/ antara goroutine tidak saling peduli karna goroutine di eksekusi di alamat memori yg berbeda

  done := make(chan bool)         //  membuat channel dengan tipe data boolean

  go helloGo(done)                    //  pembuatan Goroutine dengan keyword "go". untuk synchronize antara goroutine main() dan helloGo()

  <-done                              //  mengambil/menerima value variabel done, yaitu true (channer receiver)
  //  time.Sleep(1 * time.Second)     //  <- tidak dibutuhkan karna sudah ada channel u/ synchronize
  fmt.Println("Ini fungsi main")
}

func helloGo(done chan bool)  {   //  done <- refactor dengan tipe data chan boolean
  fmt.Println("Hello Go")
  done <- true                    // mengisi channel done dgn boolean true (channel sender)
}
