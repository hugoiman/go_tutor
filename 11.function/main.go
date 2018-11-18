package main

	import "fmt"

	func main() {
  // 1	->	fungsi return integer
    x := 5
    y := 10
    z := add(x,y)
    fmt.Println("Nilai z: ",z)

  // 2	->	fungsi retun string
    name := "ogen"
    result := hello(name)
    fmt.Println(result)

  // 3	->	fungsi return multiple
    xMulti := "hello"
    yMulti := "golang"
    resultX, resultY := multi(xMulti, yMulti)		//	karna mengembalikan 2 nilai, maka dibutuhkan 2 variabel u/ menampung
    fmt.Println(resultX, resultY)

  // 4	->	variabel di jadikan fungsi (fungsi sebagai value)
    plus := func(a,b int) int {
      return a + b
    }
    fmt.Println(plus(4,9))		//	panggil variabel dengan memberikan nilai ke fungsi

  // 5	->	memanggil fungsi dengan variabel
    nextValue := genNumber()		// nextValue u/ menampung fungsi genNumber
    fmt.Println(nextValue())		// output: 1
    fmt.Println(nextValue())		// output: 2

  // 6	-> memanggil fungsi berparameter dengan variabel
    lv := love(name)						//	masuk ke fungsi pertama
    fmt.Println(lv("You"))			//	masuk ke fungsi ke dua (anonymous fungsi)

  // 7	->
    f := func(v string) bool{				//	fungsi sebagai value.
      return v == "Golang"
    }
    resultMatch := match("Golang",f)	//	fungsi sebagai value. memanggil fungsi match sekaligus mengirim parameter
    fmt.Println(resultMatch)
  }

  // 1
  func add(x int, y int) int {				// bisa juga ditulis: func add(x, y int) int {}
    return x + y
  }

  // 2
  func hello(name string) string {
    return fmt.Sprintf("Hello %s", name)		//	Sprintf: seperti concat/menggabungkan -> in this case: mengembalikan nilai dalam bentuk string
  }

  // 3.Multiple return
  func multi(x string, y string) (string,string) {
    return y,x
  }

  // 5.Fungsi return fungsi (Closure)
  func genNumber() func() int {		//	mengembalikan fungsi lain yg memiliki tipe data int.	func() int: disebut anonymous fungsi karna tidak punya nama
    c := 0
    return func() int {
      c = c + 1
      return c
    }
  }

  // 6. Sama seperti poin 5 (Closure)
  func love(nama string) func(string) string {
    return func(sesuatu string) string {
      return fmt.Sprintf("%s love %s", nama, sesuatu)
    }
  }

  // 7.Fungsi sebagai parameter/argument
  func match(v string, f func(string) bool) bool  {		//	func(string) bool: anonymous fungsi
    if f(v) {
      return true
    }
    return false
  }
