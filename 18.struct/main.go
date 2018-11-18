//  Struct: kumpulan definisi variabel dengan tipe data yang bervariasi, dengan struct data terbungkus lebih rapi dan mudah di maintain
//          digunakan u/ mencetak variabel objek, seperti konsep class pada pemrograman berbasis objek

package main

  import "fmt"

  type Vector struct {
    X,Y int;
  }

  type Player struct {
    ID int
    Name string
  }

  func main()  {
    var v Vector
    v.X = 10
    v.Y = 5

    fmt.Println(v)
    fmt.Println("X = ",v.X)
    fmt.Println("Y = ",v.Y)

    player1 := Player{ID: 1,Name: "Neymar"}

    fmt.Println(player1.ID)
    fmt.Println(player1.Name)
  }
