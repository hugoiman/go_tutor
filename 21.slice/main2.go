package main

import "fmt"

func main()  {
  var players []Player
  players = []Player {Player{ID: 1, Name: "Hazard"}, Player{ID: 2,Name: "Messi"}}
  players = append(players, Player{ID: 3,Name: "Ronaldo"})

  for _,v := range players {
    fmt.Println(v)
  }
}

type Player struct {
  ID int
  Name string
}
