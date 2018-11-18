//	pointer : variabel yang berisikan ALAMAT MEMORI suatu nilai

package main

	import "fmt"

  func main(){
    var hello string = "Hello"
    var strPointer *string			//	u/ membuat pointer dengan nama variabel strPointer bertipe string (metode deferencing)

    strPointer = &hello					//	u/ mengakses alamat memori dengan '&' didepan variabel (metode referencing)
    fmt.Println(&hello)
    fmt.Println(strPointer)

    fmt.Println(hello)
    change(hello)
    fmt.Println(hello)					//	output :Hello

    changePointer(&hello)
    fmt.Println(hello)
  }

  func change(v string)  {
    v = v + " Golang"
  }
  //Perubahan pada pointer
  func changePointer(v *string)  {	//	mendapatkan alamat memori
    *v = *v + " Golang"							//	output: Hello Golang
  }
