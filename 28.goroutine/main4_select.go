//  Select: u/ sharing data antar goroutine, bisa juga u/ kontrol goroutine yg banyak
//          penggunaannya seperti switch

package main

import (
  "fmt"
)

func main()  {
  fmt.Println("Golang Select Channel")

    var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
    fmt.Println("numbers :", numbers)

    var ch1 = make(chan float64)
    go getAverage(numbers, ch1)

    var ch2 = make(chan int)
    go getMax(numbers, ch2)

    for i := 0; i < 2; i++ {
        select {
        case avg := <-ch1:                    //  akan terpenuhi ketika ada penerimaan data dari channel ch1, yang kemudian akan ditampung oleh variabel avg.
          fmt.Printf("Avg \t: %.2f \n", avg)
        case max := <-ch2:
          fmt.Printf("Max \t: %d \n", max)    //  akan terpenuhi ketika ada penerimaan data dari channel ch2, yang kemudian akan ditampung oleh variabel max.
        }
    }
}

func getAverage(numbers []int, ch chan<- float64)  {
  var sum = 0
  for _,e := range numbers {
    sum += e
  }
  ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan<- int)  {    //  mengirim ch (sender)
  var max = numbers[0]
  for _,e := range numbers {
    if max < e {
      max = e
    }
  }
  ch <- max   //  menginisialisasi ch dengan max (receiver)
}
