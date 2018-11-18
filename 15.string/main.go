package main

  // kumpulan library : https://golang.org/pkg/

	import "fmt"
  import "strings"  //import library strings
  import "strconv"  //convert tipe data

	func main() {
    myString := "Hello Golang"
    fmt.Println(myString)
    fmt.Println(len(myString))

    myStringUpper := strings.ToUpper(myString)
    fmt.Println(myStringUpper)

    myStringLower := strings.ToLower(myString)
    fmt.Println(myStringLower)

    resultContains := strings.Contains(myString,"Go") //cek "go" ada di myString ?
    fmt.Println(resultContains)

    resultSplit := strings.Split(myString," ") //split berdasarkan spasi   //return dynamic array/slice
    for _,v := range resultSplit {
      fmt.Println(v)
    }

    myStringTwo := "10"
    resultConvInt, err := strconv.Atoi(myStringTwo) //konversi string menjadi integer
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(resultConvInt * 5)

    resultConvStr := strconv.Itoa(10) //konversi integer menjadi string
    fmt.Println(resultConvStr)
  }
