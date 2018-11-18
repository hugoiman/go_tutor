//  Interface kosong: bisa menampung segala jenis data, bahkan array, bisa pointer bisa tidak (dynamic typing)

package main

import (
  "fmt"
)

func main() {
  var secret interface{}    //  <- interface menjadi tipe data, bukan objek

  secret = "ethan hunt"
  fmt.Println(secret)

  secret = []string{"apple","banana","manggo"}
  fmt.Println(secret)

  secret = 12
  var number = secret.(int) * 10    //  casting variabel interface kosong ke integer
  fmt.Println(number)

  var x interface{} = &person{name: "wick", age: 27}
  var name = x.(*person).name     //  casting variabel interface kosong ke objek pointer
  fmt.Println(name)

  //  kombinasi map dengan interface
  var orang = []map[string]interface{}{
    {"name": "Wick", "age": 23},
    {"name": "Ethan", "age": 23},
    {"name": "Bourne", "age": 22},
  }

  for _, each := range orang {
    fmt.Println(each["name"], "age is", each["age"])
  }

  //  kombinasi slice denga interface
  var fruits = []interface{}{
    map[string]interface{}{"name": "strawberry", "total": 10},
    []string{"manggo", "pineapple", "papaya"},
    "orange",
  }

  for _, each := range fruits {
    fmt.Println(each)
  }
}

type person struct {
  name string
  age  int
}
