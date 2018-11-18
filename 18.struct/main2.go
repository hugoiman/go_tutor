//  Struct sebagai parameter fungsi
package main

  import "fmt"

  func main()  {
    p := Person{
      ID: 1,
      Name: "Ogen",
      Age: 22,
    }
    printPerson(p)
  }

  type Person struct {
    ID int
    Name string
    Age int
  }

  func printPerson(p Person)  {   //  struct sebagai parameter fungsi
    fmt.Println("ID = ",p.ID)
    fmt.Println("Name = ",p.Name)
    fmt.Println("Age = ",p.Age)
  }
