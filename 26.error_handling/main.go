package main

import (
  "fmt"
  "strconv"
  "errors"
)

func main()  {
  iStr := "a"
  i,err := strconv.Atoi(iStr)   //  built in dari go

  if err != nil {   //  cek jika terjadi error
    fmt.Println("Terjadi error",err.Error())
  }else{
    fmt.Println(i)
  }

  // cara 1
  r, err := Div(10,0)
  if err != nil {
    fmt.Println("Terjadi error : ",err.Error())
  }else {
    fmt.Println(r)
  }

  //  cara 2 (inline: lebih simpel)
  if x,err := Div(25,2); err != nil{
    fmt.Println("Terjadi error : ",err.Error())
  }else {
    fmt.Println(x)
  }
}

func Div(x,y int) (int, error)  {
  if y == 0 {
    return 0, errors.New("Tidak bisa membagi dengan 0")
  }
  result := x / y
  return result, nil    //  retrun result = int, nil = error ( dari parameter)
}
