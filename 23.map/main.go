//  Map: seperti Slice, bedanya pada Map indeknya dapat ditentukan sendiri

package main

import "fmt"

func main()  {
  var mapPerson map[int] string     //  membuat map. int adalah key, dan valuenya = string
  mapPerson = make(map[int] string) //  make: fungsi built in dari go

  mapPerson[1] = "Alex"
  mapPerson[2] = "Ben"
  mapPerson[3] = "Geno"
  mapPerson[4] = "ogen"

  for k,v := range mapPerson {  // k = key, v = value
    fmt.Println(k, v)           // k jika di array bernama index (i)
  }

  ogen, ok := mapPerson[4]  //  mencari value ogen apakah true(terdapat di mapPerson index ke-4)
  if !ok {
    fmt.Println("Ogen tidak ada di mapPerson")
  }else {
    fmt.Println(ogen)
  }
}
