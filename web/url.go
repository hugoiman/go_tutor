//  Data string url bisa dikonversi kedalam bentuk url.URL. Dengan menggunakan tipe tersebut akan ada banyak informasi yang bisa kita manfaatkan,
//  diantaranya adalah jenis protokol yang digunakan, path yang diakses, query, dan lainnya.

package main

import "fmt"
import "net/url"

func main() {
  var urlString = "http://depeloper.com:80/hello?name=john wick&age=27"
  var u, e = url.Parse(urlString)   //  url.Parse() u/ parsing string ke bentuk url. Mengembalikan 2 data, variabel objek bertipe url.URL dan error (jika ada).
  if e != nil {
    fmt.Println(e.Error())
    return
  }

  fmt.Printf("url: %s\n", urlString)

  fmt.Printf("protocol: %s\n", u.Scheme) // http
  fmt.Printf("host: %s\n", u.Host)       // depeloper.com:80
  fmt.Printf("path: %s\n", u.Path)       // /hello

  //  query yang ada pada url akan otomatis diparsing juga, menjadi bentuk map[string][]string,
  //  dengan key adalah nama elemen query, dan value array string yang berisikan value elemen query.
  var name = u.Query()["name"][0] // john wick
  var age = u.Query()["age"][0]   // 27
  fmt.Printf("name: %s, age: %s\n", name, age)
}
