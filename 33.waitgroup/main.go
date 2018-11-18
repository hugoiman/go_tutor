// Wait Group: u/ synchronize goroutine
//  https://dasarpemrogramangolang.novalagung.com/56-waitgroup.html

package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  var wg sync.WaitGroup       //  membuat variabel Wait Group

  wg.Add(3)                   //  menambahkan 3 goroutine, yaitu: p1,p2,p3

  p1 := <-process(1, 3, &wg)
  p2 := <-process(2, 1, &wg)
  p3 := <-process(3, 4, &wg)

  fmt.Println(p1)
  fmt.Println(p2)
  fmt.Println(p3)

  wg.Wait()   //  menunggu goroutine yg di jalankan. proses eksekusi program tidak akan diteruskan ke baris selanjutnya, sebelum sejumlah 5 goroutine selesai.
}

func process(id int, delay time.Duration, wg *sync.WaitGroup) <-chan string {
  output := make(chan string)

  go func ()  {
    defer wg.Done()   //  memberikan informasi bahwa goroutine ini telah selesai di panggil
    time.Sleep(time.Second * delay)

    output <- fmt.Sprintf("Process %d done",id)
  }()

  return output
}
