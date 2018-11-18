//  Slice: array dinamis. isi nya dapat ditambah berpapun

package main

  import "fmt"

  //Slice = elemen seperti array hanya saja bersifat dinamis (dapat ditambahkan datanya)
  func main(){
    mySlice := []int{10,20,30,40,50}
    fmt.Println(mySlice[0])           // cetak slice = sperti array

    for i,v := range mySlice {
      fmt.Println(i,v)                //cetak index & value slice
    }

    mySliceStr := []string{"Ogen1","Ogen2","Ogen3","Ogen4"}
    for _,v := range mySliceStr {     //  jika tidak ingin mencetak index menggunakan "_"
      fmt.Println(v)
    }

    mySliceStr = append(mySliceStr,"Ogen5")   //  menambah data di akhir
    for _,v := range mySliceStr {
      fmt.Println(v)
    }
  }
