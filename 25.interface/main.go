//  Interface : kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), dan dibungkus dengan nama tertentu.
//              biasa disebut Polymorfisme

package main

import (
  "fmt"
  "math"
)

func main()  {
  circle := Circle {X: 1.0, Y: 1.0, Radius:10.5}
  rectangle := Rectangle {Width:10.0, Height:5.0}

  fmt.Println("Area dari circle adalah: \n",getArea(circle))
  fmt.Println("Area dari rectangle adalah: \n",getArea(rectangle))
}

type Shape interface {  //  declare Interface
  Area() float64        //  Shape mempunyai method Area
}

type Circle struct {
  X float64
  Y float64
  Radius float64
}

type Rectangle struct {
  Width float64
  Height float64
}

func (c Circle) Area() float64  {       //  menjadi polimorfisme dari Shape, maka harus meng-implementasi-kan method Area()
  return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64  {
  return r.Width * r.Height
}

func getArea(s Shape) float64{          //
  return s.Area()
}
