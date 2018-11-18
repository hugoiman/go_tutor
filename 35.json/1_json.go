//  JSON: Javascript Object Notation adalah notasi standar yang umum digunakan untuk komunikasi data via web. JSON merupakan subset dari javascript.

//  1. import package dan siapkan struct
//  2.
package main

import (
  "fmt"
  "encoding/json"
)

func main()  {
  player1 := Player{ID: 1, Name: "Iniesta"}       //  <- akan di encode, dijadikan json

  p1, err := json.Marshal(player1)                //  json.Marshall() u/ mengembalikan data bertipe []byte dan error
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(string(p1))                         //  u/ konversi dari byte ke string

  fmt.Println("=====================")

  data := []byte(`{"id": 1, "name": "Ronaldo"}`)           //  nama field harus sama persis dgn nama field di struct

  var p2 Player
  err = json.Unmarshal(data, &p2)                          // Unmarshal: u/ konversi json string menjadi objek. param1 bertipe []byte, param2 bertipe pointer
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(p2)
}

type Player struct {
  ID int `json:"id"`                              // u/ custom field menjadi lower cast
  Name string `json:"name"`
}
