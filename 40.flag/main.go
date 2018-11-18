//  Flag: Flag memiliki kegunaan yang sama seperti arguments, yaitu untuk parameterize eksekusi program, dengan penulisan dalam bentuk key-value.
// https://dasarpemrogramangolang.novalagung.com/45-command-line-args-flag.html

package main

import "fmt"
import "flag"

func main() {
  var name = flag.String("name", "anonymous", "type your name")   //flag bertipe string. param1 = key, param2 = nilai default param3 = keterangan
  var age = flag.Int64("age", 25, "type your age")

  flag.Parse()
  fmt.Printf("name\t: %s\n", *name)   //  nilai balik dari flag.String() adalah string pointer, maka harus di dereference agar bisa mendapat nilai aslinya
  fmt.Printf("age\t: %d\n", *age)
}

//  run : go run main.go -name="ogen" -age=22
