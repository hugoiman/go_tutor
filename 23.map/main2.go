package main

import "fmt"

func main()  {
  // var mapA map[int] string         //static declare
  // mapA := make(map[int] string)    //dynamic declare
  mapPlayer :=  make(map[int] Player)
  mapPlayer[1] = Player{ID: 1,Name: "Zlatan"}
  mapPlayer[2] = Player{ID: 2,Name: "Iniesta"}
  mapPlayer[3] = Player{ID: 3,Name: "Xavi"}

  for _,v := range mapPlayer {
    fmt.Println(v.Name)
  }
}

type Player struct {
  ID int
  Name string
}
