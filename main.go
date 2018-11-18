package main

	import "fmt"

	func main() {
	//static type declaration
		// var a int
		// var b float64
		// a = 10
		// b = 5.5
		// fmt.Println(a)
		// fmt.Println(b)
		// fmt.Printf("b bertipe data %T\n",b)

	//dynamic type declaration
		// c := "ogen"
		// d := 20
		// fmt.Println(c)
		// fmt.Println(d)

	// operator
		// x := 10
		// y := 5
		//  xx := true
		// yy := false

		// fmt.Println(x+y)
		// fmt.Println(x-y)
		// fmt.Println(x/y)
		// fmt.Println(x*y)
		// fmt.Println(x%y)

		// fmt.Println(x == y)
		// fmt.Println(x != y)
		// fmt.Println(xx && yy)
		// fmt.Println(xx || yy)
		// fmt.Println(!xx)

	//if else
		// x := 10
		// y := 5
		// if x == y {
		// 	fmt.Println("ini sama")
		// } else if x < y {
		// 	fmt.Println("Ini X")
		// } else {
		// 	fmt.Println("Default")
		// }

	//switch case
		// x := 10
		// switch x {
		// case 60:
		// 	fmt.Println("Ini 60")
		// case 30:
		// 	fmt.Println("Ini 30")
		// case 10:
		// 	fmt.Println("Ini 10")
		//	fallthrough						//	u/ memelanjutkan pengecekan di case selanjutnya
		// default:
		// 	fmt.Println("Default")
		// }

		// var y interface{}
		// y = 5
		// switch y.(type) {
		// case int:
		// 	fmt.Println("Integer")
		// case float64:
		// 	fmt.Println("String")
		// default:
		// 	fmt.Println("Default")
		// }

	//loop
		x := 4
		// y := 10
		// for i := 1;i <= 5;i++ {
		// 	fmt.Println(i)
		// }

		// for x < y {
		// 	x++
		// 	fmt.Println(x)
		// }

		for {
			x++
			if x == 6 {
				continue
			}else if x == 8 {
				break
			}
			fmt.Println(x)
		}

	}
