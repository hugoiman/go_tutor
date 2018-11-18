//  Time, Parsing Time, Format Time: pemanfaatan data bertipe waktu, method-method yang disediakan, dan juga format & parsing data string ke tipe time.Time dan sebaliknya
//  https://dasarpemrogramangolang.novalagung.com/38-time.html
package main

import (
  "fmt"
  "time"
)

func main() {
  var time1 = time.Now()                                          //  mengembalikan informasi waktu sekarang
  fmt.Printf("time1 %v\n", time1)
  // time1 2015-09-01 17:59:31.73600891 +0700 WIB

  var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)     //  mengembalikan informasi waktu yang ditentukan sendiri. parameter(thn,bulan,tgl,jam,mnt,nanodetik,timezone)
  fmt.Printf("time2 %v\n", time2)
  // time2 2011-12-24 10:20:00 +0000 UTC


  //  Parsing Time
      // Data string bisa dikonversi menjadi time.Time dengan memanfaatkan time.Parse. Fungsi ini membutuhkan 2 parameter:
      // Parameter ke-1 adalah layout format dari data waktu yang akan diparsing.
      // Parameter ke-2 adalah data string yang ingin diparsing.
  var layoutFormat, value string
  var date time.Time

  layoutFormat = "2006-01-02 15:04:05"        //  <- 2006 = tahun 4 dgt, 01 = bulan 2 dgt, 02 = tgl 2 dgt, 15 = format 24 jam, 04 = menit 2 dgt, 05 = detik 2 dgt
  value = "2015-09-02 08:04:00"
  date, _ = time.Parse(layoutFormat, value)
  fmt.Println(value, "\t->", date.String())   // 2015-09-02 08:04:00 +0000 UTC

  layoutFormat = "02/01/2006 MST"
  value = "02/09/2015 WIB"
  date, _ = time.Parse(layoutFormat, value)
  fmt.Println(value, "\t\t->", date.String()) // 2015-09-02 00:00:00 +0700 WIB


  //  Predefined Layout Format Untuk Keperluan Parsing Time -> auto format

  var date2, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
  fmt.Println(date2.String())    // 2015-09-02 08:00:00 +0700 WIB
}
