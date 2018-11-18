//  Goroutine (Mini Thread):  cara mengeksekusi fungsi tanpa menggangu fungsi lain.
//                            misal: 1 fungsi terdapat corrupt, maka fungsi yg lain tidak terganggu eksekusinya

package main

import (
  "fmt"
  "time"
)

func main()  {                    //  sebenarnya main() adalah goroutine/ antara goroutine tidak saling peduli karna goroutine di eksekusi di alamat memori yg berbeda
  go helloGo()                    //  pembuatan Goroutine dengan keyword "go"

  time.Sleep(1 * time.Second)     // <-  untuk menunggu sampai go routine di eksekusi
  fmt.Println("Ini fungsi main")
}

func helloGo()  {
  fmt.Println("Hello Go")
}
