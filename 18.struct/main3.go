//Struct sebagai

package main

  import "fmt"

  func main()  {
    p := Person{
      ID: 1,
      Name: "Ogen",
      Age: 22,
    }
    fmt.Println(p.getID())
    fmt.Println(p.getName())
    fmt.Println(p.getAge())
    p.changeName("Boiboi")
    fmt.Println(p.getName())

    x := Person{             //  Pointer Person
      ID: 2,
      Name: "Hugo",
      Age: 23,
    }
    fmt.Println(x.getID())
    fmt.Println(x.getName())
    fmt.Println(x.getAge())
    x.changeName("Boiboi2")
    fmt.Println(x.getName())
  }

  type Person struct {
    ID int
    Name string
    Age int
  }

  func (p *Person) changeName (newName string)  {     //  (p Person) : Receiver
    p.Name = newName
  }

  func (p Person) getID() int  {
    return p.ID
  }

  func (p Person) getName() string {
    return p.Name
  }

  func (p Person) getAge() int  {
    return p.Age
  }
