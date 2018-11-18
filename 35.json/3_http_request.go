//  https://dasarpemrogramangolang.novalagung.com/52-http-request.html

package main

import "fmt"
import "net/http"
import "encoding/json"
import "bytes"
import "net/url"

var baseURL = "http://localhost:8080"

//  Penggunaan HTTP Request
// func main() {
//   var users, err = fetchUsers()
//   if err != nil {
//     fmt.Println("Error!", err.Error())
//     return
//   }
//
//   for _, each := range users {
//     fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
//   }
// }

//  HTTP Request Dengan Form Data
func main() {
  var user1, err = fetchUser("E001")
    if err != nil {
      fmt.Println("Error!", err.Error())
      return
    }

    fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}

type student struct {
  ID    string
  Name  string
  Grade int
}

// fetchUsers() bertugas melakukan request ke http://localhost:8080/users, menerima response dari request tersebut, lalu menampilkannya.
func fetchUsers() ([]student, error) {
  var err error
  var client = &http.Client{}   // u/ menghasilkan instance http.Client. Objek ini nantinya diperlukan untuk eksekusi request.
  var data []student

  request, err := http.NewRequest("POST", baseURL+"/users", nil)  //  NewRequest() u/ membuat request baru.  u/ menghasilkan instance http.Request.
                                                                  //  param1: method, param2: url tujuan, param3: form data request (jika ada)
  if err != nil {
    return nil, err
  }

  response, err := client.Do(request) //  Do() adalah cara eksekusi request. mengembalikan instance bertipe http.Response
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)  //  Body: u/ mengambil data response dalam bentuk string
                                                      //  json.NewDecoder(response.Body).Decode(&data): mengkonversi string menjadi bentuk JSON.
  if err != nil {
    return nil, err
  }

  return data,nil
}

//  Response dari rute /user bukan berupa array objek, melainkan sebuah objek.
//  Maka pada saat decode pastikan tipe variabel penampung hasil decode data response adalah student (bukan []student).
func fetchUser(ID string) (student, error) {
    var err error
    var client = &http.Client{}
    var data student

    var param = url.Values{}  //  akan menghasilkan objek yang nantinya digunakan sebagai form data request.
    param.Set("id", ID)       //  objek diatas perlu di set data apa saja yang ingin dikirimkan
    var payload = bytes.NewBufferString(param.Encode())   //  objek form data di-encode lalu diubah menjadi bentuk bytes.Buffer,
                                                          //  yang nantinya disisipkan pada parameter ketiga pemanggilan fungsi http.NewRequest().

    request, err := http.NewRequest("POST", baseURL+"/user", payload)
    if err != nil {
      return data, err
    }
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded") //  Karena data yang akan dikirim di-encode, maka pada header perlu di set tipe konten request-nya sbg param2

    response, err := client.Do(request)
    if err != nil {
      return data, err
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
      return data, err
    }

    return data, nil
}
